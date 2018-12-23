package MongoDB

import (
	"fmt"
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
	fmt.Printf("%v\n", User)
	// -------------------------------------
	Collection = Database.C("User")
	err = Collection.Insert(User)
	ERRs(err)
	// -------------------------------------
	result := bson.M{}
	err = Collection.Find(bson.M{"Email": AccountIDPW.Account}).One(&result)
	ERRs(err)
	fmt.Printf("%v\n", result["_id"])
	Oid, _ := result["_id"].(bson.ObjectId)
	// -------------------------------------
	Register := ConvertRegisters(AccountIDPW, Oid)
	Collection = Database.C("Register")
	err = Collection.Insert(Register)
	ERRs(err)
	// -------------------------------------
	return Base.CreateReturn{Status: Base.SelfSuccess(1), ID: AccountIDPW.Account}
}

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
		// -------------------------------------

		// -------------------------------------
		return ConverLogInToken(Account, 1)
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
