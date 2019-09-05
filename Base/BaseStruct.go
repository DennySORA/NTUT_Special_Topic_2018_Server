// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package Base

var ServerStop chan int

type Accesses struct {
	Certification bool          `json:"Certification"`
	PermitTime    string        `json:"PermitTime"`
	Level         int           `json:"Level"`
	PermitLog     []*PermitLogs `json:"Permit_log"`
}

type AccountHas struct {
	Status StatusData `json:"Status"`
	Has    bool       `json:"Has"`
}

type CarData struct {
	Status      StatusData `json:"Status"`
	CarID       string     `json:"CarID"`
	CarName     string     `json:"CarName"`
	RefreshTime *string    `json:"RefreshTime"`
	CreateTime  *string    `json:"CreateTime"`
}

type CarIDReturn struct {
	Status    StatusData `json:"Status"`
	AccountID string     `json:"AccountID"`
	CarToken  string     `json:"CarToken"`
}

type CreateReturn struct {
	Status    StatusData `json:"Status"`
	AccountID string     `json:"AccountID"`
}

type Historys struct {
	Times    string `json:"Times"`
	UseToken string `json:"UseToken"`
	Types    int    `json:"Types"`
	Device   string `json:"Device"`
}

type LogInToken struct {
	Status       StatusData `json:"Status"`
	GetTimes     string     `json:"GetTimes"`
	AccountToken string     `json:"AccountToken"`
	AccountID    string     `json:"AccountID"`
}

type Logformation struct {
	Type   string `json:"Type"`
	Device string `json:"Device"`
}

type MonitorData struct {
	Status         StatusData        `json:"Status"`
	WaterStatus    MonitorStatusData `json:"WaterStatus"`
	GasolineStatus MonitorStatusData `json:"GasolineStatus"`
	BatteryStatus  MonitorStatusData `json:"BatteryStatus"`
}

type MonitorStatus struct {
	CarToken     string `json:"CarToken"`
	SelectObject string `json:"SelectObject"`
	StatusCode   int    `json:"StatusCode"`
}

type MonitorStatusData struct {
	StatusCode  int     `json:"StatusCode"`
	RefreshTime *string `json:"RefreshTime"`
}

type NewAccountIDPw struct {
	AccountID string `json:"AccountID"`
	Password  string `json:"Password"`
}

type NewAccountUser struct {
	Name          string `json:"Name"`
	Gender        int    `json:"Gender"`
	CountryNumber string `json:"CountryNumber"`
	PhoneNumber   string `json:"PhoneNumber"`
}

type NewCarName struct {
	AccountID string `json:"AccountID"`
	CarID     string `json:"CarID"`
	CarName   string `json:"CarName"`
}

type PermitLogs struct {
	Level     int    `json:"Level"`
	Times     string `json:"Times"`
	Authority string `json:"Authority"`
}

type Phones struct {
	CountryNumber string `json:"CountryNumber"`
	PhoneNumber   string `json:"PhoneNumber"`
}

type Profiles struct {
	Name   string `json:"Name"`
	Gender int    `json:"Gender"`
	Phone  Phones `json:"Phone"`
}

type SecurityData struct {
	Status       StatusData            `json:"Status"`
	DoorStatus   []*SecurityStatusData `json:"DoorStatus"`
	WindowStatus []*SecurityStatusData `json:"WindowStatus"`
	LightStatus  []*SecurityStatusData `json:"LightStatus"`
}

type SecurityStatus struct {
	CarToken     string `json:"CarToken"`
	Name         string `json:"Name"`
	SelectObject string `json:"SelectObject"`
	StatusCode   int    `json:"StatusCode"`
}

type SecurityStatusData struct {
	Name        string  `json:"Name"`
	StatusCode  int     `json:"StatusCode"`
	RefreshTime *string `json:"RefreshTime"`
}

type StatusData struct {
	StatusCode  int    `json:"StatusCode"`
	Description string `json:"Description"`
}

type TemporarilyTokenData struct {
	Status   StatusData `json:"Status"`
	Token    string     `json:"Token"`
	GetTimes string     `json:"GetTimes"`
}

type Users struct {
	Status        StatusData `json:"Status"`
	Car           []CarData  `json:"Car"`
	Profile       Profiles   `json:"Profile"`
	Accesse       Accesses   `json:"Accesse"`
	SiginHistory  []Historys `json:"SiginHistory"`
	LogoutHistory []Historys `json:"LogoutHistory"`
}
