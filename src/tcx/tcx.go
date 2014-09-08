package tcx

import (
	"encoding/xml"
	"gpx"
	"os"
	"slf"
)

type TrainingCenterDatabase struct {
	xmlns string `xml:"xmlns,attr"`
}

func (t *TrainingCenterDatabase) FromLog(s *slf.Log) {
	// TODO
}

func (t *TrainingCenterDatabase) ReplaceTrack(track *gpx.Gpx) error {
	// TODO
}

func (t *TrainingCenterDatabase) Save(path string) error {
	t.xmlns = "http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2"
	var err error
	var file *os.File
	if file, err = os.Create(path); err != nil {
		return err
	}
	defer file.Close()
	if err = xml.NewEncoder(file).Encode(t); err != nil {
		return err
	}
	return nil
}
