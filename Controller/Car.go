package Controller

import (
	"SORA/Project/Go_Back_End_SEGA_Project/Base"
	"SORA/Project/Go_Back_End_SEGA_Project/MongoDB"
)

// ============================================[GetCarID]

func ExaminationGetCarID(Certification Base.InputCertification) ([]Base.CarData, error) {
	Ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if Ok == false {
		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(number)}}, nil
	} else if MongoDB.DBAccountHas(Certification.Account) == false {
		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(5)}}, nil
	}
	Data, Status := MongoDB.DBGetCarID(Certification.Account, Certification.Token)
	if Status != 0 {
		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(Status)}}, nil
	} else {
		return Data, nil
	}
}

// ============================================[GetTemporarilyToken]
func ExaminationGetTemporarilyToken(Certification Base.InputCertification) (Base.TemporarilyTokenData, error) {
	Ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if Ok == false {
		return Base.TemporarilyTokenData{Status: Base.SelfErrors(number)}, nil
	} else if MongoDB.TokenCheck(Certification.Account, Certification.Token, 1) == false {
		return Base.TemporarilyTokenData{Status: Base.SelfErrors(6)}, nil
	} else {
		return Base.TemporarilyTokenData{
			Status:   Base.SelfSuccess(3),
			Token:    MongoDB.GetAccountToken(Certification.Account, "", 2),
			GetTimes: MongoDB.GetUTCTime()}, nil
	}
}

// ============================================[AddCarID]
func ExaminationAddCarID(InputCarNews Base.CarNews) (Base.CarIDReturn, error) {
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
