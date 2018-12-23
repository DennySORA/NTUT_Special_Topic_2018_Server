package MongoDB

import (
	Base "sega/Base"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DatabaseURL = "127.0.0.1:27017"
const DatabaseName = "Test"

// ===========================================================[Account]
// ============================================[CreateAccount]
func DBCreateAccount(AccountIDPW Base.NewAccountIDPW, AccountData Base.NewAccountUser) Base.CreateReturn {
	// ====================================[DBLink]
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
	Collection := Database.C("Register")
	// -------------------------------------
	User := UserConvert(AccountData, AccountIDPW.Account)
	// fmt.Printf("%v\n", User)
	// -------------------------------------
	Collection = Database.C("User")
	err = Collection.Insert(User)
	ERRs(err)
	// -------------------------------------
	result := bson.M{}
	err = Collection.Find(bson.M{"Email": AccountIDPW.Account}).One(&result)
	ERRs(err)
	// fmt.Printf("%v\n", result["_id"])
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
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
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

// ============================================[GetUser]
func DBGetUser(Account string, Token string) (Base.Users, int) {
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
	Collection := Database.C("User")
	result := bson.M{}
	err = Collection.Find(bson.M{"Email": Account}).One(&result)
	if err != nil {
		return Base.Users{}, 8
	} else {
		return ReturnUserConvert(result), 0
	}
}

// ===========================================================[CarID]
// ============================================[AddCarID]
func DBAddCarID(Account string, CarID string, CarName string) (string, int) {
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
	Collection := Database.C("User")
	// ========================================
	selects := bson.M{"Email": Account}
	data := bson.M{"$addToSet": bson.M{
		"Car": bson.M{
			"CarID":       CarID,
			"CarName":     CarName,
			"RefreshTime": GetUTCTime(),
			"CreateTime":  GetUTCTime(),
		}}}
	// ========================================
	err = Collection.Update(selects, data)
	if err != nil {
		return "nil", 7
	} else {
		return GetAccountToken(Account, CarID, 3), 0
	}
}

// ===============================================================================
// ============================================[AccountHas]
func DBAccountHas(ID string) bool {
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
	Collection := Database.C("Register")
	result := bson.M{}
	err = Collection.Find(bson.M{"Accountid": ID}).One(&result)
	if err == nil {
		return true
	} else {
		return false
	}
}
