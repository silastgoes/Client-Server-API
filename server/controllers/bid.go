package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/silastgoes/client-server-api/server/awesomeapi"
	"github.com/silastgoes/client-server-api/server/db"
	repobid "github.com/silastgoes/client-server-api/server/repository/bid"
)

func SetRoutes() {
	http.HandleFunc("/", get)
}

func get(w http.ResponseWriter, r *http.Request) {
	db, err := db.GetConnection()
	if err != nil {
		panic(err)
	}
	bid, err := awesomeapi.GetBidUSDBRL()
	if err != nil {
		panic(err)
	}
	if err = repobid.Save(db, bid); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(bid.Bid)
}
