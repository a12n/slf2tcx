package main

import (
	"fmt"
	"os"
)

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
