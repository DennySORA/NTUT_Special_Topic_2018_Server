package Controller

import (
	"SORA/Base"
	"SORA/MongoDB"
	TokenBox "SORA/Token"
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
func ExaminationGetTemporarilyToken(Certification Base.InputCertification) (*Base.TemporarilyTokenData, error) {
	Ok, number := RoutineInspection(Certification.Account, Certification.Token)
	if Ok == false {
		return &Base.TemporarilyTokenData{Status: Base.SelfErrors(number)}, nil
	} else if TokenBox.Token.EqualToekn(Certification.Token, "Account") == false {
		return &Base.TemporarilyTokenData{Status: Base.SelfErrors(6)}, nil
	} else {
		if token, ok := TokenBox.Token.GetToken(Certification.Account, "Certification", 30); !ok {
			return &Base.TemporarilyTokenData{Status: Base.SelfErrors(9)}, nil
		} else {
			return &Base.TemporarilyTokenData{
				Status:   Base.SelfSuccess(3),
				Token:    token,
				GetTimes: MongoDB.GetUTCTime()}, nil
		}
	}
}

// ============================================[AddCarID]
func ExaminationAddCarID(InputCarNews Base.CarNews) (*Base.CarIDReturn, error) {
	Ok, number := RoutineInspection(InputCarNews.ID, InputCarNews.TemporarilyToken)
	if Ok == false {
		return &Base.CarIDReturn{Status: Base.SelfErrors(number)}, nil
	} else if TokenBox.Token.EqualToekn(InputCarNews.TemporarilyToken, "Certification") == false {
		return &Base.CarIDReturn{Status: Base.SelfErrors(6)}, nil
	} else {
		Token, Status := MongoDB.DBAddCarID(InputCarNews.ID, InputCarNews.CarID, InputCarNews.CarName)
		if Status != 0 {
			return &Base.CarIDReturn{Status: Base.SelfErrors(Status)}, nil
		} else {
			return &Base.CarIDReturn{
				Status: Base.SelfSuccess(4),
				ID:     InputCarNews.ID,
				CarID:  InputCarNews.CarID,
				Token:  Token,
			}, nil
		}
	}
}
