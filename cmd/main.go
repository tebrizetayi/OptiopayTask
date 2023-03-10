package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tebrizetayi/optiopay/api"
	"github.com/tebrizetayi/optiopay/internal/corporate"
	"github.com/tebrizetayi/optiopay/internal/storage"
)

func main() {

	port := os.Getenv("SERVER_LISTEN_ADDR")
	if port == "" {
		port = ":8080"
	}

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	ctx := context.Background()

	// StorageService
	storageService := storage.NewStorage()
	initEmployees(ctx, storageService)

	// Services
	directoryService := corparate.NewCorporate(ctx, storageService)
	controller := api.NewController(directoryService)

	// Start the HTTP service listening for requests.
	api := http.Server{
		Addr:           port,
		Handler:        api.NewAPI(controller),
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("main : API Listening %s", port)
		serverErrors <- api.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown
	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("main : Error starting server: %+v", err)

	case sig := <-shutdown:
		log.Printf("main : %v : Start shutdown..", sig)
	}
}

func initEmployees(ctx context.Context, s storage.Storage) {
	s.CreateEmployee(ctx, storage.Employee{Id: 1, Name: "Claire", ManagerId: 0})
	s.CreateEmployee(ctx, storage.Employee{Id: 2, Name: "John", ManagerId: 1})
	s.CreateEmployee(ctx, storage.Employee{Id: 3, Name: "Mary", ManagerId: 1})
	s.CreateEmployee(ctx, storage.Employee{Id: 4, Name: "Alice", ManagerId: 2})
	s.CreateEmployee(ctx, storage.Employee{Id: 5, Name: "Bob", ManagerId: 2})
	s.CreateEmployee(ctx, storage.Employee{Id: 6, Name: "Charlie", ManagerId: 3})
}
