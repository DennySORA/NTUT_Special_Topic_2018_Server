package Controller

import (
	"SORA/Base"
	"SORA/MongoDB"
	TokenBox "SORA/Token"
	"strings"
)

// ============================================[GetUser]

func ExaminationGetUser(token string, getHistorysNumber int) (*Base.Users, error) {
	var accountID string
	var ok bool
	if len(strings.Trim(token, " ")) == 0 {
		return &Base.Users{Status: Base.SelfErrors(1)}, nil
	} else if accountID, ok = TokenBox.Token.EqualToeknGetAccount(token, "Account"); !ok {
		return &Base.Users{Status: Base.SelfErrors(6)}, nil
	}
	Data, Status := MongoDB.DBGetUser(accountID, getHistorysNumber)
	if Status != 0 {
		return &Base.Users{Status: Base.SelfErrors(Status)}, nil
	} else {
		return &Data, nil
	}
}

func ExaminationUpdateUser(token string, user Base.NewAccountUser) (*Base.CreateReturn, error) {
	if len(strings.Trim(token, " ")) == 0 {
		return &Base.CreateReturn{Status: Base.SelfErrors(1)}, nil
	} else if accountID, ok := TokenBox.Token.EqualToeknGetAccount(token, "Account"); !ok {
		return &Base.CreateReturn{Status: Base.SelfErrors(6)}, nil
	} else {
		return MongoDB.DBUpdateUser(accountID, user), nil
	}
}
