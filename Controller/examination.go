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

// ============================================[GetUser]

func ExaminationGetUser(Account string, Token string) (Base.Users, error) {
	Ok, number := RoutineInspection(Account, Token)
	if Ok == false {
		return Base.Users{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Account) == false {
		return Base.Users{Status: Base.SelfErrors(5)}, nil
	} else if MongoDB.TokenCheck(Account, Token, 1) == false {
		return Base.Users{Status: Base.SelfErrors(6)}, nil
	}
	Data, Status := MongoDB.DBGetUser(Account, Token)
	if Status != 0 {
		return Base.Users{Status: Base.SelfErrors(Status)}, nil
	} else {
		return Data, nil
	}
}

// ===========================================================[CarID]
// ============================================[GetCarID]
// func ExaminationGetCarID(Account string, Token string) ([]Base.CarData, error) {
// 	Ok, number := RoutineInspection(Account, Token)
// 	if Ok == false {
// 		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(number)}}, nil
// 	} else if MongoDB.DBAccountHas(Account) == false {
// 		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(5)}}, nil
// 	} else {
// 	}
// }

// ============================================[GetTemporarilyToken]
func ExaminationGetTemporarilyToken(Account string, Token string) (Base.TemporarilyTokenData, error) {
	Ok, number := RoutineInspection(Account, Token)
	if Ok == false {
		return Base.TemporarilyTokenData{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.TokenCheck(Account, Token, 1) == false {
		return Base.TemporarilyTokenData{Status: Base.SelfErrors(6)}, nil
	} else {
		return Base.TemporarilyTokenData{
			Status:   Base.SelfSuccess(3),
			Token:    MongoDB.GetAccountToken(Account, "", 2),
			GetTimes: MongoDB.GetUTCTime()}, nil
	}
}

// ============================================[AddCarID]
func ExaminationAddCarID(Token Base.InputToken, InputCarNews Base.CarNews) (Base.CarIDReturn, error) {
	Ok, number := RoutineInspection(InputCarNews.ID, InputCarNews.TemporarilyToken)
	if Ok == false {
		return Base.CarIDReturn{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.TokenCheck(InputCarNews.ID, InputCarNews.TemporarilyToken, 2) == false {
		return Base.CarIDReturn{Status: Base.SelfErrors(6)}, nil
	} else {
		Token, Status := MongoDB.DBAddCarID(InputCarNews.ID, InputCarNews.CarID, InputCarNews.CarName)
		if Status != 0 {
			return Base.CarIDReturn{Status: Base.SelfErrors(Status)}, nil
		} else {
			return Base.CarIDReturn{
				Status: Base.SelfSuccess(4),
				ID:     InputCarNews.ID,
				CarID:  InputCarNews.CarID,
				Token:  Token,
			}, nil
		}
	}
}

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
