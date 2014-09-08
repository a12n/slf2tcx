package slf

import (
	"encoding/xml"
	"os"
	"sort"
	"time"
)

type SigmaTime struct {
	time.Time
}

func (t *SigmaTime) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	var v string
	var x time.Time
	decoder.DecodeElement(&v, &start)
	if x, err = time.Parse("Mon Jan _2 15:04:05 GMT-0700 2006", v); err == nil {
		// Ignore time zone from the device, use local time zone
		t.Time = time.Date(
			x.Year(), x.Month(), x.Day(),
			x.Hour(), x.Minute(), x.Second(), x.Nanosecond(),
			time.Local)
	}
	return
}

// Sigma Log File, revision 3xx (Sigma Data Center v3.x)

const (
	// Gender
	Female = "female"
	Male = "male"
	// Weather in MetaInformation
	Cloudless = 0
	LightCloud = 1
	Cloudy = 2
	LightRain = 3
	Rain = 4
	Storm = 5
	Snow = 6
	Fog = 7
	// Wind in MetaInformation
	NoWind = 0
	LightWind = 1
	StrongWind = 2
	Gale = 3
	// Profile in MetaInformation
	Flat = 0
	SlightlyHilly = 1
	Hilly = 2
	Mountainous = 3
	Steep = 4
	// MarkerType in Marker
	Lap = "l"
	Pause = "p"
)

type GeneralInformation struct {
	SerialNumber int `xml:"serialNumber,attr"`
	Unit string `xml:"unit,attr"`
	LogType string `xml:"logType,attr"`
	// FileDate
	Name string
	StartDate SigmaTime
	// DateCode
	SamplingRate float64		// s
	WheelSize float64			// mm
	Measurement string
	Age int						// year
	Gender string
	BodyWeight float64
	BodyWeightUnit string
	Bike string
}

type Person struct {
	Color int `xml:"color,attr"`
	Gender string `xml:"gender,attr"`
	PersonName string
}

type Participant struct {
	Person []Person
}

type MetaInformation struct {
	Statistic bool
	Notes string
	Rating int
	Weather int
	Wind int
	Profile int
	Participant Participant
	TrainingType string
	ExternalLink string
	Temperature float64
}

type LogEntry struct {
	Number int
	Rotations int				// rpm?
	Speed float64				// m/s
	Heartrate int				// bpm
	Altitude float64			// mm
	Temperature float64			// °C
	RideTime float64			// s
	Distance float64			// m
	DistanceDownhill float64	// m
	DistanceUphill float64		// m
	AltitudeDifferenceDownhill float64 // mm?
	AltitudeDifferenceUphill float64   // mm?
	RideTimeUphill int				   // s
	RideTimeDownhill int			   // s
	Cadence int						   // rpm
	// IntensityZone
	// TargetZone
	Calories float64			// kcal
}

type LogValues struct {
	TrainingTime int			// s
	PauseTime int				// ?
	Distance float64			// m
	MinimumSpeed float64		// m/s
	AverageSpeed float64		// m/s
	MaximumSpeed float64		// m/s
	HrMax int					// bpm
	LowerLimit int				// bpm?
	UpperLimit int				// bpm?
	AverageHeartrate int		// bpm
	Calories float64			// kcal
	CaloriesDifferenceFactor float64 // ?
	IntensityZone1Start int		  // bpm
	IntensityZone2Start int		  // bpm
	IntensityZone3Start int		  // bpm
	IntensityZone4Start int		  // bpm
	TimeInIntensityZone1 int		  // s
	TimeInIntensityZone2 int		  // s
	TimeInIntensityZone3 int		  // s
	TimeInIntensityZone4 int		  // s
	TimeUnderIntensityZone int		  // s
	TimeOverIntensityZone int		  // s
	TargetZone string
	TimeInTargetZone int		// s
	TimeOverTargetZone int		// s
	TimeUnderTargetZone int		// s
	MinimumAltitude float64		// mm
	MaximumAltitude float64		// mm
	AverageAltitude float64		// mm
	RideTimeUphill int			// s
	RideTimeDownhill int		// s
	DistanceUphill float64		// m
	DistanceDownhill float64	// m
	AltitudeDifferencesUphill float64 // mm
	AltitudeDifferencesDownhill float64 // mm
}

type Marker struct {
	MarkerType string			// ?
	MarkerNumber int
	Title string
	Description string
	Time int					// s?
	TimeAbsolute int			// s?
	Duration int				// s?
	Distance float64			// m?
	DistanceAbsolute float64	// m?
	MinimumHeartrate int		// bpm?
	MaximumHeartrate int		// bpm?
	AverageHeartrate int		// bpm?
	Calories float64			// kcal?
	MinimumSpeed float64		// m/s?
	MaximumSpeed float64		// m/s?
	AverageSpeed float64		// m/s?
	AverageCadence int			// rpm?
	AverageAltitude float64		// mm?
	// Uphill
	// Downhill
}

type Log struct {
	Revision int `xml:"revision,attr"`
	GeneralInformation GeneralInformation
	MetaInformation MetaInformation
	// EncodedData
	LogValues LogValues
	LogEntry []LogEntry `xml:"LogEntries>LogEntry"`
	Marker []Marker `xml:"Markers>Marker"`
}

type byNumber []LogEntry
func (a byNumber) Len() int { return len(a) }
func (a byNumber) Less(i, j int) bool { return a[i].Number < a[j].Number }
func (a byNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type byTimeAbsolute []Marker
func (a byTimeAbsolute) Len() int { return len(a) }
func (a byTimeAbsolute) Less(i, j int) bool { return a[i].TimeAbsolute < a[j].TimeAbsolute }
func (a byTimeAbsolute) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Load(path string) (ans *Log, err error) {
	// FIXME: duplicates gpx.Load
	var file *os.File
	if file, err = os.Open(path); err == nil {
		defer file.Close()
		ans = new(Log)
		if err = xml.NewDecoder(file).Decode(ans); err == nil {
			sort.Sort(byNumber(ans.LogEntry))
			sort.Sort(byTimeAbsolute(ans.Marker))
		}
	}
	return
}
