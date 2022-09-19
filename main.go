package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"undefeated-davout/go-api-sample/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.SetLevel(log.INFO)

	// Routes
	e.GET("/health", health)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctxTimeout); err != nil {
		e.Logger.Fatal(err)
	}
	return nil
}

// Handler
func health(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
