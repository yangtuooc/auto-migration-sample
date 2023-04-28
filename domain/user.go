package domain

import "gorm.io/gorm"

type User struct {
	ID   int64  `json:"id" gorm:"primary_key"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

func (u User) TableName() string {
	return "user"
}

func EmptyUser() User {
	return User{}
}

func (u User) AfterCreate(db *gorm.DB) error {
	return db.Select("id").Where("name=?", u.Name).First(&u).Error
}
