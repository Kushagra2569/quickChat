package main

import (
	"fmt"
	"log"
	"net/http"

	"quickChat/login"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/login", &login.LoginHandler{})
	log.Println("Listening on :3000...")
	http.ListenAndServe(":3000", mux)
	fmt.Println("Server started on :3000")
}
