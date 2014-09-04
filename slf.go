package main

type GeneralInformation struct {
	SerialNumber int `xml:"serialNumber,attr"`
	Unit string `xml:"unit,attr"`
	// TODO
}

type MetaInformation struct {
	Statistic bool
	Notes string
	Rating int
	Weather int
	Wind int
	Profile int
	// Participant
	TrainingType string
	ExternalLink string
	Temperature int
}

type LogEntry struct {
	IsPause bool
	PauseTime int
	IsWaypoint bool
	Title string
	Description string
	Rotations int
	RelativeRotations int
	Speed float64
	Heartrate int
	Altitude float64
	Temperature int
	RideTime float64
	Distance float64
	Incline float64
	RiseRate float64
	DistanceDownhill float64
	DistanceUphill float64
	AltitudeDifferenceDownhill float64
	AltitudeDifferenceUphill float64
	RideTimeUphill float64
	RideTimeDownhill float64
}

type LogValues struct {
	// TODO
}

type LogEntries struct {
	LogEntry []LogEntry
}

type Log struct {
	Revision int `xml:"revision,attr"`
	GeneralInformation GeneralInformation
	MetaInformation MetaInformation
	LogValues LogValues
	LogEntries LogEntries
}
