package api

import (
	"fmt"
	"net/http"

	"github.com/KidPudel/learn-go-api/db"
)

type WishesHandler struct {
}

func (handler WishesHandler) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Println("error while getting db")
		return
	}

	switch {
	case req.Method == "GET" && len(req.URL.Query()) == 0:
		query := `select name, rate from wishes;`
		var name string
		var rate int
		// write to the pointed values
		db.ConnnectionPool.QueryRow(query).Scan(&name, &rate)
		fmt.Fprintf(resWriter, "name: %v, rate: %v", name, rate)
	default:
		fmt.Fprintln(resWriter, "unhandled request")
	}
}
