package router

import (
	"net/http"
	"leetduel-backend/controller"
	"firebase.google.com/go/v4/auth"
)

// Required callback when passed auth client
func authWrapper (authClient *auth.Client, fn func(http.ResponseWriter, *http.Request, *auth.Client)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, authClient)
	}
}

	func RegisterRoutes(mux *http.ServeMux, authClient *auth.Client) {
		mux.HandleFunc("/question/getQuestion", controller.GetRandomQuestion)
		mux.HandleFunc("/auth/getUser", authWrapper(authClient, controller.GetUser))
	}