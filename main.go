package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"undefeated-davout/echo-api-sample/config"
	"undefeated-davout/echo-api-sample/frameworks_drivers/database"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways"

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

func run(ctx context.Context) (err error) {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.SetLevel(log.INFO)

	// DB用意
	db, err := database.InitDB(cfg)
	if err != nil {
		return err
	}

	// ルーティング
	if err := gateways.NewRouter(ctx, e, db, cfg); err != nil {
		return err
	}

	// サーバ起動
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// グレースフルシャットダウン設定
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
