package main

import (
	"fmt"
	_ "codestory/elevator"
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
