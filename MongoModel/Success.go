package MongoModel

import Base "sega/Base"

func SelfSuccess(Number int) Base.StatusData {
	if Number == 1 {
		return Base.StatusData{
			StatusCode:  1,
			Description: "Success Create Account",
		}
	} else if Number == 2 {
		return Base.StatusData{
			StatusCode:  2,
			Description: "Success LogIn",
		}
	} else {
		return Base.StatusData{}
	}
}
