package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KidPudel/learn-go-api/db"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Wish struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Rate       int    `json:"rate,omitempty"`
	WishListId int    `json:"wish_list_id,omitempty"`
}

type WishesHandler struct {
}

func (handler WishesHandler) ServeHTTP(resWriter http.ResponseWriter, req *http.Request) {
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Println("error while getting db", err)
		return
	}

	connnectionPool, err := db.Pool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintln(resWriter, "error estblishing connection", err)
		return
	}
	connection := connnectionPool.Conn()
	defer connection.Close(context.Background())

	switch {
	case req.Method == "GET" && len(req.URL.Query()) == 0:
		query := `select * from wishes w where w.id = $1;`
		var wish Wish

		template, err := connection.Prepare(context.Background(), "get_wish", query)
		if err != nil {
			return
		}
		// write to the pointed values
		rows, err := connection.Query(context.Background(), template.SQL, 2)
		if err != nil {
			fmt.Fprintln(resWriter, err.Error())
			return
		}
		// first check, go on 0 index
		for rows.Next() {
			err := pgxscan.ScanRow(&wish, rows)
			if err != nil {
				fmt.Println(err)
				break
			}

		}

		// or

		// []*Wish
		var wishes []map[string]string
		pgxscan.ScanAll(&wishes, rows)
		fmt.Fprintf(resWriter, "wish: %v", wishes)

		fmt.Fprintf(resWriter, "wish: %v", wish)
	default:
		fmt.Fprintln(resWriter, "unhandled request")
	}
}
