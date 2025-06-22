package controller

import (
	"net/http"
	"math/rand"
	"fmt"
)

func GetRandomQuestion(w http.ResponseWriter, r *http.Request) {
	var randomNum int = rand.Intn(3) // TODO: Change this number based on total number of question (or other criteria)
	// TODO: return the json to the client as reponse here
	// return randomNum
	fmt.Println(randomNum)
}