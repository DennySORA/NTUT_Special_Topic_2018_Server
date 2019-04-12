package Controller

import (
	"regexp"
	"strings"
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
