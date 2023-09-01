package models

type Security struct {
	UserId      int64  `json:"userId"`
	FirsrtName  string `json:"firstName"`
	LastName    string `json:"lastName"`
	From        string `json:"from"`
	To          string `json:"to"`
	FirsrtNameD string `json:"firstNameD"`
	LastNameD   string `json:"lastNameD"`
	CarNumber   string `json:"carNumber"`
	TimeStart   string `json:"timeStart"`
}
