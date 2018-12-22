package MongoModel

import Base "sega/Base"

func SelfErrors(Number int) Base.StatusData {
	if Number == 1 {
		return Base.StatusData{
			StatusCode:  -1,
			Description: "Required fields can not be blank.",
		}
	} else if Number == 2 {
		return Base.StatusData{
			StatusCode:  -2,
			Description: "Account exists.",
		}
	} else if Number == 3 {
		return Base.StatusData{
			StatusCode:  -3,
			Description: "Account need Email.",
		}
	} else if Number == 4 {
		return Base.StatusData{
			StatusCode:  -4,
			Description: "Login failed.",
		}
	} else if Number == 5 {
		return Base.StatusData{
			StatusCode:  -5,
			Description: "Not found Account.",
		}
	} else {
		return Base.StatusData{
			StatusCode:  -3,
			Description: "Internal error",
		}
	}
}

func UserERROE() Base.Users {
	return Base.Users{
		Status: Base.StatusData{
			StatusCode:  3,
			Description: "Token nonexists",
		},
	}
}

func LogInERROR() Base.LogInToken {
	return Base.LogInToken{
		Status: Base.StatusData{
			StatusCode:  4,
			Description: "Account error",
		},
	}
}
