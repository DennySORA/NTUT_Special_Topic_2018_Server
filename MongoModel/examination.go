package MongoModel

import (
	regexp "regexp"
	Base "sega/Base"
)

// ===============================================================================
// ===========================================================[Account]
// ============================================[CreateAccount]
func ExaminationCreateAccount(AccountIDPW Base.NewAccountIDPW, User Base.NewAccountUser) (Base.CreateReturn, error) {
	//==========================================================
	HasEmail, _ := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, AccountIDPW.Account)
	//==========================================================
	if AccountIDPW.Account == "" || AccountIDPW.Password == "" {
		return Base.CreateReturn{Status: SelfErrors(1)}, nil
	} else if HasEmail == false {
		return Base.CreateReturn{Status: SelfErrors(3)}, nil
	} else if DBAccountHas(AccountIDPW.Account) == true {
		return Base.CreateReturn{Status: SelfErrors(2)}, nil
	} else {
		return DBCreateAccount(AccountIDPW, User), nil
	}
}

// ============================================[LogIn]
func ExaminationLogIn(Account string, Password string) (Base.LogInToken, error) {
	//==========================================================
	HasEmail, _ := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, Account)
	//==========================================================
	if Account == "" || Password == "" {
		return Base.LogInToken{Status: SelfErrors(1)}, nil
	} else if HasEmail == false {
		return Base.LogInToken{Status: SelfErrors(3)}, nil
	} else if DBAccountHas(Account) == false {
		return Base.LogInToken{Status: SelfErrors(5)}, nil
	} else {
		return DBLogIn(Account, Password), nil
	}
}

// ===============================================================================
// ============================================[DB]
