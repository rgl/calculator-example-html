package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(0)

	var listenAddress = flag.String("listen", "localhost:8000", "Listen address.")

	flag.Parse()

	if flag.NArg() != 0 {
		flag.Usage()
		log.Fatalf("\nERROR You MUST NOT pass any positional arguments")
	}

	fmt.Printf("Listening at http://%s\n", *listenAddress)

	http.Handle("/", http.FileServer(assets))

	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatalf("Failed to ListenAndServe: %v", err)
	}
}
