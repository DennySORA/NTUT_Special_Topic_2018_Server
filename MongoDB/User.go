package MongoDB

import (
	"SORA/Base"
	"SORA/Config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ============================================[GetUser]
func DBGetUser(Account string, Token string) (Base.Users, int) {
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
	Collection := Database.C("User")
	result := bson.M{}
	err = Collection.Find(bson.M{"Email": Account}).One(&result)
	if err != nil {
		return Base.Users{}, 8
	} else {
		return ReturnUserConvert(result), 0
	}
}

// ============================================[GetUser]

func DBGetCarID(Account string, Token string) ([]Base.CarData, int) {
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
	Collection := Database.C("User")
	result := bson.M{}
	err = Collection.Find(bson.M{"Email": Account}).Select(bson.M{"_id": 0, "Car": 1}).One(&result)
	if err != nil {
		return []Base.CarData{}, 8
	} else {
		return ReturnCarIDConvert(result), 0
	}
}

// ============================================[GetUser]

func DBUpdateUser(Account string, AccountData Base.NewAccountUser) *Base.CreateReturn {
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
	Collection := Database.C("User")
	User := UserConvert(AccountData, Account)["Profile"].(bson.M)
	User = bson.M{"$set": bson.M{"Profile": User}}
	err = Collection.Update(bson.M{"Email": Account}, User)
	if err != nil {
		return &Base.CreateReturn{Status: Base.SelfErrors(0)}
	} else {
		return &Base.CreateReturn{Status: Base.SelfSuccess(8), ID: Account}
	}
}
