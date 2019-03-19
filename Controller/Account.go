package Controller

import (
	"SORA/Base"
	"SORA/MongoDB"
)

// ============================================[CreateAccount]

func ExaminationCreateAccount(AccountIDPW Base.NewAccountIDPw, User Base.NewAccountUser) (*Base.CreateReturn, error) {
	Ok, number := RoutineInspection(AccountIDPW.Account, AccountIDPW.Password)
	if Ok == false {
		return &Base.CreateReturn{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(AccountIDPW.Account) == true {
		return &Base.CreateReturn{Status: Base.SelfErrors(2)}, nil
	} else {
		return MongoDB.DBCreateAccount(AccountIDPW, User), nil
	}
}

// ============================================[LogIn]

func ExaminationLogIn(Account string, Password string) (*Base.LogInToken, error) {
	Ok, number := RoutineInspection(Account, Password)
	if Ok == false {
		return &Base.LogInToken{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Account) == false {
		return &Base.LogInToken{Status: Base.SelfErrors(5)}, nil
	} else {
		return MongoDB.DBLogIn(Account, Password), nil
	}
}

// ============================================[LogOut]

func ExaminationLogOut(Certification Base.InputCertification) (*Base.StatusData, error) {
	Ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if Ok == false {
		returnData := Base.SelfErrors(number)
		return &returnData, nil
	} else if MongoDB.DBAccountHas(Certification.Account) == false {
		returnData := Base.SelfErrors(5)
		return &returnData, nil
	} else {
		return MongoDB.DBLogOut(Certification.Account, Certification.Token), nil
	}
}

// ============================================[CheckAccountHas]

func ExaminationCheckAccountHas(Account string) (*Base.AccountHas, error) {
	Ok, number := RoutineInspection(Account, "Temp")
	if Ok == false {
		return &Base.AccountHas{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Account) == false {
		return &Base.AccountHas{Status: Base.SelfSuccess(6), Has: false}, nil
	} else {
		return &Base.AccountHas{Status: Base.SelfSuccess(6), Has: true}, nil
	}
}
