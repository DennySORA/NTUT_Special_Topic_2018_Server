package Base

func SelfErrors(Number int) StatusData {
	if Number == 1 {
		return StatusData{
			StatusCode:  -1,
			Description: "Required fields can not be blank.",
		}
	} else if Number == 2 {
		return StatusData{
			StatusCode:  -2,
			Description: "Account exists.",
		}
	} else if Number == 3 {
		return StatusData{
			StatusCode:  -3,
			Description: "Account need Email.",
		}
	} else if Number == 4 {
		return StatusData{
			StatusCode:  -4,
			Description: "Login failed.",
		}
	} else if Number == 5 {
		return StatusData{
			StatusCode:  -5,
			Description: "Not found Account.",
		}
	} else if Number == 6 {
		return StatusData{
			StatusCode:  -6,
			Description: "Token Invalid.",
		}
	} else {
		return StatusData{
			StatusCode:  0,
			Description: "Internal error",
		}
	}
}

func UserERROE() Users {
	return Users{
		Status: StatusData{
			StatusCode:  3,
			Description: "Token nonexists",
		},
	}
}

func LogInERROR() LogInToken {
	return LogInToken{
		Status: StatusData{
			StatusCode:  4,
			Description: "Account error",
		},
	}
}
