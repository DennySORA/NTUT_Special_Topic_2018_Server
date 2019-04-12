package Base

import (
	"SORA/Config"

	"gopkg.in/mgo.v2"
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
