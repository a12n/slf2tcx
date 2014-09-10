package main

import (
	"fmt"
	"log"
	"os"
	"tcx"
)

func dump(wrk *tcx.TrainingCenterDatabase) {
	for nTrkpt, trkpt := range wrk.Activity[0].Lap[0].Track[0].Trackpoint {
		if trkpt.Altitude != nil {
			fmt.Printf("%d\t%f\n", trkpt.Time.Unix(), *trkpt.Altitude)
		} else {
			log.Printf("Trackpoint %d has nil altitude\n", nTrkpt)
		}
	}
}

func main() {
	var wrk *tcx.TrainingCenterDatabase
	var err error
	if wrk, err = tcx.Load(os.Stdin); err != nil {
		log.Fatal(err)
	}
	dump(wrk)
}
