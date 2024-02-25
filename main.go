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

	mux.HandleFunc("/simple", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "simple handler, that does not contain anything.")
	})

	mux.Handle("/getWishes", api.WishesHandler{})

	log.Fatal(http.ListenAndServe(":8081", mux))
}

// instead of declaring db (symbolising all big and important data) at the root, since we doesn't need it every where
// it's _obvious_ to separate it into its own **case**, meaning, we need specific handler as well, where further filtering will be proceed
// what is the difference between handing at the base, it avoids declaring (sticking) db at the root or even more stupid, in each handler

// don't stress and overthink
// don't make simple things hard
// router is to navigate the request!!!

// Logically place project
