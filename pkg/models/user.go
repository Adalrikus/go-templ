package models

import (
	"fmt"
	"log"

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

var db *gorm.DB

func InitDB(filename string) (*gorm.DB, error) {
  db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
  if err != nil {
    log.Fatal("failed to connect database")
    return nil, err
  }
  db.AutoMigrate(&User{})
  return db, nil
}

func (u *User) CreateNewUser() error {
  db.Create(u)
  return u.Login()
}

func (u *User) Login() error {
  var result User
  db.Where("Username = ?", u.Username).Where("Password = ?", u.Password).First(&result)
  if result.Username != u.Username {
    return fmt.Errorf("User not found!")
  } else if result.Password != u.Password {
    return fmt.Errorf("Password incorrect!")
  }
  return nil
}
