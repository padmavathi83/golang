package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("oops"))
			return
		}
		fmt.Fprintf(rw, "hiiiii_paddu %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Godbye world")

	})
	http.ListenAndServe(":9090", nil)
}
