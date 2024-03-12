package routes

import (
	"github.com/adalrikus/go-templ/pkg/auth"
	"github.com/adalrikus/go-templ/pkg/controllers"
	"github.com/adalrikus/go-templ/pkg/handlers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.IndexHandler)
	e.POST("/login", controllers.LoginUser)
	e.POST("/register", controllers.RegisterNewUser)
	e.POST("/logout", controllers.LogoutUser)
	e.GET("/login", handlers.LoginHandler)
	e.GET("/logout", handlers.LogoutHandler)
	e.GET("/profile", handlers.ProfileHandler)
	e.GET("/register", handlers.RegisterHandler)
	profieGroup := e.Group("/profile")
	config := echojwt.Config{
		SigningKey:  []byte(auth.GetJWTSecret()),
		TokenLookup: "cookie:access-token",
	}

	profieGroup.Use(echojwt.WithConfig(config))
	profieGroup.GET("", handlers.ProfileHandler)
}
