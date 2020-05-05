package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/matjam/discotime/internal/bot"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func formatQuotedMessage(i interface{}) string {
	if s, ok := i.(string); ok {
		return fmt.Sprintf("\"%s\"", strings.ReplaceAll(s, "\"", "\\\""))
	}

	return fmt.Sprintf("%s", i)
}

func main() {
	port := os.Getenv("PORT")
	token := os.Getenv("DISCORD_AUTH")

	out := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    true,
		TimeFormat: ""}
	out.FormatMessage = func(i interface{}) string { return fmt.Sprintf("message=%v", formatQuotedMessage(i)) }
	out.FormatTimestamp = func(i interface{}) string { return "" }
	out.FormatLevel = func(i interface{}) string { return fmt.Sprintf("level=%s", i) }
	out.FormatFieldName = func(i interface{}) string { return fmt.Sprintf("%s=", i) }
	out.FormatFieldValue = formatQuotedMessage
	log.Logger = log.Output(out)

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	go bot.Run(token)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
