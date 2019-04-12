package Controller

import (
	"SORA/Base"
	"SORA/MongoDB"
	TokenBox "SORA/Token"
)

// ============================================[GetUser]

func ExaminationGetUser(Certification Base.InputCertification, GetHistorysNumber int) (*Base.Users, error) {
	Ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if Ok == false {
		return &Base.Users{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Certification.Account) == false {
		return &Base.Users{Status: Base.SelfErrors(5)}, nil
	} else if TokenBox.Token.EqualToekn(Certification.Token, "Account") == false {
		return &Base.Users{Status: Base.SelfErrors(6)}, nil
	}
	Data, Status := MongoDB.DBGetUser(Certification.Account, Certification.Token, GetHistorysNumber)
	if Status != 0 {
		return &Base.Users{Status: Base.SelfErrors(Status)}, nil
	} else {
		return &Data, nil
	}
}

func ExaminationUpdateUser(Certification Base.InputCertification, User Base.NewAccountUser) (*Base.CreateReturn, error) {
	ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if ok == false {
		return &Base.CreateReturn{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Certification.Account) == false {
		return &Base.CreateReturn{Status: Base.SelfErrors(5)}, nil
	} else if TokenBox.Token.EqualToekn(Certification.Token, "Account") == false {
		return &Base.CreateReturn{Status: Base.SelfErrors(6)}, nil
	} else {
		return MongoDB.DBUpdateUser(Certification.Account, User), nil
	}
}
