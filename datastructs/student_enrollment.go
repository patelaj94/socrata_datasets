package datastructs

type StudentEnrollment struct {
	SchoolYear     string
	DistrictCode   int `json:"DistrictCode,string,omitempty"`
	District       string
	SchoolCode     int `json:"SchoolCode,string,omitempty"`
	Organization   string
	Race           string
	Gender         string
	Grade          string
	SpecialDemo    string
	Geography      string
	SubGroup       string
	RowStatus      string
	Students       int `json:"Students,string,omitempty"`
	EOYEnrollment  int `json:"EOYEnrollment,string,omitempty"`
	FallEnrollment int `json:"FallEnrollment,string,omitempty"`
}

type StudentEnrollmentData []StudentEnrollment

// TODO: params should all be lowercase when making request
type StudentEnrollmentParams struct {
	SchoolYear     string
	DistrictCode   int `json:"DistrictCode,string,omitempty"`
	District       string
	SchoolCode     int `json:"SchoolCode,string,omitempty"`
	Organization   string
	Race           string
	Gender         string
	Grade          string
	SpecialDemo    string
	Geography      string
	SubGroup       string
	RowStatus      string
	Students       int `json:"Students,string,omitempty"`
	EOYEnrollment  int `json:"EOYEnrollment,string,omitempty"`
	FallEnrollment int `json:"FallEnrollment,string,omitempty"`
}
