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
	Activities []Activity `xml:"Activities>Activity"`
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

func FromLog(log *slf.Log) (*TrainingCenterDatabase, error) {
	// TODO
	return new(TrainingCenterDatabase), nil
}

func (t *TrainingCenterDatabase) ReplaceTrack(track *gpx.Gpx) error {
	// TODO
	return nil
}

func (t *TrainingCenterDatabase) Save(path string) (err error) {
	var file *os.File
	if file, err = os.Create(path); err != nil {
		defer file.Close()
		var encoder *xml.Encoder = xml.NewEncoder(file)
		encoder.Indent("", "\t")
		err = encoder.Encode(t)
	}
	return
}
