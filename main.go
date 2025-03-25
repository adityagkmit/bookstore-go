package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/adityagkmit/bookstore/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to Database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize the app
	app := NewApp(db)

	// Handle graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}
