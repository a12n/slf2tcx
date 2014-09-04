package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func decodeFile(name string, v interface{}) {
	var err error
	var file *os.File
	if file, err = os.Open(name); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err = xml.NewDecoder(file).Decode(v); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// var slf *Log
	// var gpx *Gpx
	// var tcx *Tcx
	switch len(os.Args) {
	case 3:
		// slf = new(Log)
		// tcx = new(Tcx)
	case 4:
		// slf = new(Log)
		// gpx = new(Gpx)
		// tcx = new(Tcx)
	default:
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " input.slf [replace_trk.gpx] output.tcx\n")
		os.Exit(1)
	}
}
