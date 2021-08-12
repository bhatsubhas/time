package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	router.HandleFunc("/api/time", getCurrentTime)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
