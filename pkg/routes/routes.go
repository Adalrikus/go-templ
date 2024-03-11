package routes

import (
  "github.com/adalrikus/go-templ/pkg/controllers"
  "github.com/adalrikus/go-templ/pkg/handlers"
  "github.com/labstack/echo/v4"
  "github.com/adalrikus/go-templ/pkg/views/profile"
)

func InitRoutes(e *echo.Echo) {
  e.GET("/", handlers.IndexHandler)
  e.POST("/login", controllers.LoginUser)
  e.POST("/register", controllers.RegisterNewUser)
  e.GET("/login", handlers.Login)
  e.GET("/logout", handlers.IndexHandler)
  e.GET("/profile", handlers.Profile)
}
