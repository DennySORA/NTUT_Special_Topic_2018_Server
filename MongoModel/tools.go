package MongoModel

import (
	"encoding/hex"
	"fmt"
	Base "sega/Base"
	"time"

	"golang.org/x/crypto/sha3"
	"gopkg.in/mgo.v2/bson"
)

// =========================================

var TokenBox map[string]string

func CreateTokenBox() {
	if TokenBox == nil {
		TokenBox = make(map[string]string)
	}
}

// =========================================
func ConvertUser(AccountData Base.NewAccountData) Base.Users {
	NowTimes := GetUTCTime()
	return Base.Users{
		// ------------------------------------
		Car: []Base.Cars{
			Base.Cars{
				CarID:       AccountData.CarID,
				RefreshTime: NowTimes,
				CreateTime:  NowTimes,
			},
		},
		// ------------------------------------
		Profile: Base.Profiles{
			Name:   AccountData.Name,
			Gender: AccountData.Gender,
			Email:  AccountData.Email,
			Phone: Base.Phones{
				Country: AccountData.Country,
				Number:  AccountData.Number,
			},
		},
		// ------------------------------------
		Accesse: Base.Accesses{
			PermitTime: NowTimes,
			Level:      0,
			PermitLog: []Base.PermitLogs{
				Base.PermitLogs{
					Level:     0,
					Times:     NowTimes,
					Authority: "System",
				},
			},
		},
		// ------------------------------------
	}
}

func ConvertDBUser(Data bson.M) Base.Users {
	fmt.Printf("%v\n", Data)
	return Base.Users{
		// 	// ------------------------------------
		// 	Car: []Cars{
		// 		Cars{
		// 			CarID:       AccountData.CarID,
		// 			RefreshTime: NowTimes,
		// 			CreateTime:  NowTimes,
		// 		},
		// 	},
		// 	// ------------------------------------
		// 	Profile: Profiles{
		// 		Name:   AccountData.Name,
		// 		Gender: AccountData.Gender,
		// 		Email:  AccountData.Email,
		// 		Phone: Phones{
		// 			Country: AccountData.Country,
		// 			Number:  AccountData.Number,
		// 		},
		// 	},
		// 	// ------------------------------------
		// 	Accesse: Accesses{
		// 		PermitTime: NowTimes,
		// 		Level:      0,
		// 		PermitLog: []PermitLogs{
		// 			PermitLogs{
		// 				Level:     0,
		// 				Times:     NowTimes,
		// 				Authority: "System",
		// 			},
		// 		},
		// 	},
	}
}

func ConvertRegisters(AccountIDPW Base.NewAccountIDPW, Oid bson.ObjectId) bson.M {
	NowTimes := GetUTCTime()
	return bson.M{
		"_id":         Oid,
		"createtime":  NowTimes,
		"refreshtime": NowTimes,
		"accountid":   AccountIDPW.Account,
		"password":    AccountIDPW.Password,
		"isVerify":    false,
	}
}

func ConvertCreateReturn(AccountData Base.NewAccountData, Code int, Description string) Base.CreateReturn {
	return Base.CreateReturn{
		Status: Base.StatusData{
			StatusCode:  Code,
			Description: Description,
		},
		Account: AccountData.Email,
		Name:    AccountData.Name,
	}
}

func ConverLogInToken(ID string) Base.LogInToken {
	CreateTokenBox()
	NowTimes := GetUTCTime()
	Token := GetSHAString(ID + NowTimes + "4das56sa4d1")
	TokenBox[Token] = ID
	return Base.LogInToken{
		Status:       Base.StatusData{},
		GetTimes:     NowTimes,
		AccountToken: Token,
		AccountID:    ID,
	}
}

// ========================================================
func GetSHAString(Data string) string {
	SHAMain := sha3.New512()
	SHAMain.Write([]byte(Data))
	SHACode := SHAMain.Sum([]byte(Data))
	// fmt.Println(hex.EncodeToString(bs[:]))
	SHAString := hex.EncodeToString(SHACode[:])
	return SHAString
}

// ========================================================

func LogInERROR() Base.LogInToken {
	return Base.LogInToken{
		Status: Base.StatusData{
			StatusCode:  -1,
			Description: "Account error",
		},
	}
}

func CreateERROR() Base.CreateReturn {
	return Base.CreateReturn{
		Status: Base.StatusData{
			StatusCode:  -1,
			Description: "Account exists ",
		},
	}
}

func UserERROE() Base.Users {
	return Base.Users{
		Status: Base.StatusData{
			StatusCode:  -1,
			Description: "Token nonexists",
		},
	}
}

// ========================================================
// ========================================================
// ========================================================

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
