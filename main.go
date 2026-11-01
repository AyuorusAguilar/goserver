package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AyuorusAguilar/goserver/internal/handlers"
)

func main()  {
	/* Getting the server up */
	// Good old mutex
	mutex := http.NewServeMux()
	mutex.HandleFunc("/", handlers.HandleIndex)

	port := 8010
	srv_conf := http.Server{
		Handler: mutex,
		Addr: fmt.Sprintf(":%d", port),
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	fmt.Printf("Server listening port %d", port)
	err:= srv_conf.ListenAndServe()
	log.Fatal(err)
}