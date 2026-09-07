package main

import (
	"net/http"
)

func main() {
	var h http.Handler = func(http.ResponseWriter, *http.Request) {}
	_ = h
}
