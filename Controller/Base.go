package Controller

import "regexp"

// ===========================================================[RoutineInspection]
func RoutineInspection(Account string, Other string) (bool, int) {
	HasEmail, _ := regexp.MatchString(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`, Account)
	if Account == "" || Other == "" {
		return false, 1
	} else if HasEmail == false {
		return false, 3
	} else {
		return true, 0
	}
}
