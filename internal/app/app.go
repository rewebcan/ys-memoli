package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/rewebcan/ys-memoli/internal/controller/http/v1"
	"github.com/rewebcan/ys-memoli/internal/usecase/repo"
	"github.com/rewebcan/ys-memoli/pkg/memoli"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(cfg Config) {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.RequestID())

	mdb, err := memoli.NewBucket(
		"ys",
		memoli.SnapshotPath(cfg.MemoliPath),
		memoli.SnapshotWindow(time.Second), // for test purposes
	)
	if err != nil {
		panic(err)
	}
	settingsRepository := repo.NewMemoliDBSettingsRepository(mdb)

	v1.NewRouter(e, settingsRepository)

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	mdb.Close()
}
