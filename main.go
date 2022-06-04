package main

import "net/http"

func main() {
	http.Handle("/", new(myHandler))
	http.ListenAndServe(":8080", nil)
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
