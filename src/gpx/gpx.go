package gpx

import (
	"encoding/xml"
	"os"
	"time"
)

type Email struct {
	Id string `xml:"id,attr"`
	Domain string `xml:"domain,attr"`
}

type Person struct {
	Name *string `xml:"name"`
	Email *Email `xml:"email"`
	// Link *Link `xml:"link"`
}

type Wpt struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
	Ele *float64 `xml:"ele"`
	Time *time.Time `xml"time"`
	// MagVar *float64 `xml"magvar"`
	// GeoidHeight *float64 `xml"geoidheight"`
	Name *string `xml:"name"`
	Cmt *string `xml:"cmt"`
	Desc *string `xml:"desc"`
	// Link []Link `xml:"link"`
	Sym *string `xml:"sym"`
	Type *string `xml:"type"`
	// Fix *Fix `xml:"fix"`
	// Sat *uint `xml:"sat"`
	// Hdop *float64 `xml:"hdop"`
	// Vdop *float64 `xml:"vdop"`
	// Pdop *float64 `xml:"pdop"`
	// AgeOfDgpsData *float64 `xml:"ageofdgpsdata"`
	// DgpsId *DgpsId `xml:"dgpsid"`
}

type TrkSeg struct {
	TrkPt []*Wpt `xml:"trkpt"`
}

type Trk struct {
	// Name *string `xml:"name"`
	// Cmt *string `xml:"cmt"`
	// Desc *string `xml:"desc"`
	// Src *string `xml:"src"`
	// Link []Link `xml:"link"`
	// Number *uint `xml:"number"`
	// Type *string `xml:"type"`
	TrkSeg []*TrkSeg `xml:"trkseg"`
}

type Metadata struct {
	Name *string `xml:"name"`
	Desc *string `xml:"desc"`
	Author *Person `xml:"author"`
	// Copyright *Copyright `xml:"copyright"`
	// Link []Link `xml:"link"`
	Time *time.Time `xml:"time"`
	// Keywords *string `xml:"keywords"`
	// Bounds *Bounds `xml:"bounds"`
}

type Gpx struct {
	// Version string `xml:"version,attr"`
	// Creator string `xml:"creator,attr"`
	Metadata *Metadata `xml:"metadata"`
	Wpt []*Wpt `xml:"wpt"`
	// Rte []*Rte `xml:"rte"`
	Trk []*Trk `xml:"trk"`
}

func Load(path string) (ans *Gpx, err error) {
	// FIXME: duplicates slf.Load
	var file *os.File
	if file, err = os.Open(path); err == nil {
		defer file.Close()
		ans = new(Gpx)
		err = xml.NewDecoder(file).Decode(ans)
	}
	return
}
