package main

import (
    "net/http"
    "flag"
    "log"
)

var port = flag.String("port", "3000", "Port to use")
var root = flag.String("path", ".", "Path to serve")

func main() {
    flag.Parse()
    log.Println("Magic happens at port: " + *port)
    panic(http.ListenAndServe(":" + *port, http.FileServer(http.Dir(*root))))
}
