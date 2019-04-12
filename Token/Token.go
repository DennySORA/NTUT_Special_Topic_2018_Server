package TokenBox

import (
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

const AccountSetTime int64 = 60 * 60 * 24

type TokenSet struct {
	Account            map[string]AccountInformation
	AccountToken       map[string]Tokens
	CarToken           map[string]Tokens
	CertificationToken map[string]Tokens
}

type Tokens struct {
	GetTime     int64
	InvalidTime int64
	Account     string
}

type AccountInformation struct {
	AccountNumber       int
	LimitAccountNumber  int
	CarNumber           int
	CertificationNumber int
}

var Token TokenSet

func init() {
	Token = TokenSet{}
	Token.AccountToken = make(map[string]Tokens)
	Token.CertificationToken = make(map[string]Tokens)
	Token.CarToken = make(map[string]Tokens)
	Token.Account = make(map[string]AccountInformation)
}

func (t *TokenSet) GetToken(account string, types string, setInvalidTime int64) (string, bool) {
	var Token string
	accountInformation := t.Account[account]
	if setInvalidTime == 0 {
		setInvalidTime = AccountSetTime
	}
	// --------------------------------------------------------
	InsertToken := Tokens{
		GetTime:     time.Now().Unix(),
		InvalidTime: setInvalidTime,
		Account:     account,
	}
	tokenBox := t.SwitchToken(types)
	switch types {
	case "Account":
		if accountInformation.LimitAccountNumber != 0 && accountInformation.LimitAccountNumber >= accountInformation.AccountNumber {
			return "", false
		}
		accountInformation.AccountNumber += 1
		Token = uuid.NewV4().String()
	case "Car":
		accountInformation.CarNumber += 1
		Token = uuid.NewV4().String()
	case "Certification":
		accountInformation.CertificationNumber += 1
		rand.Seed(time.Now().UTC().UnixNano())
		Token = fmt.Sprintf("%06d", rand.Intn(999999))
	}
	tokenBox[Token] = InsertToken
	fmt.Println(tokenBox)
	t.CalculationTokenNumber(types, account, 1)
	return Token, true
}
func (t *TokenSet) EqualToekn(token string, types string) bool {
	tokenBox := t.SwitchToken(types)
	information, ok := tokenBox[token]
	fmt.Println(tokenBox)
	// --------------------------------------------------------
	if ok == false {
		return false
	} else if information.InvalidTime != -1 && information.GetTime+information.InvalidTime < time.Now().Unix() {
		t.RemoveToken(token, types)
		return false
	} else {
		return true
	}
}
func (t *TokenSet) RemoveToken(token string, types string) bool {
	tokenBox := t.SwitchToken(types)
	account, ok := tokenBox[token]
	if ok == false {
		return false
	} else {
		t.CalculationTokenNumber(types, account.Account, -1)
		delete(tokenBox, token)
		return true
	}
}

func (t *TokenSet) CalculationTokenNumber(types, account string, number int) {
	Temp := t.Account[account]
	switch types {
	case "Account":
		Temp.AccountNumber += number
	case "Car":
		Temp.CarNumber += number
	case "Certification":
		Temp.CertificationNumber += number
	}
}

func (t *TokenSet) SwitchToken(types string) map[string]Tokens {
	switch types {
	case "Account":
		return t.AccountToken
	case "Car":
		return t.CarToken
	case "Certification":
		return t.CertificationToken
	default:
		return t.AccountToken
	}
}

// package MongoDB

// import (
// 	"fmt"
// 	"math/rand"
// 	"strconv"
// 	"time"

// 	uuid "github.com/satori/go.uuid"
// )

// var TokenBox map[string]map[string]map[string]string

// func GetAccountToken(Account string, others string, Mode int) string {
// 	// ========================================
// 	if TokenBox == nil {
// 		TokenBox = make(map[string]map[string]map[string]string)
// 	}
// 	if TokenBox[Account] == nil {
// 		TokenBox[Account] = make(map[string]map[string]string)
// 	}
// 	// ========================================
// 	Token := ""
// 	if Mode == 1 {
// 		rand.Seed(time.Now().UTC().UnixNano())
// 		Token = uuid.NewV4().String()
// 		TokenBox[Account][Token] = make(map[string]string)
// 		TokenBox[Account][Token]["Level"] = others // This is Level number
// 		TokenBox[Account][Token]["Token"] = Token
// 		TokenBox[Account][Token]["Time"] = strconv.FormatInt(time.Now().Unix(), 10)
// 		return Token
// 	} else if Mode == 2 {
// 		rand.Seed(time.Now().UTC().UnixNano())
// 		Token = fmt.Sprintf("%06d", rand.Intn(999999))
// 		TokenBox[Account][Token] = make(map[string]string)
// 		TokenBox[Account][Token]["Token"] = Token
// 		TokenBox[Account][Token]["Time"] = strconv.FormatInt(time.Now().Unix(), 10)
// 		return Token
// 	} else if Mode == 3 {
// 		Token = uuid.NewV4().String()
// 		TokenBox[Account]["CarToken"] = make(map[string]string)
// 		TokenBox[Account]["CarToken"][Token] = others // This is CarID
// 		return Token
// 	} else {
// 		return ""
// 	}
// 	// ========================================
// }

// func TokenCheck(Account string, Token string, Mode int) bool {
// 	if TokenBox == nil || TokenBox[Account] == nil || (len(Token) < 10 && Mode == 1) || TokenBox[Account][Token] == nil || CheckTime(Account, Token, Mode) == false {
// 		return false
// 	} else if Mode == 2 {
// 		TokenInvalid(Account, Token, 1)
// 		return true
// 	} else {
// 		return true
// 	}
// }

// func TokenInvalid(Account string, Token string, Mode int) bool {
// 	_, AccountHas := TokenBox[Account]
// 	_, TokenHas := TokenBox[Account][Token]
// 	if TokenHas && Mode == 1 {
// 		delete(TokenBox[Account], Token)
// 		return true
// 	} else if AccountHas && Mode == 2 {
// 		delete(TokenBox, Account)
// 		return true
// 	} else if Mode == 3 && TokenHas {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func CheckTime(Account string, Token string, Mode int) bool {
// 	Times, _ := strconv.ParseInt(TokenBox[Account][Token]["Time"], 10, 64)
// 	// ========================================
// 	var SetTime int64
// 	if Mode == 1 {
// 		SetTime = AccountSetTime
// 	} else if Mode == 2 {
// 		SetTime = TemporarilySetTime
// 	}
// 	// ========================================
// 	if Times+SetTime < time.Now().Unix() {
// 		TokenInvalid(Account, Token, 1)
// 		return false
// 	} else {
// 		return true
// 	}
// }
