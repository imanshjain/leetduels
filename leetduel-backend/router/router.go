package router

import (
	"leetduel-backend/controller"
	"net/http"
)

func RegisterRoutes(a_ctx *controller.AppContext, mux *http.ServeMux) {
	mux.HandleFunc("/question/getQuestion", controller.GetRandomQuestion)
	mux.HandleFunc("/auth/getUser", a_ctx.ChainMiddleware(controller.GetUser, a_ctx.AuthMiddleware))
	mux.HandleFunc("/ws/initMessage", a_ctx.ChainMiddleware(controller.InitMessage, a_ctx.HubMiddleware, a_ctx.DBMiddleware, a_ctx.AuthMiddleware))
}
