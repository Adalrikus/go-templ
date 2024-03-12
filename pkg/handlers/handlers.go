package handlers

import (
	"github.com/adalrikus/go-templ/pkg/models"
	"github.com/adalrikus/go-templ/pkg/views/profile"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return profile.Index().Render(c.Request().Context(), c.Response().Writer)
}

func LoginHandler(c echo.Context) error {
	return profile.Login().Render(c.Request().Context(), c.Response().Writer)
}

func LogoutHandler(c echo.Context) error {
	return profile.Logout().Render(c.Request().Context(), c.Response().Writer)
}

func RegisterHandler(c echo.Context) error {
	return profile.Register().Render(c.Request().Context(), c.Response().Writer)
}

func ProfileHandler(c echo.Context) error {
	userCookie, err := c.Cookie("user")
	if err != nil {
		return err
	}
	user := models.User{
		Username: userCookie.Value,
	}
	if err := user.Find(); err != nil {
		return err
	}
	return profile.Profile(user).Render(c.Request().Context(), c.Response().Writer)
}
