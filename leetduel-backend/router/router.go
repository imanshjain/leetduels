package router

import (
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"

	"leetduel-backend/controller"
	"leetduel-backend/ws"
)

// Required callback when passed auth client
func authWrapper (authClient *auth.Client, fn func(http.ResponseWriter, *http.Request, *auth.Client)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, authClient)
	}
}


func RegisterRoutes(mux *http.ServeMux, authClient *auth.Client, hub *ws.Hub) {
	
	// The router must send ws connection NOT an HTTP one.
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("At the /ws endpoint handler")
		ws.ServeWS(hub, w, r)
	})
	mux.HandleFunc("/auth/getUser", authWrapper(authClient, controller.GetUser))
}