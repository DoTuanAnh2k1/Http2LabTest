package main

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	messHeader := r.URL.Query().Get("name")
	fmt.Println("Get mess ", messHeader)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", TestHandler)

	return mux
}

func newServer() *http.Server {
	mux := newRouter()
	return &http.Server{
		Addr:    ":3654",
		Handler: mux,
	}
}

func main() {

}
