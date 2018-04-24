package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

//create router and return the value of router
//ex: /hello
func newRouter() *mux.Router  {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird", GetBirdHandler).Methods("GET")
	r.HandleFunc("/bird", CreateBirdHandler).Methods("POST")
	staticFileDirectory := http.Dir("./assets/")

	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	// Declare a new router by calling newRouter method
	r := newRouter()
	http.ListenAndServe(":2000", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

