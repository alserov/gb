package app

import (
	"fmt"
	"github.com/alserov/fuze"
	"github.com/alserov/gb/sm_2_4/internal/config"
	"github.com/alserov/gb/sm_2_4/internal/db/postgres"
	"github.com/alserov/gb/sm_2_4/internal/log"
	"github.com/alserov/gb/sm_2_4/internal/routes"
	"github.com/alserov/gb/sm_2_4/internal/server"
	"github.com/alserov/gb/sm_2_4/internal/service"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	l := log.MustSetup(cfg.Env)

	defer func() {
		err := recover()
		if err != nil {
			l.Err("panic recovery", fmt.Errorf(err.(string)))
		}
	}()

	l.Debug("starting app", slog.Any("cfg", cfg))

	var (
		addr = fmt.Sprintf(":%d", cfg.Port)
	)

	// router init
	a := fuze.NewApp(fuze.WithAddr(addr), fuze.WithTimeouts(cfg.Timeout.Read, cfg.Timeout.Write))

	// repo
	repo := postgres.NewRepository(postgres.MustConnect(cfg.DB.Addr))

	// bll
	srvc := service.NewService(l, repo)

	// controller
	srvr := server.NewServer(l, srvc)

	routes.Setup(a.Controller, srvr)

	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		l.Info("app is running")
		if err := a.Run(); err != nil {
			l.Err("failed to run app", err)
		}
	}()

	<-chStop
	if err := a.GracefulShutdown(); err != nil {
		l.Err("failed to shutdown gracefully", err)
	}
	l.Info("app was stopped")
}
