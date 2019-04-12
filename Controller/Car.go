package Controller

import (
	"SORA/Base"
	"SORA/Config"
	"SORA/MongoDB"
	TokenBox "SORA/Token"
	"strings"
)

// ============================================[GetCarID]

func ExaminationGetCarID(token string) ([]Base.CarData, error) {
	var accountID string
	var ok bool
	if len(strings.Trim(token, " ")) == 0 {
		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(1)}}, nil
	} else if accountID, ok = TokenBox.Token.EqualToeknGetAccount(token, "Account"); !ok {
		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(6)}}, nil
	}
	Data, Status := MongoDB.DBGetCarID(accountID)
	if Status != 0 {
		return []Base.CarData{Base.CarData{Status: Base.SelfErrors(Status)}}, nil
	} else {
		return Data, nil
	}
}

// ============================================[GetTemporarilyToken]
func ExaminationGetTemporarilyToken(token string) (*Base.TemporarilyTokenData, error) {
	if len(strings.Trim(token, " ")) == 0 {
		return &Base.TemporarilyTokenData{Status: Base.SelfErrors(1)}, nil
	} else if accountID, ok := TokenBox.Token.EqualToeknGetAccount(token, "Account"); !ok {
		return &Base.TemporarilyTokenData{Status: Base.SelfErrors(6)}, nil
	} else if tokens, ok := TokenBox.Token.GetToken(accountID, "Certification", Config.TokenTime); !ok {
		return &Base.TemporarilyTokenData{Status: Base.SelfErrors(9)}, nil
	} else {
		return &Base.TemporarilyTokenData{
			Status:   Base.SelfSuccess(3),
			Token:    tokens,
			GetTimes: MongoDB.GetUTCTime()}, nil
	}
}

// ============================================[AddCarID]
func ExaminationAddCarID(accountID string, carName string, temporarilyToken string) (*Base.CarIDReturn, error) {
	Ok, number := RoutineInspection(accountID, temporarilyToken)
	if Ok == false {
		return &Base.CarIDReturn{Status: Base.SelfErrors(number)}, nil
	} else if TokenBox.Token.EqualToekn(temporarilyToken, "Certification") == false {
		return &Base.CarIDReturn{Status: Base.SelfErrors(6)}, nil
	} else {
		Token, Status := MongoDB.DBAddCarID(accountID, carName)
		if Status != 0 {
			return &Base.CarIDReturn{Status: Base.SelfErrors(Status)}, nil
		} else {
			return &Base.CarIDReturn{
				Status:    Base.SelfSuccess(4),
				AccountID: accountID,
				CarToken:  Token,
			}, nil
		}
	}
}
