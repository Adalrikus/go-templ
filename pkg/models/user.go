package models

import (
	"errors"

  "github.com/golang-jwt/jwt/v5"
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

type JWTCustomClaims struct {
  Username  string `json:"username"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Email     string `json:"email"`
  jwt.RegisteredClaims
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

func (claims *JWTCustomClaims) Create() (string, error) {
  var token = jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
  return token.SignedString([]byte("secret"))
}
