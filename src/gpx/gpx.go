package gpx

import (
	"encoding/xml"
	"os"
	"sort"
	"time"
)

type Email struct {
	Id string `xml:"id,attr"`
	Domain string `xml:"domain,attr"`
}

type Person struct {
	Name string `xml:"name,omitempty"`
	Email *Email `xml:"email,omitempty"`
	// Link *Link `xml:"link,omitempty"`
}

type Wpt struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
	Ele *float64 `xml:"ele,omitempty"`
	Time *time.Time `xml:"time,omitempty"`
	// MagVar *float64 `xml"magvar"`
	// GeoidHeight *float64 `xml"geoidheight"`
	Name string `xml:"name,omitempty"`
	Cmt string `xml:"cmt,omitempty"`
	Desc string `xml:"desc,omitempty"`
	// Link []Link `xml:"link"`
	Sym string `xml:"sym,omitempty"`
	Type string `xml:"type,omitempty"`
	// Fix *Fix `xml:"fix,omitempty"`
	// Sat *uint `xml:"sat,omitempty"`
	// Hdop *float64 `xml:"hdop,omitempty"`
	// Vdop *float64 `xml:"vdop,omitempty"`
	// Pdop *float64 `xml:"pdop,omitempty"`
	// AgeOfDgpsData *float64 `xml:"ageofdgpsdata,omitempty"`
	// DgpsId *DgpsId `xml:"dgpsid,omitempty"`
}

type TrkSeg struct {
	TrkPt []Wpt `xml:"trkpt,omitempty"`
}

type Trk struct {
	// Name string `xml:"name,omitempty"`
	// Cmt string `xml:"cmt,omitempty"`
	// Desc string `xml:"desc,omitempty"`
	// Src string `xml:"src,omitempty"`
	// Link []Link `xml:"link,omitempty"`
	// Number *uint `xml:"number,omitempty"`
	// Type string `xml:"type,omitempty"`
	TrkSeg []TrkSeg `xml:"trkseg,omitempty"`
}

type Metadata struct {
	Name string `xml:"name,omitempty"`
	Desc string `xml:"desc,omitempty"`
	Author *Person `xml:"author,omitempty"`
	// Copyright *Copyright `xml:"copyright,omitempty"`
	// Link []Link `xml:"link,omitempty"`
	Time *time.Time `xml:"time,omitempty"`
	// Keywords string `xml:"keywords,omitempty"`
	// Bounds *Bounds `xml:"bounds,omitempty"`
}

type Gpx struct {
	// Version string `xml:"version,attr"`
	// Creator string `xml:"creator,attr"`
	Metadata *Metadata `xml:"metadata,omitempty"`
	Wpt []Wpt `xml:"wpt,omitempty"`
	// Rte []Rte `xml:"rte,omitempty"`
	Trk []Trk `xml:"trk,omitempty"`
}

type byTime []Wpt

func (a byTime) Len() int { return len(a) }
func (a byTime) Less(i, j int) bool { return (a[i].Time != nil && a[j].Time != nil) && a[i].Time.Before(*a[j].Time) }
func (a byTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Load(path string) (ans *Gpx, err error) {
	// FIXME: duplicates slf.Load
	var file *os.File
	if file, err = os.Open(path); err == nil {
		defer file.Close()
		ans = new(Gpx)
		if err = xml.NewDecoder(file).Decode(ans); err == nil {
			for _, trk := range ans.Trk {
				for _, trkSeg := range trk.TrkSeg {
					sort.Sort(byTime(trkSeg.TrkPt))
				}
			}
		}
	}
	return
}
