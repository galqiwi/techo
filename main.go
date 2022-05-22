package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var args = getArgs()
var bot = NewBot(args.botToken)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/echo/{message}", echoHandler)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", args.port), r))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message, ok := vars["message"]
	if !ok {
		panic("no message")
	}
	err := bot.SendMessage(args.ownerId, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "internal error\n")
		return
	}
	_, _ = fmt.Fprintf(w, "ok\n")
}
