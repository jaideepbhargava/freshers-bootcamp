package Models

type Student struct{
	Id int `json:"id"`
	Name string `json:"name"`
	LastName string `json:"last_name"`
	DOB string `json:"dob"`
	Address string `json:"address"`
	Subject string `json:"subject"`
	Marks int `json:"marks"`
}

func (student *Student) StudentTable() string{
	return "student"
}