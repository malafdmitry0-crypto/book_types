package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, r.Method) })
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("PATCH", "/", nil))
	fmt.Println(rec.Code, rec.Body.String())
}
