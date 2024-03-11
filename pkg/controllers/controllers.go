package controllers

import (
	"github.com/adalrikus/go-templ/pkg/models"
  "github.com/adalrikus/go-templ/pkg/views/profile"

  "github.com/labstack/echo/v4"

  "net/http"
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
  
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
  return profile.Profile(user).Render(c.Request().Context(), c.Response().Writer)
}

func LoginUser(c echo.Context) error {
  var user = models.User{
    Username: c.FormValue("username"),
    Password: c.FormValue("password"),
  }
  if err := user.Login(); err != nil {
    return err
  }

  c.Response().Writer.WriteHeader(http.StatusOK)
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
  return profile.Profile(user).Render(c.Request().Context(), c.Response().Writer)
}

