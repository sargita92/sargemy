package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auth/internal/infra/postgres"
	httpiface "auth/internal/interfaces/http"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := postgres.New(ctx)
	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	httpiface.RegisterHealth(mux)

	server := &http.Server{
		Addr:    ":" + getEnv("HTTP_PORT", "8080"),
		Handler: mux,
	}

	go func() {
		log.Printf("auth service listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = server.Shutdown(shutdownCtx)
	log.Println("auth service stopped gracefully")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
