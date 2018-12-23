package Controller

import (
	regexp "regexp"
	Base "sega/Base"
	MongoDB "sega/MongoDB"
)

// ===============================================================================
// ===========================================================[Account]
// ============================================[CreateAccount]
func ExaminationCreateAccount(AccountIDPW Base.NewAccountIDPW, User Base.NewAccountUser) (Base.CreateReturn, error) {
	Ok, number := RoutineInspection(AccountIDPW.Account, AccountIDPW.Password)
	if Ok == false {
		return Base.CreateReturn{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(AccountIDPW.Account) == true {
		return Base.CreateReturn{Status: Base.SelfErrors(2)}, nil
	} else {
		return MongoDB.DBCreateAccount(AccountIDPW, User), nil
	}
}

// ============================================[LogIn]
func ExaminationLogIn(Account string, Password string) (Base.LogInToken, error) {
	Ok, number := RoutineInspection(Account, Password)
	if Ok == false {
		return Base.LogInToken{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Account) == false {
		return Base.LogInToken{Status: Base.SelfErrors(5)}, nil
	} else {
		return MongoDB.DBLogIn(Account, Password), nil
	}
}

// ===========================================================[]
// ============================================[GetTemporarilyToken]

func ExaminationGetTemporarilyToken(Account string, Token string) (Base.TemporarilyTokenData, error) {
	Ok, number := RoutineInspection(Account, Token)
	if Ok == false {
		return Base.TemporarilyTokenData{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.TokenCheck(Account, Token) == false {
		return Base.TemporarilyTokenData{Status: Base.SelfErrors(6)}, nil
	} else {
		return Base.TemporarilyTokenData{
			Status:   Base.SelfSuccess(3),
			Token:    MongoDB.GetAccountToken(Account, 0, 2),
			GetTimes: MongoDB.GetUTCTime()}, nil
	}
}

// func ExaminationAddCarID(Token Base.InputToken, InputCarNews Base.CarNews) (Base.CarIDReturn, error) {

// }

// ===============================================================================
// ===========================================================[RoutineInspection]
func RoutineInspection(Account string, Other string) (bool, int) {
	HasEmail, _ := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, Account)
	if Account == "" || Other == "" {
		return false, 1
	} else if HasEmail == false {
		return false, 3
	} else {
		return true, 0
	}
}
