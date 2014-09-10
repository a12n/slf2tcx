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
	k := 0
	for nLap, lap := range wrk.Activity[0].Lap {
		for nWrkPt, wrkPt := range lap.Track[0].Trackpoint {
			m := -1
			for i := k; i < len(trk.Trk[0].TrkSeg[0].TrkPt); i++ {
				// log.Printf("TrkPt[%d] %#v\n", i, track.Trk[0].TrkSeg[0].TrkPt[i])
				if wrkPt.Time.Before(*trk.Trk[0].TrkSeg[0].TrkPt[i].Time) {
					m = i
					break
				}
			}
			if m > 0 {
				q1 := trk.Trk[0].TrkSeg[0].TrkPt[m - 1]
				q2 := trk.Trk[0].TrkSeg[0].TrkPt[m]
				dq := q2.Time.Sub(*q1.Time)
				dp := wrkPt.Time.Sub(*q1.Time)
				t := (float64)(dp) / (float64)(dq)
				// log.Printf("dq %d, dp %d, t %f\n", dq, dp, t)
				ans.Activity[0].Lap[nLap].Track[0].Trackpoint[nWrkPt].Position = new(tcx.Position)
				ans.Activity[0].Lap[nLap].Track[0].Trackpoint[nWrkPt].Position.Latitude = lerp(t, q1.Lat, q2.Lat)
				ans.Activity[0].Lap[nLap].Track[0].Trackpoint[nWrkPt].Position.Longitude = lerp(t, q1.Lon, q2.Lon)
				ans.Activity[0].Lap[nLap].Track[0].Trackpoint[nWrkPt].Altitude = new(float64)
				*ans.Activity[0].Lap[nLap].Track[0].Trackpoint[nWrkPt].Altitude = lerp(t, *q1.Ele, *q2.Ele)
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
	if err = merge(wrk, trk).Save(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
