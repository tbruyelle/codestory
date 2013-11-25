package main

import (
	"fmt"
	_ "misc/codestory/elevator"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
