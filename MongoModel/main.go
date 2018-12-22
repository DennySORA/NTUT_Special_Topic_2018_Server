package MongoModel

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
	return Base.CreateReturn{Status: SelfSuccess(1), ID: AccountIDPW.Account}
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
		return Base.LogInToken{Status: SelfErrors(4)}
	} else {
		// -------------------------------------

		// -------------------------------------
		return ConverLogInToken(Account)
	}
}

// func GetUser(Token string) Base.Users {
// 	ID, ok := TokenBox[Token]
// 	Session, err := mgo.Dial(DatabaseURL)
// 	defer Session.Close()
// 	ERRs(err)
// 	Database := Session.DB(DatabaseName)

// 	if ok == false {
// 		return UserERROE()
// 	} else {
// 		result := bson.M{}
// 		Collection := Database.C("Register")
// 		err = Collection.Find(bson.M{"accountid": ID}).One(&result)
// 		ERRs(err)
// 		Collection = Database.C("User")
// 		err = Collection.Find(bson.M{"_id": result["_id"]}).One(&result)
// 		ERRs(err)
// 		return ConvertDBUser(result)
// 	}
// }

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
