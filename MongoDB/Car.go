package MongoDB

import (
	"SORA/Base"

	"gopkg.in/mgo.v2/bson"
)

// ============================================[AddCarID]
func DBAddCarID(Account string, CarID string, CarName string) (string, int) {
	Collection := Base.DBCol.C("User")
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
	if err := Collection.Update(selects, data); err != nil {
		return "nil", 7
	} else {
		return GetAccountToken(Account, CarID, 3), 0
	}
}
