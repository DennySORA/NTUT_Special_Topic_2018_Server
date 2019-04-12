package MongoDB

import (
	"SORA/Base"
	TokenBox "SORA/Token"

	"gopkg.in/mgo.v2/bson"
)

// ============================================[AddCarID]
func DBAddCarID(Account string, CarName string) (string, int) {
	Token, ok := TokenBox.Token.GetToken(Account, "Car", 0)
	Collection := Base.DBCol.C("User")
	// ========================================
	selects := bson.M{"Email": Account}
	data := bson.M{"$addToSet": bson.M{
		"Car": bson.M{
			"CarID":       Token,
			"CarName":     CarName,
			"RefreshTime": GetUTCTime(),
			"CreateTime":  GetUTCTime(),
		}}}
	// ========================================
	if err := Collection.Update(selects, data); err != nil || !ok {
		return "nil", 7
	} else {
		return Token, 0
	}
}
