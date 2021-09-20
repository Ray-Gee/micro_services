package main

import (
	"context"
	"os/signal"

	"github.com/Ryuichi-g/micro_services/handlers"

	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr: ":9090",
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