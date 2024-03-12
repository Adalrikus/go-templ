package controllers

import (
	"net/http"
	"time"

	"github.com/adalrikus/go-templ/pkg/auth"
	"github.com/adalrikus/go-templ/pkg/models"

	"github.com/labstack/echo/v4"
)

func RegisterNewUser(c echo.Context) error {
	user := models.User{
		Username:  c.FormValue("username"),
		Password:  c.FormValue("password"),
		FirstName: c.FormValue("firstname"),
		LastName:  c.FormValue("lastname"),
		Email:     c.FormValue("email"),
	}
	if err := user.CreateNewUser(); err != nil {
		return err
	}

	if err := auth.GenerateTokensAndSetCookies(c, &user); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusMovedPermanently, "/profile")
}

func LoginUser(c echo.Context) error {
	user := models.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	if err := user.Login(); err != nil {
		return err
	}

	if err := auth.GenerateTokensAndSetCookies(c, &user); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return c.Redirect(http.StatusMovedPermanently, "/profile")
}

func LogoutUser(c echo.Context) error {
	expireCookie(c, auth.GetTokenKey())
	expireCookie(c, auth.GetUserKey())
	return c.Redirect(http.StatusMovedPermanently, "/logout")
}

func expireCookie(c echo.Context, key string) {
	c.SetCookie(&http.Cookie{
		Name:     key,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
}
