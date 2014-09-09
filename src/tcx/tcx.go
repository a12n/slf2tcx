package tcx

import (
	"encoding/xml"
	"gpx"
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

func FromLog(log *slf.Log) (ans *TrainingCenterDatabase, err error) {
	ans = new(TrainingCenterDatabase)
	ans.Activity = append(ans.Activity, Activity{Sport: Biking})
	ans.Activity[0].Id = log.GeneralInformation.FileDate.Time
	ans.Activity[0].Lap = append(ans.Activity[0].Lap, ActivityLap{Intensity: Active, TriggerMethod: Manual})
	ans.Activity[0].Lap[0].Track = append(ans.Activity[0].Lap[0].Track, Track{})

	var t time.Time = log.GeneralInformation.StartDate.Time
	var k float64 = 0

	ans.Activity[0].Lap[0].StartTime = t

	ans.Activity[0].Lap[0].MaximumSpeed = new(float64)

	for _, entry := range log.LogEntry {
		var p Trackpoint
		p.Time = t
		p.Altitude = new(float64)
		*p.Altitude = (float64)(entry.Altitude) * 1.0E-3
		p.Distance = new(float64)
		*p.Distance = entry.Distance
		if len(ans.Activity[0].Lap[0].Track[0].Trackpoint) > 0 {
			*p.Distance += *ans.Activity[0].Lap[0].Track[0].Trackpoint[
				len(ans.Activity[0].Lap[0].Track[0].Trackpoint) - 1].Distance
		}
		p.HeartRate = new(int)
		*p.HeartRate = entry.Heartrate
		p.Cadence = new(int)
		*p.Cadence = entry.Cadence
		if (entry.Speed * 3600.0 / 1000.0) > *ans.Activity[0].Lap[0].MaximumSpeed {
			*ans.Activity[0].Lap[0].MaximumSpeed = (entry.Speed * 3600.0 / 1000.0)
		}
		k += entry.Calories
		ans.Activity[0].Lap[0].Track[0].Trackpoint = append(ans.Activity[0].Lap[0].Track[0].Trackpoint, p)
		t = t.Add((time.Duration)(entry.RideTime * (float64)(time.Second)))
		ans.Activity[0].Lap[0].TotalTime += entry.RideTime
	}
	ans.Activity[0].Lap[0].Calories = (int)(k)
	ans.Activity[0].Lap[0].Distance =
		*ans.Activity[0].Lap[0].Track[0].Trackpoint[len(ans.Activity[0].Lap[0].Track[0].Trackpoint) - 1].Distance

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
