package main

import (
	// "fmt"
	"gpx"
	"log"
	"os"
	"math"
	"time"
	"tcx"
)

func lerp(t, a, b float64) float64 {
	return (1 - t) * a + t * b
}

func merge(wrk *tcx.TrainingCenterDatabase, trk *gpx.Gpx) {
	// Remove trackpoints from ans
	savedTrack := make([][]tcx.Trackpoint, len(wrk.Activity[0].Lap))
	for nLap, _ := range wrk.Activity[0].Lap {
		savedTrack[nLap] = wrk.Activity[0].Lap[nLap].Track[0].Trackpoint
		wrk.Activity[0].Lap[nLap].Track[0].Trackpoint = nil
	}
	// Create trackpoints with position and elevation
	nLap := 0
	for _, trkPt := range trk.Trk[0].TrkSeg[0].TrkPt {
		if nLap < (len(wrk.Activity[0].Lap) - 1) {
			if ! trkPt.Time.Before(wrk.Activity[0].Lap[nLap + 1].StartTime) {
				nLap++
			}
		}
		if ! trkPt.Time.Before(wrk.Activity[0].Lap[nLap].StartTime) {
			newTrackpoint := tcx.Trackpoint{}
			newTrackpoint.Time = *trkPt.Time
			newTrackpoint.Position = new(tcx.Position)
			newTrackpoint.Position.Latitude = trkPt.Lat
			newTrackpoint.Position.Longitude = trkPt.Lon
			newTrackpoint.Altitude = new(float64)
			*newTrackpoint.Altitude = *trkPt.Ele
			wrk.Activity[0].Lap[nLap].Track[0].Trackpoint =
				append(wrk.Activity[0].Lap[nLap].Track[0].Trackpoint, newTrackpoint)
		}
	}
	// Sample original TCX for heart rate and cadence
	for nLap, lap := range wrk.Activity[0].Lap {
		for nTrackpoint, trackpoint := range lap.Track[0].Trackpoint {
			d := (float64)(10 * time.Hour)
			p := tcx.Trackpoint{}
			for _, origTrackpoint := range savedTrack[nLap] {
				if f := math.Abs((float64)(trackpoint.Time.Sub(origTrackpoint.Time))); f < d {
					d = f
					p = origTrackpoint
				}
			}
			if d < (float64)(5 * time.Second) {
				if p.HeartRate != nil {
					wrk.Activity[0].Lap[nLap].Track[0].Trackpoint[nTrackpoint].HeartRate = new(int)
					*wrk.Activity[0].Lap[nLap].Track[0].Trackpoint[nTrackpoint].HeartRate = *p.HeartRate
				}
				if p.Cadence != nil {
					wrk.Activity[0].Lap[nLap].Track[0].Trackpoint[nTrackpoint].Cadence = new(int)
					*wrk.Activity[0].Lap[nLap].Track[0].Trackpoint[nTrackpoint].Cadence = *p.Cadence
				}
			}
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
	merge(wrk, trk)
	if err = wrk.Save(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
