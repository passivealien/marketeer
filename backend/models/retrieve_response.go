package models

type RetrieveResponse struct {
	Code       int32 `json:"code"`
	Status     string
	Message    string `json:"message,omitempty"`
	ID         int    `json:"ID"`
	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	EMail      string `json:"E-Mail"`
	ContactNum int64  `json:"ContactNumber,string"`
	BirthDate  string `json:"BirthDate"`
	Address    string `json:"Address"`
	UserName   string `json:"Username"`
	Password   string `json:"Password"`
	//Data    interface{} `json:",omitempty"`

}
