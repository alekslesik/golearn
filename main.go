package main

import (
	"fmt"
	"net/http"
)

func main()  {
	
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	http.ListenAndServe(":80", mux)
}