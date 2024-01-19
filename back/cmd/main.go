package main

import (
	"back/internal/api"
	"back/internal/config"
	"back/internal/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := start()
	if err != nil {
		fmt.Printf("Error starting app: %s\n", err.Error())
	}
}

func start() error {
	// Config
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Database
	err = database.InitDB()
	if err != nil {
		return err
	}

	// API
	router := api.InitRoutes()

	// Server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	server := api.InitServer(router)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-stop:
		cancel()
	default:
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println("Graceful shutdown failed. Forcing shutdown.")
		}
	}

	fmt.Println("Server gracefully shut down")

	return nil
}
