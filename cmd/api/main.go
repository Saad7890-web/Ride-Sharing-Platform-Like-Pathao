package main

import (
	"context"
	"log"

	"os"
	"os/signal"
	"syscall"
	"time"

	"ride/internal/auth/service"
	"ride/internal/config"
	"ride/internal/database"

	"ride/internal/router"
	"ride/internal/server"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg.DB)
	if err != nil {
		log.Fatal("db connection failed:", err)
	}
	defer db.Close()

	jwtSvc := service.NewJwtService("super-secret-key", 24*time.Hour)

	handler := router.New(db, jwtSvc)

	srv := server.New(":"+cfg.HTTPPort, handler)


	go srv.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown failed:", err)
	}

	log.Println("Server exited cleanly")
}
