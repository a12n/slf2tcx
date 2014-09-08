package tcx

import (
	"gpx"
	"slf"
)

type TrainingCenterDatabase struct {
	xmlns string `xml:"xmlns,attr"`
}

func (t *TrainingCenterDatabase) FromLog(s *slf.Log) {
	// TODO
}

func (t *TrainingCenterDatabase) ReplaceTrack(s *gpx.Gpx) {
	// TODO
}

func (t *TrainingCenterDatabase) Save(path string) error {
	t.xmlns = "http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2"
	// TODO
	return nil
}
