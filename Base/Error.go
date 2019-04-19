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
	} else if Number == 7 {
		return StatusData{
			StatusCode:  -7,
			Description: "Add failure.",
		}
	} else if Number == 8 {
		return StatusData{
			StatusCode:  -8,
			Description: "Token remove error.",
		}
	} else if Number == 9 {
		return StatusData{
			StatusCode:  -9,
			Description: "Token Limit.",
		}
	} else if Number == 10 {
		return StatusData{
			StatusCode:  -10,
			Description: "Strong Password Not Good.",
		}
	} else if Number == 11 {
		return StatusData{
			StatusCode:  -11,
			Description: "Password is error.",
		}
	} else if Number == 12 {
		return StatusData{
			StatusCode:  -12,
			Description: "Password is same.",
		}
	} else {
		return StatusData{
			StatusCode:  0,
			Description: "Internal error",
		}
	}
}
