package main

import (
	// "fmt"
	"gpx"
	"log"
	"os"
	"tcx"
)

func lerp(t, a, b float64) float64 {
	return (1 - t) * a + t * b
}

func merge(wrk *tcx.TrainingCenterDatabase, trk *gpx.Gpx) (ans *tcx.TrainingCenterDatabase) {
	ans = new(tcx.TrainingCenterDatabase)
	*ans = *wrk
	// Remove trackpoints from ans
	for nLap, _ := range ans.Activity[0].Lap {
		ans.Activity[0].Lap[nLap].Track[0].Trackpoint = make([]tcx.Trackpoint, 0)
	}
	// Create trackpoints with position and elevation
	nLap := 0
	for _, trkPt := range trk.Trk[0].TrkSeg[0].TrkPt {
		if nLap < (len(ans.Activity[0].Lap) - 1) {
			if ! trkPt.Time.Before(ans.Activity[0].Lap[nLap + 1].StartTime) {
				nLap++
			}
		}
		if ! trkPt.Time.Before(ans.Activity[0].Lap[nLap].StartTime) {
			newTrackpoint := tcx.Trackpoint{}
			newTrackpoint.Time = *trkPt.Time
			newTrackpoint.Position = new(tcx.Position)
			newTrackpoint.Position.Latitude = trkPt.Lat
			newTrackpoint.Position.Longitude = trkPt.Lon
			newTrackpoint.Altitude = new(float64)
			*newTrackpoint.Altitude = *trkPt.Ele
			ans.Activity[0].Lap[nLap].Track[0].Trackpoint =
				append(ans.Activity[0].Lap[nLap].Track[0].Trackpoint, newTrackpoint)
		}
	}
	// Sample original TCX for heart rate and cadence
	for nLap, _ := range ans.Activity[0].Lap {
		for nTrackpoint, trackpoint := range lap.Track[0].Trackpoint {
			// TODO
		}
	}
	return
}

func main() {
	var wrk *tcx.TrainingCenterDatabase
	var trk *gpx.Gpx
	var err error
	if trk, err = gpx.LoadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	}
	if wrk, err = tcx.Load(os.Stdin); err != nil {
		log.Fatal(err)
	}
	if err = merge(wrk, trk).Save(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
