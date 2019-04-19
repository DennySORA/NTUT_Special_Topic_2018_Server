package Base

import (
	"SORA/Config"
	TokenBox "SORA/Token"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	DBCol *mgo.Database
)

func InitMongoDB() error {
	if Session, err := mgo.Dial(Config.DatabaseURL); err != nil {
		return err
	} else {
		DBCol = Session.DB(Config.DatabaseName)
	}
	return nil
}

func InitCarToken() error {
	results := []bson.M{}
	Collection := DBCol.C("User")
	Collection.Find(bson.M{}).Select(bson.M{"Email": 1, "Car.CarID": 1, "Car.CreateTime": 1}).All(&results)
	for _, val := range results {
		accuntID := val["Email"]
		for _, Car := range val["Car"].([]interface{}) {
			tokenData := TokenBox.Tokens{
				GetTime:     timeConvert(Car.(bson.M)["CreateTime"].(string)),
				Account:     accuntID.(string),
				InvalidTime: -1,
			}
			TokenBox.Token.CarToken[Car.(bson.M)["CarID"].(string)] = tokenData
		}
	}
	return nil
}

func timeConvert(times string) int64 {
	layout := "2006-01-02 15:04:05.0000000 -0700 MST"
	t, err := time.Parse(layout, times)
	if err != nil {
		return 0
	} else {
		return t.Unix()
	}
}
