package main

import (
	"net/http"
	"fmt"
	"os"
	//"github.com/VividCortex/godaemon"
)

var parsedString string

func exit(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func boi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "BOIIIII\n")
}

func main() {
	//address := "192.168.1.40:8081"
	//godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	http.HandleFunc("/BOI", boi)
	http.HandleFunc("/serverFunctions/exit", exit)
	http.ListenAndServe(":7000", nil)
	//fmt.Print("%s", address)
}

