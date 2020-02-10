package main

import (
	"flag"
	"fmt"
	"os"

	"log"
	"net/http"
)

func main() {

	port := flag.String("port", "8080", "server port")
	path := flag.String("path", "www", "server path")
	flag.Parse()

	_, err := os.Stat(*path)
	if !os.IsNotExist(err) {
		log.Printf("serving on http://localhost:%s", *port)
		log.Fatalf("%v", http.ListenAndServe(fmt.Sprintf(":%s", *port), http.FileServer(http.Dir(*path))))
	} else {
		log.Printf("path: %s does not exists", *path)
	}

}
