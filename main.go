package main

import (
	"net/http"
	"os"
)

type Student struct {
	Name  string
	ID    int32
	Class string
}

func init() {

}

func main() {

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
