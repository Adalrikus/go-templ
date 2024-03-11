package routes

import (
  "github.com/adalrikus/go-templ/pkg/auth"
  "github.com/adalrikus/go-templ/pkg/controllers"
  "github.com/adalrikus/go-templ/pkg/handlers"

  "github.com/labstack/echo/v4"
  echojwt "github.com/labstack/echo-jwt/v4"
  "github.com/golang-jwt/jwt/v5"
)

func InitRoutes(e *echo.Echo) {
  e.GET("/", handlers.IndexHandler)
  e.POST("/login", controllers.LoginUser)
  e.POST("/register", controllers.RegisterNewUser)
  e.GET("/login", handlers.LoginHandler)
  e.GET("/logout", handlers.IndexHandler)
  e.GET("/profile", handlers.ProfileHandler)
  e.GET("/register", handlers.RegisterHandler)
  var profieGroup = e.Group("/profile")
  var config = echojwt.Config{
    NewClaimsFunc: func(c echo.Context) jwt.Claims {
      return new(auth.JWTCustomClaims)
    },
    SigningKey: []byte("secret"),
  }
  profieGroup.Use(echojwt.WithConfig(config))
  profieGroup.GET("", handlers.ProfileHandler)
}
