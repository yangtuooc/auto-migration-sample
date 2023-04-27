package domain

type User struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func EmptyUser() User {
	return User{}
}
