package server

import (
	"fmt"
fmt.Fprintf(w, "Go Server Working!")	"log"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("default: ", r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("param: ", r.Form["test_param"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Go Server Working!")
}

func Start() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("ListenAndServe Started! -> Port(8080)")
	}
}
