package Models

type User struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func (user *User) TableName() string{
	return "user"
}