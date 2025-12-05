package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Thripada/task-manager-api-go/handlers"
	"github.com/Thripada/task-manager-api-go/internal/store"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ts := store.NewTaskStore()
	h := handlers.NewTasksHandler(ts)

	r := mux.NewRouter()

	r.HandleFunc("/tasks", h.Create).Methods("POST")
	r.HandleFunc("/tasks", h.List).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.Get).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/tasks/{id}", h.Delete).Methods("DELETE")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		log.Printf("Server listening on :%s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-done
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server exited properly")
}
