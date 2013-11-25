package elevator

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Elevator interface {
	NextCommand() string
	Reset(lowerFloor, higherFloor, cabinSize int, cause string)
	Call(atFloor int, to string)
	Go(floorToGo int)
	UserHasEntered()
	UserHasExited()
	Ditdlamerde()
	Debug(enable bool)
}

var elevator Elevator

func main() {
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	debug := false
	if len(os.Args) >= 2 {
		debug = os.Args[1] == "-d"
		if debug {
			fmt.Println("Debug enabled")
		}
	}
	elevator = NewCabin(0, 5, 50, debug)

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/nextCommand", nextCommandHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/call", callHandler)
	http.HandleFunc("/go", goHandler)
	http.HandleFunc("/userHasEntered", userHasEnteredHandler)
	http.HandleFunc("/userHasExited", userHasExitedHandler)
	http.HandleFunc("/ditdlamerde", shitHandler)
	http.HandleFunc("/debug", debugHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
}

func nextCommandHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, elevator.NextCommand())
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	l, err := strconv.Atoi(r.FormValue("lowerFloor"))
	if err != nil {
		fmt.Println(err)
		l = 0
	}
	h, err := strconv.Atoi(r.FormValue("higherFloor"))
	if err != nil {
		fmt.Println(err)
		h = 5
	}
	c, err := strconv.Atoi(r.FormValue("cabinSize"))
	if err != nil {
		fmt.Println(err)
		c = 50
	}
	elevator.Reset(l, h, c, r.FormValue("cause"))
}

func callHandler(w http.ResponseWriter, r *http.Request) {
	atFloor, err := strconv.Atoi(r.FormValue("atFloor"))
	if err != nil {
		fmt.Println(err)
		return
	}
	elevator.Call(atFloor, r.FormValue("to"))
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

func shitHandler(w http.ResponseWriter, r *http.Request) {
	elevator.Ditdlamerde()
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	elevator.Debug(r.FormValue("enabled") == "true")
}
