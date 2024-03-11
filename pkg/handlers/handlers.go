package handlers

import (
	"github.com/adalrikus/go-templ/pkg/auth"
	"github.com/adalrikus/go-templ/pkg/views/profile"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
  return profile.Index().Render(c.Request().Context(), c.Response().Writer)
}

func LoginHandler(c echo.Context) error {
  return profile.Login().Render(c.Request().Context(), c.Response().Writer)
}

func RegisterHandler(c echo.Context) error {
  return profile.Register().Render(c.Request().Context(), c.Response().Writer)
}

func ProfileHandler(c echo.Context) error {
  var user = c.Get("user").(*jwt.Token)
  var claims = user.Claims.(*auth.JWTCustomClaims)
  return profile.Profile(*claims).Render(c.Request().Context(), c.Response().Writer)
}
