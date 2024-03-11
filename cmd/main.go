package main

import (
  "log"
  "net/http"

  "github.com/adalrikus/go-templ/pkg/models"
  "github.com/adalrikus/go-templ/pkg/routes"

  "github.com/labstack/echo/v4"
  "github.com/labstack/echo-jwt/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  var e = echo.New()
  e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
  
  models.InitDB("database.db")
  routes.InitRoutes(e)
  if err := e.Start(":8080"); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}

