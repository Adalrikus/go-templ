package auth

import (
  "github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
  Username  string `json:"username"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Email     string `json:"email"`
  jwt.RegisteredClaims
}

func (claims *JWTCustomClaims) Create() (string, error) {
  var token = jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
  return token.SignedString([]byte("secret"))
}
