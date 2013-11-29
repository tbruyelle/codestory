// +build !heroku

package main

import (
	_ "bitbucket.org/tbruyelle/codestory/elevator"
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
