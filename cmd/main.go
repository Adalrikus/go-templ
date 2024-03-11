package main

import (
  "log"
  "net/http"

  "github.com/adalrikus/go-templ/pkg/models"
  "github.com/adalrikus/go-templ/pkg/routes"

  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  var e = echo.New()
  e.Use(middleware.Logger())
	e.Use(middleware.Recover())

  if err := models.InitDB("database.db"); err != nil {
    log.Fatal(err)
  }
  routes.InitRoutes(e)
  if err := e.Start(":8080"); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}

