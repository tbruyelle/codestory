package main

import (
	"fmt"
	"net/http"
)

type Elevator interface {
	NextCommand() string
	Reset()
}

var elevator Elevator

func main() {
	elevator = NewOmnibus()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/nextCommand", nextCommand)
	http.HandleFunc("/reset", reset)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
}

func nextCommand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, elevator.NextCommand())
}

func reset(w http.ResponseWriter, r *http.Request) {
	elevator.Reset()
}
