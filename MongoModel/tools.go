package MongoModel

import (
	"encoding/hex"
	"fmt"
	Base "sega/Base"
	"time"

	"golang.org/x/crypto/sha3"
	"gopkg.in/mgo.v2/bson"
)

// ============================================[]

func UserConvert(User Base.NewAccountUser, ID string) bson.M {
	TempTime := GetUTCTime()
	return bson.M{
		"Email": ID,
		"Car":   []bson.M{},
		"Profile": bson.M{
			"Name":   User.Name,
			"Gender": User.Gender,
			"Phone": bson.M{
				"Country": User.Country,
				"Number":  User.Number,
			},
		},
		"Accesse": bson.M{
			"Certification": false,
			"PermitTime":    TempTime,
			"Level":         1,
			"PermitLog": []bson.M{
				bson.M{
					"Level":     1,
					"Times":     TempTime,
					"Authority": "Auto",
				},
			},
		},
		"SiginHistory":  []bson.M{},
		"LogoutHistory": []bson.M{},
	}
}

func ConvertRegisters(AccountIDPW Base.NewAccountIDPW, Oid bson.ObjectId) bson.M {
	TempTime := GetUTCTime()
	return bson.M{
		"_id":         Oid,
		"Createtime":  TempTime,
		"Refreshtime": TempTime,
		"Accountid":   AccountIDPW.Account,
		"Password":    GetSHAString(AccountIDPW.Password),
		"IsVerify":    false,
	}
}

func ConverLogInToken(Account string) Base.LogInToken {
	return Base.LogInToken{
		Status:       SelfSuccess(2),
		GetTimes:     GetUTCTime(),
		AccountToken: GetLogInToken(Account),
		AccountID:    Account,
	}
}

// ============================================[]

func GetSHAString(Data string) string {
	SHAMain := sha3.New512()
	SHAMain.Write([]byte(Data))
	SHACode := SHAMain.Sum([]byte(Data))
	// fmt.Println(hex.EncodeToString(SHACode[:]))
	SHAString := hex.EncodeToString(SHACode[:])
	return SHAString
}

func GetUTCTime() string {
	t := time.Now()
	local_location, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
	}
	time_UTC := t.In(local_location)
	return time_UTC.String()
}

func ERRs(logs error) {
	if logs != nil {
		fmt.Println(logs)
	}
}
