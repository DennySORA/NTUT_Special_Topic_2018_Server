package MongoDB

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Base"
	"SORA/Project/Go_Back_End_SEGA_Project/Config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ============================================[CreateAccount]

func DBCreateAccount(AccountIDPW Base.NewAccountIDPW, AccountData Base.NewAccountUser) Base.CreateReturn {
	// ====================================[DBLink]
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
	Collection := Database.C("Register")
	// -------------------------------------
	User := UserConvert(AccountData, AccountIDPW.Account)
	// -------------------------------------
	Collection = Database.C("User")
	err = Collection.Insert(User)
	ERRs(err)
	// -------------------------------------
	result := bson.M{}
	err = Collection.Find(bson.M{"Email": AccountIDPW.Account}).One(&result)
	ERRs(err)
	Oid, _ := result["_id"].(bson.ObjectId)
	// -------------------------------------
	Register := ConvertRegisters(AccountIDPW, Oid)
	Collection = Database.C("Register")
	err = Collection.Insert(Register)
	ERRs(err)
	// -------------------------------------
	return Base.CreateReturn{Status: Base.SelfSuccess(1), ID: AccountIDPW.Account}
}

// ============================================[LogIn]

func DBLogIn(Account string, Password string) Base.LogInToken {
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
	Collection := Database.C("Register")
	result := bson.M{}
	// -------------------------------------
	err = Collection.Find(bson.M{"Accountid": Account, "Password": GetSHAString(Password)}).One(&result)
	// -------------------------------------
	if err != nil {
		return Base.LogInToken{Status: Base.SelfErrors(4)}
	} else {
		ReturnData := ConverLogInToken(Account, 1)
		// -------------------------------------
		Collection := Database.C("User")
		// ========================================
		selects := bson.M{"Email": Account}
		data := bson.M{"$addToSet": bson.M{
			"SiginHistory": bson.M{
				"Times":    GetUTCTime(),
				"UseToken": ReturnData.AccountToken[:10],
				"Types":    "Null",
				"Device":   "Null",
			}}}
		// ========================================
		err = Collection.Update(selects, data)
		if err != nil {
			return Base.LogInToken{Status: Base.SelfErrors(4)}
		} else {
			return ReturnData
		}
		// -------------------------------------
	}
}

// ============================================[LogOut]

func DBLogOut(Account string, Token string) Base.StatusData {
	Invalid := TokenInvalid(Account, Token, 1)
	if Invalid == false {
		return Base.SelfErrors(8)
	}
	// ========================================
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
	Collection := Database.C("User")
	// ========================================
	selects := bson.M{"Email": Account}
	data := bson.M{"$addToSet": bson.M{
		"LogoutHistory": bson.M{
			"Times":    GetUTCTime(),
			"UseToken": Token[:10],
			"Types":    "Null",
			"Device":   "Null",
		}}}
	// ========================================
	err = Collection.Update(selects, data)
	if err != nil {
		return Base.SelfErrors(0)
	} else {
		return Base.SelfSuccess(9)
	}
	// -------------------------------------
}

// ============================================[AccountHas]

func DBAccountHas(ID string) bool {
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
	Collection := Database.C("Register")
	result := bson.M{}
	err = Collection.Find(bson.M{"Accountid": ID}).One(&result)
	if err == nil {
		return true
	} else {
		return false
	}
}
