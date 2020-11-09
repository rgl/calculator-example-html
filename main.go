package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	log.SetFlags(0)

	var listenAddress = flag.String("listen", "localhost:8000", "Listen address. You can use localhost:0 to listen at a random port.")

	flag.Parse()

	if flag.NArg() != 0 {
		flag.Usage()
		log.Fatalf("\nERROR You MUST NOT pass any positional arguments")
	}

	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	addr := listener.Addr().(*net.TCPAddr)
	fmt.Printf("Listening at http://%s\n", addr)

	err = http.Serve(listener, http.FileServer(assets))
	if err != nil {
		log.Fatalf("Failed to ListenAndServe: %v", err)
	}
}
