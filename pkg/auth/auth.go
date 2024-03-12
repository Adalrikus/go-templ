package auth

import (
  "time"
  "net/http"

  "github.com/adalrikus/go-templ/pkg/models"
  "github.com/golang-jwt/jwt/v5"

  "github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
  Username  string `json:"username"`
  jwt.RegisteredClaims
}

const (
  TOKEN_KEY = "access-token"
  USER_KEY  = "user"
  SECRET    = "secret"
)

func GenerateTokensAndSetCookies(c echo.Context, user *models.User) error {
  var claims = &JWTCustomClaims{
    user.Username,
    jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
    },
  }

  var token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET))
  if err != nil {
    return err
  }

  setCookies(c, TOKEN_KEY, token)
  setCookies(c, USER_KEY, user.Username)

  return nil
}

func setCookies(c echo.Context, name string, value string) {
  c.SetCookie(&http.Cookie{
    Name:     name,
    Value:    value,
    Path:     "/",
    HttpOnly: true,
    Expires:  time.Now().Add(time.Hour * 1),
  })
}

func GetJWTSecret() string {
  return SECRET
}

func GetTokenKey() string {
  return TOKEN_KEY
}

func GetUserKey() string {
  return USER_KEY
}
