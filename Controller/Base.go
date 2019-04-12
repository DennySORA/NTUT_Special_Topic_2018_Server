package Controller

import (
	"regexp"
	"strings"
	"unicode"
)

// ===========================================================[RoutineInspection]
func RoutineInspection(Account string, Other string) (bool, int) {
	HasEmail, _ := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, Account)
	if len(strings.Trim(Account, " ")) == 0 || len(strings.Trim(Other, " ")) == 0 {
		return false, 1
	} else if HasEmail == false {
		return false, 3
	} else {
		return true, 0
	}
}

func PasswordStrong(password string) bool {
	var number, upper, lower bool
	for _, i := range password {
		switch {
		case unicode.IsNumber(i):
			number = true
		case unicode.IsUpper(i):
			upper = true
		case unicode.IsLower(i):
			lower = true
		default:
		}
	}
	return number && upper && lower && len(password) >= 8 && len(password) <= (15)
}
