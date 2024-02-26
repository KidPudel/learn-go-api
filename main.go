package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KidPudel/learn-go-api/api"
)

func main() {
	// router handler
	mux := http.NewServeMux()

	// simple handler
	mux.HandleFunc("/simple", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "simple handler, that does not contain anything.")
	})

	// for bigger stuff
	mux.Handle("/getWishes", api.WishesHandler{})

	// run a server with wrapped handler, Since servemux
	log.Fatal(http.ListenAndServe(":8081", mux))
}
