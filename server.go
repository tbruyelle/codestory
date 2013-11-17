package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Elevator interface {
	NextCommand() string
	Reset()
	Call(atFloor int, to byte)
	Go(floorToGo int)
	UserHasEntered()
	UserHasExited()
}

var elevator Elevator

func main() {
	elevator = NewOmnibus()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/nextCommand", nextCommandHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/call", callHandler)
	http.HandleFunc("/go", goHandler)
	http.HandleFunc("/userHasEntered", userHasEnteredHandler)
	http.HandleFunc("/userHasExited", userHasExitedHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
}

func nextCommandHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, elevator.NextCommand())
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	elevator.Reset()
}

func callHandler(w http.ResponseWriter, r *http.Request) {
	atFloor, err := strconv.Atoi(r.FormValue("atFloor"))
	if err != nil {
		fmt.Println(err)
		return
	}
	direction := r.FormValue("to")
	var to byte
	if direction == "UP" {
		to = CALLUP
	} else {
		to = CALLDOWN
	}
	elevator.Call(atFloor, to)
}

func goHandler(w http.ResponseWriter, r *http.Request) {
	floorToGo, err := strconv.Atoi(r.FormValue("floorToGo"))
	if err != nil {
		fmt.Println(err)
		return
	}
	elevator.Go(floorToGo)
}

func userHasEnteredHandler(w http.ResponseWriter, r *http.Request) {
	elevator.UserHasEntered()
}

func userHasExitedHandler(w http.ResponseWriter, r *http.Request) {
	elevator.UserHasExited()
}
