package MongoDB

import (
	// "fmt"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var TokenBox map[string]map[string]map[string]string

const AccountSetTime int64 = 60 * 60 * 24
const TemporarilySetTime int64 = 30

func GetAccountToken(Account string, others string, Mode int) string {
	// ========================================
	if TokenBox == nil {
		TokenBox = make(map[string]map[string]map[string]string)
	}
	if TokenBox[Account] == nil {
		TokenBox[Account] = make(map[string]map[string]string)
	}
	// ========================================
	Token := ""
	if Mode == 1 {
		rand.Seed(time.Now().UTC().UnixNano())
		Token = GetSHAString(strconv.FormatInt(int64(rand.Intn(4294967295)), 10) + strconv.FormatInt(time.Now().Unix(), 10) + Account + strconv.FormatFloat(float64(rand.Intn(4294967295))+rand.Float64(), 'f', 10, 64))
		TokenBox[Account][Token] = make(map[string]string)
		TokenBox[Account][Token]["Level"] = others // This is Level number
		TokenBox[Account][Token]["Token"] = Token
		TokenBox[Account][Token]["Time"] = strconv.FormatInt(time.Now().Unix(), 10)
		return Token
	} else if Mode == 2 {
		rand.Seed(time.Now().UTC().UnixNano())
		Token = fmt.Sprintf("%06d", rand.Intn(999999))
		TokenBox[Account][Token] = make(map[string]string)
		TokenBox[Account][Token]["Token"] = Token
		TokenBox[Account][Token]["Time"] = strconv.FormatInt(time.Now().Unix(), 10)
		return Token
	} else if Mode == 3 {
		Token = GetSHAString(Account + time.Now().String() + strconv.FormatFloat(float64(rand.Intn(4294967295))+rand.Float64(), 'f', 10, 64))
		TokenBox[Account]["CarToken"] = make(map[string]string)
		TokenBox[Account]["CarToken"][Token] = others // This is CarID
		return Token
	} else {
		return ""
	}
	// ========================================
}

func TokenCheck(Account string, Token string, Mode int) bool {
	if TokenBox == nil || TokenBox[Account] == nil || (len(Token) < 10 && Mode == 1) || TokenBox[Account][Token] == nil || CheckTime(Account, Token, Mode) == false {
		return false
	} else if Mode == 2 {
		TokenInvalid(Account, Token, 1)
		return true
	} else {
		return true
	}
}

func TokenInvalid(Account string, Token string, Mode int) bool {
	_, AccountHas := TokenBox[Account]
	_, TokenHas := TokenBox[Account][Token]
	if TokenHas && Mode == 1 {
		delete(TokenBox[Account], Token)
		return true
	} else if AccountHas && Mode == 2 {
		delete(TokenBox, Account)
		return true
	} else if Mode == 3 && TokenHas {
		return true
	} else {
		return false
	}
}

func CheckTime(Account string, Token string, Mode int) bool {
	Times, _ := strconv.ParseInt(TokenBox[Account][Token]["Time"], 10, 64)
	// ========================================
	var SetTime int64
	if Mode == 1 {
		SetTime = AccountSetTime
	} else if Mode == 2 {
		SetTime = TemporarilySetTime
	}
	// ========================================
	if Times+SetTime < time.Now().Unix() {
		TokenInvalid(Account, Token, 1)
		return false
	} else {
		return true
	}
}
