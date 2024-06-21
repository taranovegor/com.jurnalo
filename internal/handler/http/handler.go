package http

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"log"
	"net/http"
)

type Handler interface {
}

func parseQuery(req *http.Request, i interface{}) {
	decoder := schema.NewDecoder()
	decoder.Decode(i, req.URL.Query())
}

func resp(rw http.ResponseWriter, status int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		log.Printf("Error during response: %s\n", err)
	}
}

func errno(rw http.ResponseWriter, status int, message string) {
	resp(rw, status, struct {
		Code    int
		Message string
	}{status, message})
}
