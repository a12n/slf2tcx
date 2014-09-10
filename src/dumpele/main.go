package main

import (
	"fmt"
	"log"
	"os"
	"gpx"
)

func dump(trk *gpx.Gpx) {
	for nTrkPt, trkPt := range trk.Trk[0].TrkSeg[0].TrkPt {
		if trkPt.Time != nil && trkPt.Ele != nil {
			fmt.Printf("%d\t%f\n", trkPt.Time.Unix(), *trkPt.Ele)
		} else {
			log.Printf("Trackpoint %d has nil time and/or elevation\n", nTrkPt)
		}
	}
}

func main() {
	var trk *gpx.Gpx
	var err error
	if trk, err = gpx.Load(os.Stdin); err != nil {
		log.Fatal(err)
	}
	dump(trk)
}
