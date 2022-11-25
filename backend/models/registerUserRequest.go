package models

type RegistrationInput struct {
	//UserID    int    `json:"User ID"`

	FirstName  string `json:"FirstName"`
	LastName   string `json:"LastName"`
	EMail      string `json:"E-Mail"`
	ContactNum int64  `json:"ContactNumber,string"`
	BirthDate  string `json:"BirthDate"`
	Address    string `json:"Address"`
	UserName   string `json:"Username"`
	Password   string `json:"Password"`
}
