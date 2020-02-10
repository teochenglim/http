package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port int

func init()  {
	var port int
	flag.IntVar(&port, "port", 8000, "specify port to use.  defaults to 8000.")
	flag.Parse()
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/ping", ping)
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
