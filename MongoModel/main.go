package MongoModel

import (
	"fmt"
	Base "sega/Base"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DatabaseURL = "127.0.0.1:27017"
const DatabaseName = "Test"

// =========================================

func CreateAccount(AccountIDPW Base.NewAccountIDPW, AccountData Base.NewAccountData) Base.CreateReturn {
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
	Collection := Database.C("Register")
	result := bson.M{}
	// =====================================
	err = Collection.Find(bson.M{"accountid": AccountIDPW.Account}).One(&result)
	fmt.Println(err, result, err == nil, result != nil)
	if result == nil || err == nil {
		return CreateERROR()
	}
	// -------------------------------------
	User := ConvertUser(AccountData)
	// -------------------------------------
	Collection = Database.C("User")
	err = Collection.Insert(User)
	ERRs(err)
	// -------------------------------------
	result = bson.M{}
	err = Collection.Find(bson.M{"profile.email": AccountData.Email}).One(&result)
	ERRs(err)
	fmt.Printf("%v\n", result["_id"])
	Oid, _ := result["_id"].(bson.ObjectId)
	// -------------------------------------
	Register := ConvertRegisters(AccountIDPW, Oid)
	Collection = Database.C("Register")
	err = Collection.Insert(Register)
	ERRs(err)
	// -------------------------------------
	CreateReturn := ConvertCreateReturn(AccountData, 0, "success")
	return CreateReturn
}

func LongIn(ID string, Password string) Base.LogInToken {
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
	Collection := Database.C("Register")
	result := bson.M{}
	err = Collection.Find(bson.M{"accountid": ID, "password": Password}).One(&result)
	if result == nil || err != nil {
		return LogInERROR()
	} else {
		return ConverLogInToken(ID)
	}
}

func AccountHas(ID string) bool {
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)
	Collection := Database.C("Register")
	result := bson.M{}
	err = Collection.Find(bson.M{"accountid": ID}).One(&result)
	if result == nil || err == nil {
		return true
	} else {
		return false
	}
}

func GetUser(Token string) Base.Users {
	ID, ok := TokenBox[Token]
	Session, err := mgo.Dial(DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(DatabaseName)

	if ok == false {
		return UserERROE()
	} else {
		result := bson.M{}
		Collection := Database.C("Register")
		err = Collection.Find(bson.M{"accountid": ID}).One(&result)
		ERRs(err)
		Collection = Database.C("User")
		err = Collection.Find(bson.M{"_id": result["_id"]}).One(&result)
		ERRs(err)
		return ConvertDBUser(result)
	}
}
