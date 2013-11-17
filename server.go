package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Elevator interface {
	NextCommand() string
	Reset()
	Call(atFloor int)
	Go(floorToGo int)
	UserHasEntered()
	UserHasExited()
}

var elevator Elevator

func main() {
	elevator = NewOmnibus()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/nextCommand", nextCommand)
	http.HandleFunc("/reset", reset)
	http.HandleFunc("/call", call)
	http.HandleFunc("/go", goHandler)
	http.HandleFunc("/userHasEntered", userHasEntered)
	http.HandleFunc("/userHasExited", userHasExited)
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

func call(w http.ResponseWriter, r *http.Request) {
	atFloor, err := strconv.Atoi(r.FormValue("atFloor"))
	if err != nil {
		fmt.Println(err)
		return
	}
	elevator.Call(atFloor)
}

func goHandler(w http.ResponseWriter, r *http.Request) {
	floorToGo, err := strconv.Atoi(r.FormValue("floorToGo"))
	if err != nil {
		fmt.Println(err)
		return
	}
	elevator.Go(floorToGo)
}

func userHasEntered(w http.ResponseWriter, r *http.Request) {
	elevator.UserHasEntered()
}

func userHasExited(w http.ResponseWriter, r *http.Request) {
	elevator.UserHasExited()
}
