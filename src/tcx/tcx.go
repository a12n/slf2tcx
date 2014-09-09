package tcx

import (
	"encoding/xml"
	"gpx"
	"log"
	"os"
	"slf"
	"time"
)

const (
	// Intensity field in ActivityLap
	Active = "Active"
	Resting = "Resting"
	// TriggerMethod field in ActivityLap
	Manual = "Manual"
	Distance = "Distance"
	Location = "Location"
	Time = "Time"
	HeartRate = "HeartRate"
	// SensorState field in Trackpoint
	Present = "Present"
	Absent = "Absent"
	// Sport attr in Activity
	Running = "Running"
	Biking = "Biking"
	Other = "Other"
)

type Position struct {
	Latitude float64 `xml:"LatitudeDegrees"`
	Longitude float64 `xml:"LongitudeDegrees"`
}

type Trackpoint struct {
	Time time.Time
	Position *Position `xml:",omitempty"`
	Altitude *float64 `xml:"AltitudeMeters,omitempty"`
	Distance *float64 `xml:"DistanceMeters,omitempty"`
	HeartRate *int `xml:"HeartRateBpm>Value,omitempty"`
	Cadence *int `xml:",omitempty"`
	SensorState string `xml:",omitempty"`
}

type Track struct {
	Trackpoint []Trackpoint
}

type ActivityLap struct {
	StartTime time.Time `xml:",attr"`
	TotalTime float64 `xml:"TotalTimeSeconds"`
	Distance float64 `xml:"DistanceMeters"`
	MaximumSpeed *float64 `xml:",omitempty"`
	Calories int
	AverageHeartRate *int `xml:"AverageHeartRateBpm>Value,omitempty"`
	MaximumHeartRate *int `xml:"MaximumHeartRateBpm>Value,omitempty"`
	Intensity string
	Cadence *int `xml:",omitempty"`
	TriggerMethod string
	Track []Track `xml:",omitempty"`
	Notes string `xml:",omitempty"`
}

type Activity struct {
	Sport string `xml:",attr"`
	Id time.Time
	Lap []ActivityLap
}

type TrainingCenterDatabase struct {
	XMLName xml.Name `xml:"http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2 TrainingCenterDatabase"`
	Activity []Activity `xml:"Activities>Activity"`
}

// func Load(path string) (*TrainingCenterDatabase, error) {
// 	var err error
// 	var file *os.File
// 	if file, err = os.Open(path); err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	var ans *TrainingCenterDatabase = new(TrainingCenterDatabase)
// 	var decoder *xml.Decoder = xml.NewDecoder(file)
// 	decoder.DefaultSpace = "http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2"
// 	if err = decoder.Decode(ans); err != nil {
// 		return nil, err
// 	}
// 	return ans, nil
// }

func FromLog(wrk *slf.Log) (ans *TrainingCenterDatabase, err error) {
	ans = new(TrainingCenterDatabase)

	var activity Activity = Activity{Id: wrk.GeneralInformation.FileDate.Time, Sport: Biking}

	var lap *ActivityLap
	var t time.Time = wrk.GeneralInformation.StartDate.Time
	var kcal float64 = 0

	var m int = 0

	for _, entry := range wrk.LogEntry {
		if lap == nil {
			lap = new(ActivityLap)
			*lap = ActivityLap{Intensity: Active, MaximumSpeed:
				new(float64), StartTime: t, Track: make([]Track, 1),
				TriggerMethod: Manual}
		}

		var point Trackpoint

		point.Time = t

		point.Altitude = new(float64)
		*point.Altitude = (float64)(entry.Altitude) * 1.0E-3

		point.Distance = new(float64)
		*point.Distance = entry.Distance

		if len(lap.Track[0].Trackpoint) > 0 {
			*point.Distance += *lap.Track[0].Trackpoint[len(lap.Track[0].Trackpoint) - 1].Distance
		}

		point.HeartRate = new(int)
		*point.HeartRate = entry.Heartrate

		point.Cadence = new(int)
		*point.Cadence = entry.Cadence

		if v := (entry.Speed * 3600.0 / 1000.0); v > *lap.MaximumSpeed {
			*lap.MaximumSpeed = v
		}

		kcal += entry.Calories

		lap.Track[0].Trackpoint = append(lap.Track[0].Trackpoint, point)

		lap.TotalTime += entry.RideTime
		t = t.Add((time.Duration)(entry.RideTime * (float64)(time.Second)))

		for i := m; i < len(wrk.Marker); i++ {
			tm := wrk.GeneralInformation.StartDate.Time.Add((time.Duration)(wrk.Marker[i].TimeAbsolute) * time.Second)
			if t.After(tm) {
				if wrk.Marker[i].MarkerType == slf.Pause {
					log.Printf("Pause at %f, duration %d\n", wrk.Marker[i].DistanceAbsolute, wrk.Marker[i].Duration)
					t = t.Add((time.Duration)(wrk.Marker[i].Duration) * time.Second)
					lap.TotalTime += (float64)(wrk.Marker[i].Duration)
				}
				m = i + 1
				break
			}
		}
	}
	lap.Calories = (int)(kcal)
	lap.Distance = *lap.Track[0].Trackpoint[len(lap.Track[0].Trackpoint) - 1].Distance

	if lap != nil {
		activity.Lap = append(activity.Lap, *lap)
	}
	ans.Activity = append(ans.Activity, activity)

	return
}

func (t *TrainingCenterDatabase) ReplaceTrack(track *gpx.Gpx) error {
	// TODO
	return nil
}

func (t *TrainingCenterDatabase) Save(path string) (err error) {
	var file *os.File
	if file, err = os.Create(path); err == nil {
		defer file.Close()
		var encoder *xml.Encoder = xml.NewEncoder(file)
		encoder.Indent("", "\t")
		err = encoder.Encode(t)
	}
	return
}
