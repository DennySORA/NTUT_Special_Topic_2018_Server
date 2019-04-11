package MongoDB

import (
	"SORA/Base"

	"gopkg.in/mgo.v2/bson"
)

// ============================================[GetUser]
func DBGetUser(Account string, Token string, GetHistorysNumber int) (Base.Users, int) {
	Collection := Base.DBCol.C("User")
	result := bson.M{}
	if err := Collection.Find(bson.M{"Email": Account}).One(&result); err != nil {
		return Base.Users{}, 8
	} else {
		return ReturnUserConvert(result, GetHistorysNumber), 0
	}
}

// ============================================[GetUser]

func DBGetCarID(Account string, Token string) ([]Base.CarData, int) {
	result := bson.M{}
	Collection := Base.DBCol.C("User")
	if err := Collection.Find(bson.M{"Email": Account}).Select(bson.M{"_id": 0, "Car": 1}).One(&result); err != nil {
		return []Base.CarData{}, 8
	} else {
		return ReturnCarIDConvert(result), 0
	}
}

// ============================================[GetUser]

func DBUpdateUser(Account string, AccountData Base.NewAccountUser) *Base.CreateReturn {
	Collection := Base.DBCol.C("User")
	User := UserConvert(AccountData, Account)["Profile"].(bson.M)
	User = bson.M{"$set": bson.M{"Profile": User}}
	if err := Collection.Update(bson.M{"Email": Account}, User); err != nil {
		return &Base.CreateReturn{Status: Base.SelfErrors(0)}
	} else {
		return &Base.CreateReturn{Status: Base.SelfSuccess(8), ID: Account}
	}
}
