package Controller

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Base"
	"SORA/Project/Go_Back_End_SEGA_Project/MongoDB"
)

// ============================================[GetUser]

func ExaminationGetUser(Certification Base.InputCertification) (Base.Users, error) {
	Ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if Ok == false {
		return Base.Users{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Certification.Account) == false {
		return Base.Users{Status: Base.SelfErrors(5)}, nil
	} else if MongoDB.TokenCheck(Certification.Account, Certification.Token, 1) == false {
		return Base.Users{Status: Base.SelfErrors(6)}, nil
	}
	Data, Status := MongoDB.DBGetUser(Certification.Account, Certification.Token)
	if Status != 0 {
		return Base.Users{Status: Base.SelfErrors(Status)}, nil
	} else {
		return Data, nil
	}
}

func ExaminationUpdateUser(Certification Base.InputCertification, User Base.NewAccountUser) (Base.CreateReturn, error) {
	ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if ok == false {
		return Base.CreateReturn{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.DBAccountHas(Certification.Account) == false {
		return Base.CreateReturn{Status: Base.SelfErrors(5)}, nil
	} else if MongoDB.TokenCheck(Certification.Account, Certification.Token, 1) == false {
		return Base.CreateReturn{Status: Base.SelfErrors(6)}, nil
	} else {
		return MongoDB.DBUpdateUser(Certification.Account, User), nil
	}
}
