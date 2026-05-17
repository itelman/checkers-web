package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/itelman/checkers-web/internal/game"
	"github.com/itelman/checkers-web/pkg/kithelper"
	"github.com/itelman/checkers-web/pkg/middleware/standard"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := newConfig(ctx)
	if err != nil {
		slog.Error("Unable to get configuration", "error", err)
		return
	}

	deps, err := NewDependencies(
		ctx,
		WithJWTClient(conf.JWTSecret),
		WithLogger(conf.ENV),
	)
	if err != nil {
		slog.Error("Unable to initialize dependencies", "error", err)
		return
	}
	defer deps.Close()

	var (
		standardMiddleware = standard.NewMiddleware(deps.logger)
		corsMiddleware     = cors.New(cors.Options{
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
		}).Handler
		serverOpts = []kithttp.ServerOption{
			kithttp.ServerErrorEncoder(kithelper.NewEncoder(deps.logger).ErrorEncoder),
		}
	)

	router := mux.NewRouter()

	game.RegisterHTTPHandlers(
		router.PathPrefix("/game").Subrouter(),
		game.MakeEndpoints(game.NewService()),
		kithelper.ServerOptions(serverOpts),
	)

	//router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui"))))

	addr := ":" + conf.Port
	server := &http.Server{
		Addr: addr,
		Handler: alice.New(
			corsMiddleware,
			standardMiddleware.RequestLogger,
			standardMiddleware.PanicRecoverer,
		).Then(router),
	}

	errs := make(chan error, 2)
	go func() {
		slog.Info(
			"listening",
			"address", addr,
		)

		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			errs <- err
		}
		slog.Info("Stopped serving new connections")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c

		slog.Info("received shutdown signal", "signal", sig)

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			errs <- err
			return
		}

		errs <- nil
	}()

	if err := <-errs; err != nil {
		slog.Error("terminated with error", "err", err)
	} else {
		slog.Info("terminated gracefully")
	}
}
