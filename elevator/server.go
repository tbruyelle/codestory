package elevator

import (
	"fmt"
	"html/template"
	"log"
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

type Elevators interface {
	NextCommands() []string
	Reset(lowerFloor, higherFloor, cabineSize, cabinCaount int, cause string)
	Call(atFloor int, to string)
	Go(floorToGo, cabin int)
	UserHasEntered(cabin int)
	UserHasExited(cabin int)
	Ditdlamerde()
	Debug(enable bool)
}

var elevators Elevators

func init() {
	debug := false
	if len(os.Args) >= 2 {
		debug = os.Args[1] == "-d"
		if debug {
			log.Println("Debug enabled")
		}
	}
	elevators = NewCabins(0, 5, 50, 2, debug)

	http.HandleFunc("/", defaultHandler)
	//http.HandleFunc("/nextCommand", nextCommandHandler)
	http.HandleFunc("/nextCommands", nextCommandsHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/call", callHandler)
	http.HandleFunc("/go", goHandler)
	http.HandleFunc("/userHasEntered", userHasEnteredHandler)
	http.HandleFunc("/userHasExited", userHasExitedHandler)
	http.HandleFunc("/ditdlamerde", shitHandler)
	http.HandleFunc("/debug", debugHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t,err:=template.ParseFiles("views/index.html")
	if err!=nil{
	fmt.Fprint(w, err)
	return
	}
	t.Execute(w,elevators) 
}

func nextCommandsHandler(w http.ResponseWriter, r *http.Request) {
	cs := elevators.NextCommands()
	for _, c := range cs {
		fmt.Fprintln(w, c)
	}
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	l, err := strconv.Atoi(r.FormValue("lowerFloor"))
	if err != nil {
		log.Println(err)
		l = 0
	}
	h, err := strconv.Atoi(r.FormValue("higherFloor"))
	if err != nil {
		log.Println(err)
		h = 5
	}
	c, err := strconv.Atoi(r.FormValue("cabinSize"))
	if err != nil {
		log.Println(err)
		c = 50
	}
	cc, err := strconv.Atoi(r.FormValue("cabinCount"))
	if err != nil {
		log.Println(err)
		cc = 2
	}
	elevators.Reset(l, h, c, cc, r.FormValue("cause"))
}

func callHandler(w http.ResponseWriter, r *http.Request) {
	atFloor, err := strconv.Atoi(r.FormValue("atFloor"))
	if err != nil {
		log.Println(err)
		return
	}
	elevators.Call(atFloor, r.FormValue("to"))
}

func goHandler(w http.ResponseWriter, r *http.Request) {
	floorToGo, err := strconv.Atoi(r.FormValue("floorToGo"))
	if err != nil {
		log.Println(err)
		return
	}
	cabin, err := strconv.Atoi(r.FormValue("cabin"))
	if err != nil {
		log.Println(err)
		return
	}

	elevators.Go(floorToGo, cabin)
}

func userHasEnteredHandler(w http.ResponseWriter, r *http.Request) {
	cabin, err := strconv.Atoi(r.FormValue("cabin"))
	if err != nil {
		log.Println(err)
		return
	}

	elevators.UserHasEntered(cabin)
}

func userHasExitedHandler(w http.ResponseWriter, r *http.Request) {
	cabin, err := strconv.Atoi(r.FormValue("cabin"))
	if err != nil {
		log.Println(err)
		return
	}
	elevators.UserHasExited(cabin)
}

func shitHandler(w http.ResponseWriter, r *http.Request) {
	elevators.Ditdlamerde()
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	elevators.Debug(r.FormValue("enabled") == "true")
}
