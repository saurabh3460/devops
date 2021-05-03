package main

import (
	"flag"
	"fmt"
)

var port int

type message string

// var serverStartMessage message

func init() {
	flag.IntVar(&port, "port", 8080, "use to custom port")
	// flag.Var(serverStartMessage, "m", "write message to show up when server starts")
}

func main() {
	flag.Parse()
	fmt.Printf("port= %d\n", port)
	fmt.Printf("other args: %+v\n", flag.Args())
}
