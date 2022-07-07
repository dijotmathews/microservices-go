package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(" root hit!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", d)

	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("good bye")
	})

	http.ListenAndServe(":9000", nil)
}
