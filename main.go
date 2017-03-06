package main

import (
	"os"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		h, _ := os.Hostname()
		w.Write([]byte(h))
	})
	
	http.ListenAndServe(":8000", nil)
}
