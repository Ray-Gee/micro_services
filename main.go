package main

import (
	"context"
	"os/signal"
	"github.com/Ryuichi-g/micro_services/handlers"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := http.Server{
		Addr: *bindAddress,
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5 *time.Second,
		WriteTimeout: 10 *time.Second,
		IdleTimeout: 120 *time.Second,
	}

	go func()  {
		l.Println("Starting server on port 9090")
		
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}