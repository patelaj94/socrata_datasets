package datastructs

type EducatorAverageSalary struct {
	SchoolYear                  string
	DistrictCode                int `json:"DistrictCode,string,omitempty"`
	District                    string
	SchoolCode                  int `json:"SchoolCode,string,omitempty"`
	Organization                string
	Race                        string
	Gender                      string
	Grade                       string
	SpecialDemo                 string
	Geography                   string
	SubGroup                    string
	Staff_Type                  string
	Staff_Category              string
	Experience                  string
	Educators_Fte               float32 `json:"Educators_Fte,string,omitempty"`
	Average_Total_Salary        float32 `json:"Average_Total_Salary,string,omitempty"`
	Average_State_Salary        float32 `json:"Average_State_Salary,string,omitempty"`
	Average_Local_Salary        float32 `json:"Average_Local_Salary,string,omitempty"`
	Average_Federal_Salary      float32 `json:"Average_Federal_Salary,string,omitempty"`
	Average_Years_Of_Experience float32 `json:"Average_Years_Of_Experience,string,omitempty"`
	Average_Years_Of_Age        float32 `json:"Average_Years_Of_Age,string,omitempty"`
}

type EducatorAverageSalaryData []EducatorAverageSalary

// TODO: params should all be lowercase when making request
type EducatorAverageSalaryParams struct {
	SchoolYear                  string
	DistrictCode                int `json:"DistrictCode,string,omitempty"`
	District                    string
	SchoolCode                  int `json:"SchoolCode,string,omitempty"`
	Organization                string
	Race                        string
	Gender                      string
	Grade                       string
	SpecialDemo                 string
	Geography                   string
	SubGroup                    string
	Staff_Type                  string
	Staff_Category              string
	Experience                  string
	Educators_Fte               float32 `json:"Educators_Fte,string,omitempty"`
	Average_Total_Salary        float32 `json:"Average_Total_Salary,string,omitempty"`
	Average_State_Salary        float32 `json:"Average_State_Salary,string,omitempty"`
	Average_Local_Salary        float32 `json:"Average_Local_Salary,string,omitempty"`
	Average_Federal_Salary      float32 `json:"Average_Federal_Salary,string,omitempty"`
	Average_Years_Of_Experience float32 `json:"Average_Years_Of_Experience,string,omitempty"`
	Average_Years_Of_Age        float32 `json:"Average_Years_Of_Age,string,omitempty"`
}
