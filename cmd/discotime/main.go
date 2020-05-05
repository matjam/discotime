package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/matjam/discotime/internal/bot"
)

func main() {
	port := os.Getenv("PORT")
	token := os.Getenv("DISCORD_AUTH")

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} uri=${uri} (${status})\n",
	}))
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	go bot.Run(token)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
