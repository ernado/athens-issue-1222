package main

import (
	"log"
	"errors"
)

var checkErr = errors.New("some failure")

func check() error {
	return checkErr
}

func main() {
	if err := check(); err != nil {
		log.Fatal(err)
	}
}
