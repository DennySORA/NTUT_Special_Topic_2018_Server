package MongoDB

import (
	"SORA/Base"
	"SORA/Config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ============================================[CreateAccount]

func DBCreateAccount(AccountIDPW Base.NewAccountIDPw, AccountData Base.NewAccountUser) *Base.CreateReturn {
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
	return &Base.CreateReturn{Status: Base.SelfSuccess(1), ID: AccountIDPW.Account}
}

// ============================================[LogIn]

func DBLogIn(Account string, Password string) *Base.LogInToken {
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
		return &Base.LogInToken{Status: Base.SelfErrors(4)}
	} else {
		ReturnData := ConverLogInToken(Account, 1)
		// -------------------------------------
		Collection := Database.C("User")
		// ========================================
		selects := bson.M{"Email": Account}
		data := bson.M{"$addToSet": bson.M{
			"SiginHistory": bson.M{
				"Times":    GetUTCTime(),
				"UseToken": ReturnData.AccountToken,
				"Types":    "Null",
				"Device":   "Null",
			}}}
		// ========================================
		err = Collection.Update(selects, data)
		if err != nil {
			return &Base.LogInToken{Status: Base.SelfErrors(4)}
		} else {
			return &ReturnData
		}
		// -------------------------------------
	}
}

// ============================================[LogOut]

func DBLogOut(Account string, Token string) *Base.StatusData {
	Invalid := TokenInvalid(Account, Token, 1)
	if Invalid == false {
		returnData := Base.SelfErrors(8)
		return &returnData
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
			"UseToken": Token,
			"Types":    "Null",
			"Device":   "Null",
		}}}
	// ========================================
	err = Collection.Update(selects, data)
	if err != nil {
		returnData := Base.SelfErrors(0)
		return &returnData
	} else {
		returnData := Base.SelfSuccess(9)
		return &returnData
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
