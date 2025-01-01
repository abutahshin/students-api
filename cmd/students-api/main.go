package main

import (
	"context"
	"github.com/abutahshin/students-api/internal/config"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//Load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students API"))
	})
	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("addr", cfg.Addr))
	//fmt.Printf("Server is started: %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown server gracefully:", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
