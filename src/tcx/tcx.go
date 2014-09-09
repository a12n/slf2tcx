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
	var t2 time.Time = wrk.GeneralInformation.StartDate.Time

	var m int = 0

	for k, entry := range wrk.LogEntry {
		log.Printf("entry %d\n", k)

		if lap == nil {
			lap = new(ActivityLap)
			*lap = ActivityLap{StartTime: t, Track: make([]Track, 1)}
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

		lap.Track[0].Trackpoint = append(lap.Track[0].Trackpoint, point)

		t = t.Add((time.Duration)(wrk.GeneralInformation.SamplingRate * (float64)(time.Second)))
		lap.TotalTime += wrk.GeneralInformation.SamplingRate

		for i := m; i < len(wrk.Marker); i++ {
			// log.Printf("wrk.Marker[%d] %#v\n", i, wrk.Marker[i])
			tm := wrk.GeneralInformation.StartDate.Time.Add((time.Duration)(wrk.Marker[i].TimeAbsolute) * time.Second)
			log.Printf("t %v, t2 %v, tm[%d] %v\n", t, t2, i, tm)
			if t2.After(tm) {
			// log.Printf("point.Distance %f, wrk.Marker[%d].DistanceAbsolute %f\n",
			// 	*point.Distance, i, wrk.Marker[i].DistanceAbsolute)
			// if *point.Distance >= wrk.Marker[i].DistanceAbsolute {
				if wrk.Marker[i].MarkerType == slf.Lap {
					// lap.TotalTime = (float64)(wrk.Marker[i].Time)
					lap.Distance = wrk.Marker[i].Distance
					lap.MaximumSpeed = new(float64)
					*lap.MaximumSpeed = wrk.Marker[i].MaximumSpeed * 3600.0 / 1000.0
					lap.Calories = (int)(wrk.Marker[i].Calories)
					lap.AverageHeartRate = new(int)
					*lap.AverageHeartRate = wrk.Marker[i].AverageHeartrate
					lap.MaximumHeartRate = new(int)
					*lap.MaximumHeartRate = wrk.Marker[i].MaximumHeartrate
					lap.Intensity = Active
					lap.Cadence = new(int)
					*lap.Cadence = wrk.Marker[i].AverageCadence
					lap.TriggerMethod = Manual
					log.Printf("Append lap\n")
					activity.Lap = append(activity.Lap, *lap)
					lap = nil
				} else if wrk.Marker[i].MarkerType == slf.Pause {
					log.Printf("Pause at %f, duration %d\n", wrk.Marker[i].DistanceAbsolute, wrk.Marker[i].Duration)
					t = t.Add((time.Duration)(wrk.Marker[i].Duration) * time.Second)
				}
				m = i + 1
				break
			}
		}

		t2 = t2.Add((time.Duration)(wrk.GeneralInformation.SamplingRate * (float64)(time.Second)))
	}

	// if lap != nil {
	// 	log.Printf("Append lap\n")
	// 	activity.Lap = append(activity.Lap, *lap)
	// 	lap = nil
	// }

	ans.Activity = append(ans.Activity, activity)

	return
}

func lerp(t, a, b float64) float64 {
	return (1.0 - t) * a + t * b
}

func (self *TrainingCenterDatabase) ReplaceTrack(track *gpx.Gpx) error {
	k := 0
	for iLap, lap := range self.Activity[0].Lap {
		for iTrackpoint, point := range lap.Track[0].Trackpoint {
			m := -1
			for i := k; i < len(track.Trk[0].TrkSeg[0].TrkPt); i++ {
				// log.Printf("TrkPt[%d] %#v\n", i, track.Trk[0].TrkSeg[0].TrkPt[i])
				if point.Time.Before(*track.Trk[0].TrkSeg[0].TrkPt[i].Time) {
					m = i
					break
				}
			}
			if m > 0 {
				q1 := track.Trk[0].TrkSeg[0].TrkPt[m - 1]
				q2 := track.Trk[0].TrkSeg[0].TrkPt[m]
				dq := q2.Time.Sub(*q1.Time)
				dp := point.Time.Sub(*q1.Time)
				t := (float64)(dp) / (float64)(dq)
				log.Printf("dq %d, dp %d, t %f\n", dq, dp, t)
				self.Activity[0].Lap[iLap].Track[0].Trackpoint[iTrackpoint].Position = new(Position)
				self.Activity[0].Lap[iLap].Track[0].Trackpoint[iTrackpoint].Position.Latitude = lerp(t, q1.Lat, q2.Lat)
				self.Activity[0].Lap[iLap].Track[0].Trackpoint[iTrackpoint].Position.Longitude = lerp(t, q1.Lon, q2.Lon)
			}
		}
	}
	return nil
}

func (t *TrainingCenterDatabase) SaveFile(path string) (err error) {
	var file *os.File
	if file, err = os.Create(path); err == nil {
		defer file.Close()
		err = t.Save(file)
	}
	return
}

func (t *TrainingCenterDatabase) Save(file *os.File) (err error) {
	var encoder *xml.Encoder = xml.NewEncoder(file)
	encoder.Indent("", "\t")
	if _, err = file.WriteString(xml.Header); err == nil {
		err = encoder.Encode(t)
	}
	return
}
