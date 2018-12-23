package MongoDB

import Base "sega/Base"

type Users struct {
	Email		  string
	Car           []Base.CarData
	Profile       Base.Profiles
	Accesse       Base.Accesses
	SiginHistory  Base.Historys
	LogoutHistory Base.Historys
}
