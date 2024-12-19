package main

import (
	"context"
	"flag"
	"goginmvc/database"
	"goginmvc/router"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var env string

func main() {
	flag.StringVar(&env, "env", "dev", "start with dev or prod")
	flag.Parse()

	// Create context that listens for the interrupt sigal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// init db
	db, err := database.InitDB(env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init gin
	r := gin.Default()

	// init routers
	router.InitRouter(r)

	// start server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()
	log.Println("Server started on :8080")

	// Listen for the interrupt signal
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C to stop")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}
	log.Println("Server exited gracefully")
}
