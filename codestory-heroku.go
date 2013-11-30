// +build heroku

package main

import (
	_ "codestory/elevator"
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
