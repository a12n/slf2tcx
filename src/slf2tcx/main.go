package main

import (
	"encoding/xml"
	"fmt"
	"gpx"
	"log"
	"os"
	"slf"
	"sort"
	"tcx"
)

func Load(path string, ans interface{}) {
	var err error
	var file *os.File
	if file, err = os.Open(path); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err = xml.NewDecoder(file).Decode(ans); err != nil {
		log.Fatal(err)
	}
}

func LoadSlf(path string) *slf.Log {
	ans := new(slf.Log)
	Load(path, ans)
	sort.Sort(slf.LogEntryArray(ans.LogEntries.LogEntry))
	sort.Sort(slf.MarkerArray(ans.Markers.Marker))
	return ans
}

func LoadGpx(path string) *gpx.Gpx {
	ans := new(gpx.Gpx)
	Load(path, ans)
	return ans
}

func main() {
	var slf *slf.Log
	var gpx *gpx.Gpx
	var tcx *tcx.TrainingCenterDatabase = new(tcx.TrainingCenterDatabase)
	switch len(os.Args) {
	case 3:
		slf = LoadSlf(os.Args[1])
	case 4:
		slf = LoadSlf(os.Args[1])
		gpx = LoadGpx(os.Args[2])
	default:
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " input.slf [replace_trk.gpx] output.tcx\n")
		os.Exit(1)
	}
	fmt.Printf("slf: %#v\n", slf)
	fmt.Printf("gpx: %#v\n", gpx)
	fmt.Printf("tcx: %#v\n", tcx)
}
