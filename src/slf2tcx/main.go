package main

import (
	// "fmt"
	"log"
	"os"
	"slf"
	"tcx"
	"time"
)

// Convert speed from m/s to km/h
func mps2kmph(v float64) float64 {
	return v * 3600 / 1000
}

// Convert length from mm to m
func mm2m(l float64) float64 {
	return l * 1.0E-3
}

// Convert Sigma Log File to Training Center Database
func conv(wrk *slf.Log, ans *tcx.TrainingCenterDatabase) (err error) {
	var curActivity tcx.Activity = tcx.Activity{Id: wrk.GeneralInformation.FileDate.Time, Sport: tcx.Biking}
	var curLap *tcx.ActivityLap
	var curTrack *tcx.Track

	var clockTime time.Time = wrk.GeneralInformation.StartDate.Time
	var samplingTime time.Duration = 0

	var nMarkerBegin int = 0

	for _, entry := range wrk.LogEntry {
		log.Printf("entry %#v\n", entry)

		if curLap == nil {
			curLap = new(tcx.ActivityLap)
			*curLap = tcx.ActivityLap{StartTime: clockTime, Track: make([]tcx.Track, 1)}
			curTrack = &curLap.Track[0]
		}

		var curTrackpoint tcx.Trackpoint

		curTrackpoint.Time = clockTime

		curTrackpoint.Altitude = new(float64)
		*curTrackpoint.Altitude = mm2m((float64)(entry.Altitude))

		curTrackpoint.Distance = new(float64)
		*curTrackpoint.Distance = entry.Distance
		if len(curTrack.Trackpoint) > 0 {
			*curTrackpoint.Distance += *curTrack.Trackpoint[len(curTrack.Trackpoint) - 1].Distance
		}

		curTrackpoint.HeartRate = new(int)
		*curTrackpoint.HeartRate = entry.HeartRate

		curTrackpoint.Cadence = new(int)
		*curTrackpoint.Cadence = entry.Cadence

		curTrack.Trackpoint = append(curTrack.Trackpoint, curTrackpoint)

		for nMarker := nMarkerBegin; nMarker < len(wrk.Marker); nMarker++ {
			curMarker := &wrk.Marker[nMarker]
			curMarkerTime := (time.Duration)(curMarker.TimeAbsolute) * time.Second
			curMarkerDuration := (time.Duration)(curMarker.Duration) * time.Second
			log.Printf("samplingTime %s, curMarker.TimeAbsolute %s\n",
				samplingTime.String(), curMarkerTime.String())
			if samplingTime >= curMarkerTime {
				if curMarker.MarkerType == slf.Pause {
					log.Printf("Pause at %f, duration %d\n", curMarker.DistanceAbsolute, curMarker.Duration)
					clockTime = clockTime.Add(curMarkerDuration)
					curLap.TotalTime += (float64)(curMarkerDuration / time.Second)
					nMarkerBegin = nMarker + 1
					break
				} else if curMarker.MarkerType == slf.Lap {
					// TODO: new lap, start time = (wrk.GeneralInformation.StartDate + Marker.TimeAbsolute)
					// Trackpoint was already appended, move it to the new lap.
				}
			}
		}

		var advanceTime time.Duration = (time.Duration)(wrk.GeneralInformation.SamplingRate) * time.Second
		if entry.RideTime > 0 {
			advanceTime = (time.Duration)(entry.RideTime) * time.Second
		}

		clockTime = clockTime.Add(advanceTime)
		samplingTime += advanceTime
		curLap.TotalTime += (float64)(advanceTime / time.Second)
	}

	// XXX: Close and append lap if there is no more log entries and
	// an unhandled lap marker.
	if nMarkerBegin == (len(wrk.Marker) - 1) {
		log.Printf("nMarkerBegin %d, len(wrk.Marker) %d\n",
			nMarkerBegin, len(wrk.Marker))
		for nMarker := nMarkerBegin; nMarker < len(wrk.Marker); nMarker++ {
			curMarker := &wrk.Marker[nMarker]
			if curMarker.MarkerType == slf.Lap {
				// Fill lap summary
				curLap.Distance = curMarker.DistanceAbsolute
				curLap.MaximumSpeed = new(float64)
				*curLap.MaximumSpeed = mps2kmph(curMarker.MaximumSpeed)
				curLap.Calories = (int)(curMarker.Calories)
				curLap.AverageHeartRate = new(int)
				*curLap.AverageHeartRate = curMarker.AverageHeartRate
				curLap.MaximumHeartRate = new(int)
				*curLap.MaximumHeartRate = curMarker.MaximumHeartRate
				curLap.Intensity = tcx.Active
				curLap.Cadence = new(int)
				*curLap.Cadence = curMarker.AverageCadence
				curLap.TriggerMethod = tcx.Manual
				// Append lap
				curActivity.Lap = append(curActivity.Lap, *curLap)
				curLap = nil
				break
			}
		}
	}

	ans.Activity = append(ans.Activity, curActivity)

	return
}

func main() {
	var ans *tcx.TrainingCenterDatabase = new(tcx.TrainingCenterDatabase)
	var wrk *slf.Log
	var err error
	if wrk, err = slf.Load(os.Stdin); err != nil {
		log.Fatal(err)
	}
	if err = conv(wrk, ans); err != nil {
		log.Fatal(err)
	}
	if err = ans.Save(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
