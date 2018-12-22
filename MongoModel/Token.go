package MongoModel

import (
	"math/rand"
	"strconv"
	"time"
)

var TokenBox map[string]string

func CreateTokenBox() {
	if TokenBox == nil {
		TokenBox = make(map[string]string)
	}
}

func GetLogInToken(Account string) string {
	CreateTokenBox()
	Token := GetSHAString(Account + time.Now().String() + strconv.FormatFloat(float64(rand.Intn(4294967295))+rand.Float64(), 'f', 10, 64))
	TokenBox[Account] = Token
	return Token
}
