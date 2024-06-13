package login

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type loginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginHandler struct {
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you're on the login page.")
	switch r.Method {
	case http.MethodPost:
		handlePost(w, r)
	default:
		fmt.Fprintln(w, "Sorry, only POST methods are supported.")
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "Error parsing form")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "Error reading body")
		return
	}
	defer r.Body.Close()

	var login loginDetails
	if err := json.Unmarshal(body, &login); err != nil {
		fmt.Fprintln(w, "Invalid JSON")
	}
	fmt.Fprintf(w, "Username: %s\n", login.Username)
	fmt.Fprintf(w, "Password: %s\n", login.Password)
}
