package MongoDB

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var TokenBox map[string]map[string]map[string]string
var SetTime int64 = 60 * 60 * 24

func GetAccountToken(Account string, Level int, Mode int) string {
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
		Token = GetSHAString(Account + time.Now().String() + strconv.FormatFloat(float64(rand.Intn(4294967295))+rand.Float64(), 'f', 10, 64))
		TokenBox[Account][Token] = make(map[string]string)
		TokenBox[Account][Token]["Level"] = strconv.Itoa(Level)
	} else if Mode == 2 {
		Token = fmt.Sprintf("%06d", rand.Intn(999999))
		TokenBox[Account][Token] = make(map[string]string)
	}
	// ========================================
	TokenBox[Account][Token]["Token"] = Token
	TokenBox[Account][Token]["Time"] = strconv.FormatInt(time.Now().Unix(), 10)
	return Token
}

func TokenCheck(Account string, Token string) bool {
	if TokenBox == nil || TokenBox[Account] == nil || len(Token) < 10 || TokenBox[Account][Token] == nil || CheckTime(Account, Token) == false {
		return false
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
	} else {
		return false
	}
}

func CheckTime(Account string, Token string) bool {
	Times, _ := strconv.ParseInt(TokenBox[Account][Token]["Time"], 10, 64)
	if Times+SetTime < time.Now().Unix() {
		TokenInvalid(Account, Token, 1)
		return false
	} else {
		return true
	}
}
