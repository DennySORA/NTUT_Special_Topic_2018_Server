package MongoDB

import (
	"SORA/Config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ============================================[AddCarID]
func DBAddCarID(Account string, CarID string, CarName string) (string, int) {
	Session, err := mgo.Dial(Config.DatabaseURL)
	defer Session.Close()
	ERRs(err)
	Database := Session.DB(Config.DatabaseName)
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
