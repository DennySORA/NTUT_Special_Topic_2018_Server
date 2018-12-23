package MongoDB

import (
	"encoding/hex"
	"fmt"
	Base "sega/Base"
	"strconv"
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

func ConverLogInToken(Account string, level int) Base.LogInToken {
	return Base.LogInToken{
		Status:       Base.SelfSuccess(2),
		GetTimes:     GetUTCTime(),
		AccountToken: GetAccountToken(Account, strconv.Itoa(level), 1),
		AccountID:    Account,
	}
}

func ReturnUserConvert(UserData bson.M) Base.Users {
	return Base.Users{
		Status: Base.SelfSuccess(5),
		Car:    ForCarList(UserData["Car"].([]interface{})),
		Profile: Base.Profiles{
			Name:   UserData["Profile"].(bson.M)["Name"].(string),
			Gender: UserData["Profile"].(bson.M)["Gender"].(int),
			Phone: Base.Phones{
				Country: UserData["Profile"].(bson.M)["Phone"].(bson.M)["Country"].(string),
				Number:  UserData["Profile"].(bson.M)["Phone"].(bson.M)["Number"].(string),
			},
		},
		Accesse: Base.Accesses{
			PermitTime: UserData["Accesse"].(bson.M)["PermitTime"].(string),
			Level:      UserData["Accesse"].(bson.M)["Level"].(int),
			PermitLog:  ForPermitLogList(UserData["Accesse"].(bson.M)["PermitLog"].([]interface{})),
		},
		SiginHistory:  ForSiginHistoryList(UserData["SiginHistory"].([]interface{})),
		LogoutHistory: ForLogoutHistoryList(UserData["LogoutHistory"].([]interface{})),
	}
}

func ForSiginHistoryList(SiginHistory []interface{}) Base.Historys {
	TempList := []Base.Historys{}
	for _, v := range SiginHistory {
		Temp, _ := strconv.Atoi(v.(bson.M)["Types"].(string))
		TempList = append(TempList, Base.Historys{
			UseToken: v.(bson.M)["UseToken"].(string),
			Times:    v.(bson.M)["Times"].(string),
			Types:    Temp,
			Device:   v.(bson.M)["Device"].(string),
		})
	}
	if len(TempList) == 0 {
		return Base.Historys{}
	} else {
		return TempList[len(TempList)-1]
	}
}

func ForLogoutHistoryList(LogoutHistory []interface{}) Base.Historys {
	TempList := []Base.Historys{}
	for _, v := range LogoutHistory {
		Temp, _ := strconv.Atoi(v.(bson.M)["Types"].(string))
		TempList = append(TempList, Base.Historys{
			UseToken: v.(bson.M)["UseToken"].(string),
			Times:    v.(bson.M)["Times"].(string),
			Types:    Temp,
			Device:   v.(bson.M)["Device"].(string),
		})
	}
	if len(TempList) == 0 {
		return Base.Historys{}
	} else {
		return TempList[len(TempList)-1]
	}
}

func ForPermitLogList(PermitLogs []interface{}) []*Base.PermitLogs {
	TempList := []Base.PermitLogs{}
	for _, v := range PermitLogs {
		TempList = append(TempList, Base.PermitLogs{
			Level:     v.(bson.M)["Level"].(int),
			Times:     v.(bson.M)["Times"].(string),
			Authority: v.(bson.M)["Authority"].(string),
		})
	}
	PointList := []*Base.PermitLogs{}
	for i := 0; i < len(TempList); i++ {
		PointList = append(PointList, &TempList[i])
	}
	return PointList
}

func ForCarList(CarList []interface{}) []Base.CarData {
	TempList := []Base.CarData{}
	for _, v := range CarList {
		RefreshTimeTemp := v.(bson.M)["RefreshTime"].(string)
		CreateTimeTemp := v.(bson.M)["CreateTime"].(string)
		RefreshTime := &RefreshTimeTemp
		CreateTime := &CreateTimeTemp
		TempList = append(TempList, Base.CarData{
			CarName:     v.(bson.M)["CarName"].(string),
			CarID:       v.(bson.M)["CarID"].(string),
			RefreshTime: RefreshTime,
			CreateTime:  CreateTime,
		})
	}
	return TempList
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
