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

func FromLog(log *slf.Log) (*TrainingCenterDatabase, error) {
	// TODO
	return new(TrainingCenterDatabase), nil
}

func (t *TrainingCenterDatabase) ReplaceTrack(track *gpx.Gpx) error {
	// TODO
	return nil
}

func (t *TrainingCenterDatabase) Save(path string) error {
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
