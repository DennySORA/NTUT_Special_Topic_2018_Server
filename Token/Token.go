package TokenBox

import (
	"SORA/Config"
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

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
	CarNumber           int
	CertificationNumber int
}

var Token TokenSet

func init() {
	go monitorTokenInvalid()
	Token = TokenSet{}
	Token.AccountToken = make(map[string]Tokens)
	Token.CertificationToken = make(map[string]Tokens)
	Token.CarToken = make(map[string]Tokens)
	Token.Account = make(map[string]AccountInformation)
}

func (t *TokenSet) GetToken(account string, types string, setInvalidTime int64) (string, bool) {
	var Token string
	if _, ok := t.Account[account]; !ok {
		t.Account[account] = AccountInformation{}
	}
	if setInvalidTime == 0 {
		setInvalidTime = Config.AccountSetTime
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
		if Config.TokenNumber != 0 && Config.TokenNumber < t.Account[account].AccountNumber+1 {
			return "", false
		}
		Token = uuid.NewV4().String()
	case "Car":
		Token = uuid.NewV4().String()
	case "Certification":
		rand.Seed(time.Now().UTC().UnixNano())
		Token = fmt.Sprintf("%06d", rand.Intn(999999))
	}
	tokenBox[Token] = InsertToken
	t.CalculationTokenNumber(types, account, 1)
	return Token, true
}

func (t *TokenSet) EqualToekn(token string, types string) bool {
	tokenBox := t.SwitchToken(types)
	information, ok := tokenBox[token]
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

func (t *TokenSet) EqualToeknGetAccount(token string, types string) (string, bool) {
	tokenBox := t.SwitchToken(types)
	information, ok := tokenBox[token]
	// --------------------------------------------------------
	if ok == false {
		return "", false
	} else if information.InvalidTime != -1 && information.GetTime+information.InvalidTime < time.Now().Unix() {
		t.RemoveToken(token, types)
		return "", false
	} else {
		return information.Account, true
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
	Temp := AccountInformation{
		AccountNumber:       t.Account[account].AccountNumber,
		CarNumber:           t.Account[account].CarNumber,
		CertificationNumber: t.Account[account].CertificationNumber,
	}
	switch types {
	case "Account":
		Temp.AccountNumber += number
	case "Car":
		Temp.CarNumber += number
	case "Certification":
		Temp.CertificationNumber += number
	}
	t.Account[account] = Temp
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

func monitorTokenInvalid() {
	tick := time.NewTicker(Config.TokenCheckTime)
	for {
		select {
		case <-tick.C:
			go checkTokenInvalid(Token.AccountToken, "Account")
			go checkTokenInvalid(Token.CarToken, "Car")
			go checkTokenInvalid(Token.CertificationToken, "Certification")
		}
	}
}

func checkTokenInvalid(tokens map[string]Tokens, types string) {
	for key, value := range tokens {
		if value.InvalidTime != -1 && value.GetTime+value.InvalidTime < time.Now().Unix() {
			Token.RemoveToken(key, types)
		}
	}
}
