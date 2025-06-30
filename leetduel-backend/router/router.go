package router

import (
	"leetduel-backend/controller"
	"net/http"
)

func RegisterRoutes(a_ctx *controller.AppContext, mux *http.ServeMux) {
	mux.HandleFunc("/question/getQuestion", controller.GetRandomQuestion)
	mux.HandleFunc("/auth/getUser", a_ctx.ChainMiddleware(controller.GetUser, a_ctx.AuthMiddleware))
}
