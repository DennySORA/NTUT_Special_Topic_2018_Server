package Base

func SelfSuccess(Number int) StatusData {
	if Number == 1 {
		return StatusData{
			StatusCode:  1,
			Description: "Success Create Account.",
		}
	} else if Number == 2 {
		return StatusData{
			StatusCode:  2,
			Description: "Success LogIn.",
		}
	} else if Number == 3 {
		return StatusData{
			StatusCode:  3,
			Description: "Success Get Temporarily Token.",
		}
	} else if Number == 4 {
		return StatusData{
			StatusCode:  4,
			Description: "Success Add Car ID to User.",
		}
	} else if Number == 5 {
		return StatusData{
			StatusCode:  5,
			Description: "Success Get User Data.",
		}
	} else if Number == 6 {
		return StatusData{
			StatusCode:  6,
			Description: "Success Return AccountHas.",
		}
	} else if Number == 7 {
		return StatusData{
			StatusCode:  7,
			Description: "Success Get Car ID.",
		}
	} else if Number == 8 {
		return StatusData{
			StatusCode:  8,
			Description: "Success Update User.",
		}
	} else if Number == 9 {
		return StatusData{
			StatusCode:  9,
			Description: "Success Log Out.",
		}
	} else {
		return StatusData{}
	}
}
