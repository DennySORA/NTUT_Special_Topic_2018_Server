package Controller

import (
	"SORA/Base"
	"SORA/MongoDB"
	TokenBox "SORA/Token"
	"strings"
)

// ============================================[CreateAccount]

func ExaminationCreateAccount(AccountIDPW Base.NewAccountIDPw, User Base.NewAccountUser) (*Base.CreateReturn, error) {
	Ok, number := RoutineInspection(AccountIDPW.AccountID, AccountIDPW.Password)
	if Ok == false {
		return &Base.CreateReturn{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(AccountIDPW.AccountID) == true {
		return &Base.CreateReturn{Status: Base.SelfErrors(2)}, nil
	} else {
		return MongoDB.DBCreateAccount(AccountIDPW, User), nil
	}
}

// ============================================[LogIn]

func ExaminationLogIn(accountID string, password string, information Base.Logformation) (*Base.LogInToken, error) {
	Ok, number := RoutineInspection(accountID, password)
	if Ok == false {
		return &Base.LogInToken{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(accountID) == false {
		return &Base.LogInToken{Status: Base.SelfErrors(5)}, nil
	} else {
		return MongoDB.DBLogIn(accountID, password, information), nil
	}
}

// ============================================[LogOut]

func ExaminationLogOut(token string, Information Base.Logformation) (*Base.StatusData, error) {
	if len(strings.Trim(token, " ")) == 0 {
		returnData := Base.SelfErrors(1)
		return &returnData, nil
	} else if accountID, ok := TokenBox.Token.EqualToeknGetAccount(token, "Account"); !ok {
		returnData := Base.SelfErrors(6)
		return &returnData, nil
	} else {
		return MongoDB.DBLogOut(accountID, token, Information), nil
	}
}

// ============================================[CheckAccountHas]

func ExaminationCheckAccountHas(accountID string) (*Base.AccountHas, error) {
	Ok, number := RoutineInspection(accountID, "Temp")
	if Ok == false {
		return &Base.AccountHas{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(accountID) == false {
		return &Base.AccountHas{Status: Base.SelfSuccess(6), Has: false}, nil
	} else {
		return &Base.AccountHas{Status: Base.SelfSuccess(6), Has: true}, nil
	}
}
