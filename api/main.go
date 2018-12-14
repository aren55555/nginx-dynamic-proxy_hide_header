package main

import (
	"log"
	"net/http"

	"github.com/aren55555/nginx-dynamic-proxy_hide_header/api/handlers/session"
)

func main() {
	http.Handle("/session", session.Handler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
