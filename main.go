package main

// import "github.com/anvh2/be-blogs/cmd"

// func main() {
// 	cmd.Execute()
// }

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Mars!")
	}))
	log.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
