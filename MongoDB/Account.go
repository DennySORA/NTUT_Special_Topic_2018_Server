package MongoDB

import (
	"SORA/Base"
	TokenBox "SORA/Token"

	"gopkg.in/mgo.v2/bson"
)

// ============================================[CreateAccount]

func DBCreateAccount(AccountIDPW Base.NewAccountIDPw, AccountData Base.NewAccountUser) *Base.CreateReturn {
	result := bson.M{}
	// ====================================[DBLink]
	Collection := Base.DBCol.C("Register")
	// -------------------------------------
	User := UserConvert(AccountData, AccountIDPW.AccountID)
	// -------------------------------------
	Collection = Base.DBCol.C("User")
	if err := Collection.Insert(User); err != nil {
		Base.Error.Println(err)
		return &Base.CreateReturn{Status: Base.SelfErrors(-1)}
	} else if err = Collection.Find(bson.M{"Email": AccountIDPW.AccountID}).One(&result); err != nil {
		Base.Error.Println(err)
		return &Base.CreateReturn{Status: Base.SelfErrors(-1)}
	}
	// -------------------------------------
	Oid, _ := result["_id"].(bson.ObjectId)
	Register := ConvertRegisters(AccountIDPW, Oid)
	Collection = Base.DBCol.C("Register")
	// -------------------------------------
	if err := Collection.Insert(Register); err != nil {
		Base.Error.Println(err)
		return &Base.CreateReturn{Status: Base.SelfErrors(-1)}
	} else {
		return &Base.CreateReturn{Status: Base.SelfSuccess(1), AccountID: AccountIDPW.AccountID}
	}
	// -------------------------------------
}

// ============================================[LogIn]

func DBLogIn(Account string, Password string, Information Base.Logformation) *Base.LogInToken {
	result := bson.M{}
	Collection := Base.DBCol.C("Register")
	// -------------------------------------
	if err := Collection.Find(bson.M{"Accountid": Account, "Password": GetSHAString(Password)}).One(&result); err != nil {
		return &Base.LogInToken{Status: Base.SelfErrors(4)}
	} else {
		ReturnData, ok := ConverLogInToken(Account, 1)
		if !ok {
			return &Base.LogInToken{Status: Base.SelfErrors(9)}
		}
		// -------------------------------------
		Collection := Base.DBCol.C("User")
		// ========================================
		selects := bson.M{"Email": Account}
		data := bson.M{"$addToSet": bson.M{
			"SiginHistory": bson.M{
				"Times":    GetUTCTime(),
				"UseToken": ReturnData.AccountToken,
				"Types":    Information.Type,
				"Device":   Information.Device,
			}}}
		// ========================================
		if err := Collection.Update(selects, data); err != nil {
			return &Base.LogInToken{Status: Base.SelfErrors(4)}
		} else {
			return &ReturnData
		}
		// -------------------------------------
	}
}

// ============================================[LogOut]

func DBLogOut(Account string, Token string, Information Base.Logformation) *Base.StatusData {
	Invalid := TokenBox.Token.RemoveToken(Token, "Account")
	if Invalid == false {
		returnData := Base.SelfErrors(8)
		return &returnData
	}
	// ========================================
	Collection := Base.DBCol.C("User")
	// ========================================
	selects := bson.M{"Email": Account}
	data := bson.M{"$addToSet": bson.M{
		"LogoutHistory": bson.M{
			"Times":    GetUTCTime(),
			"UseToken": Token,
			"Types":    Information.Type,
			"Device":   Information.Device,
		}}}
	// ========================================
	if err := Collection.Update(selects, data); err != nil {
		returnData := Base.SelfErrors(0)
		return &returnData
	} else {
		returnData := Base.SelfSuccess(9)
		return &returnData
	}
	// -------------------------------------
}

// ============================================[AccountHas]

func DBAccountHas(ID string) bool {
	result := bson.M{}
	Collection := Base.DBCol.C("Register")
	if err := Collection.Find(bson.M{"Accountid": ID}).One(&result); err == nil {
		return true
	} else {
		return false
	}
}
