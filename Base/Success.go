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
	} else {
		return StatusData{}
	}
}
