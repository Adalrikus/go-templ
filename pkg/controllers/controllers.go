package controllers

import (
  "net/http"
  "time"

	"github.com/adalrikus/go-templ/pkg/models"
  "github.com/adalrikus/go-templ/pkg/views/profile"

  "github.com/labstack/echo/v4"
  "github.com/golang-jwt/jwt/v5"
)

func RegisterNewUser(c echo.Context) error {
  var user = models.User{
    Username:  c.FormValue("username"),
    Password:  c.FormValue("password"),
    FirstName: c.FormValue("first_name"),
    LastName:  c.FormValue("last_name"),
    Email:     c.FormValue("email"),
  }
  if err := user.CreateNewUser(); err != nil {
    return err
  }
  
  var claims = models.JWTCustomClaims{
    user.Username,
    user.FirstName,
    user.LastName,
    user.Email,
    jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
    },
  }

  var token, err = claims.Create()
  if err != nil {
    return err
  }

  c.JSON(http.StatusOK, echo.Map{
    "token": token,
  })
  
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
  return profile.Profile(claims).Render(c.Request().Context(), c.Response().Writer)
}

func LoginUser(c echo.Context) error {
  var user = models.User{
    Username: c.FormValue("username"),
    Password: c.FormValue("password"),
  }
  if err := user.Login(); err != nil {
    return err
  }
  
  var claims = models.JWTCustomClaims{
    user.Username,
    user.FirstName,
    user.LastName,
    user.Email,
    jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
    },
  }

  var token, err = claims.Create()
  if err != nil {
    return err
  }

  c.JSON(http.StatusOK, echo.Map{
    "token": token,
  })

  c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
  return profile.Profile(claims).Render(c.Request().Context(), c.Response().Writer)
}


