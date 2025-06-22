package router

import (
	"net/http"
	"leetduel-backend/controller/questions"
	"leetduel-backend/controller/auth"
)

func RegisterRoutes(mux *http.ServeMux){
	mux.handleFunc("question/getQuestion", questions.GetRandomQuestion)
}