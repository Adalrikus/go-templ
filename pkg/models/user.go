package models

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `gorm:"primary_key"       json:"id"`
	Username  string `gorm:"type:varchar(100)" json:"username"`
	Password  string `gorm:"type:varchar(100)" json:"password"`
	FirstName string `gorm:"type:varchar(100)" json:"first_name"`
	LastName  string `gorm:"type:varchar(100)" json:"last_name"`
	Email     string `gorm:"type:varchar(100)" json:"email"`
}

var DB *gorm.DB

func InitDB(filename string) error {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	DB.AutoMigrate(&User{})
	return nil
}

func (u *User) CreateNewUser() error {
	DB.Create(u)
	return u.Login()
}

func (u *User) Login() error {
	var result User
	DB.Where("Username = ?", u.Username).Where("Password = ?", u.Password).First(&result)
	if result.Username != u.Username {
		return errors.New("User not found!")
	} else if result.Password != u.Password {
		return errors.New("Password incorrect!")
	}
	u = &result
	return nil
}

func (u *User) Find() error {
	var result User
	DB.Where("Username = ?", u.Username).First(&result)
	if result.Username != u.Username {
		return errors.New("User not found!")
	}
	*u = result
	return nil
}
