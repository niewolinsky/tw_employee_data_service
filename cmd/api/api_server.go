package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serveREST(port string) error {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	// Channel to signal when the shutdown is complete
	shutdownDone := make(chan struct{})

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		slog.Info(fmt.Sprintf("shutdown signal received: %s", s.String()))

		// Shutdown the server with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("error during server shutdown:", err)
		} else {
			slog.Info("server gracefully shutdown")
		}

		close(shutdownDone)
	}()

	slog.Info(fmt.Sprintf("Starting REST server on port %s", port))
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// Wait for shutdown to complete
	<-shutdownDone
	slog.Info("shutdown complete")
	return nil
}
