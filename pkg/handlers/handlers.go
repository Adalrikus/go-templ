package handlers

import (
  "github.com/adalrikus/go-templ/pkg/models"
  "github.com/labstack/echo/v4"
  "github.com/adalrikus/go-templ/pkg/views/profile"
)

func IndexHandler(c echo.Context) error {
  return profile.Index().Render(c.Request().Context(), c.Response().Writer)
}

func LoginHandler(c echo.Context) error {
  return profile.Login().Render(c.Request().Context(), c.Response().Writer)
}

func ProfileHandler(c echo.Context) error {
  var user = c.Get("user").(*models.User)
  return profile.Profile(user).Render(c.Request().Context(), c.Response().Writer)
}
