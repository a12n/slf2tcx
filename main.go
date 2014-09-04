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
	var slf *Log
	var gpx *Gpx
	var tcx *Tcx
	switch len(os.Args) {
	case 3:
		tcx = new(Tcx)
	case 4:
		gpx = new(Gpx)
		decodeFile(os.Args[2], gpx)
		tcx = new(Tcx)
	default:
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " input.slf [replace_trk.gpx] output.tcx\n")
		os.Exit(1)
	}
	slf = new(Log)
	decodeFile(os.Args[1], slf)
	fmt.Printf("slf: %+v\n", slf)
	fmt.Printf("gpx: %+v\n", gpx)
	fmt.Printf("tcx: %+v\n", tcx)
}
