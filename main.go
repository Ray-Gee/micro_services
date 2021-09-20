package main

import (
	"github.com/Ryuichi-g/micro_services/handlers"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)	

	http.ListenAndServe(":9090", sm)
}