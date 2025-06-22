package router

import (
	"net/http"
	"leetduel-backend/controller"
)

func RegisterRoutes(mux *http.ServeMux){
	mux.HandleFunc("question/getQuestion", controller.GetRandomQuestion)
}