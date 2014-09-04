package main

type GeneralInformation struct {
	SerialNumber int `xml:"serialNumber,attr"`
	Unit string `xml:"unit,attr"`
	// TODO
}

type MetaInformation struct {
	// TODO
}

type LogEntry struct {
	IsPause bool `xml:"IsPause"`
	PauseTime int `xml:"PauseTime"`
	IsWaypoint bool `xml:"IsWaypoint"`
	Title string `xml:"Title"`
	Description string `xml:"Description"`
	Rotations int `xml:"Rotations"`
	RelativeRotations int `xml:"RelativeRotations"`
	Speed float64 `xml:"Speed"`
	Heartrate int `xml:"Heartrate"`
	Altitude float64 `xml:"Altitude"`
	Temperature int `xml:"Temperature"`
	RideTime float64 `xml:"RideTime"`
	Distance float64 `xml:"Distance"`
	Incline float64 `xml:"Incline"`
	RiseRate float64 `xml:"RiseRate"`
	DistanceDownhill float64 `xml:"DistanceDownhill"`
	DistanceUphill float64 `xml:"DistanceUphill"`
	AltitudeDifferenceDownhill float64 `xml:"AltitudeDifferenceDownhill"`
	AltitudeDifferenceUphill float64 `xml:"AltitudeDifferenceUphill"`
	RideTimeUphill float64 `xml:"RideTimeUphill"`
	RideTimeDownhill float64 `xml:"RideTimeDownhill"`
}

type LogValues struct {
	// TODO
}

type LogEntries struct {
	LogEntry []LogEntry `xml:"LogEntry"`
}

type Log struct {
	Revision int `xml:"revision,attr"`
	GeneralInformation GeneralInformation `xml:"GeneralInformation"`
	MetaInformation MetaInformation `xml:"MetaInformation"`
	LogValues LogValues `xml:"LogValues"`
	LogEntries LogEntries `xml:"LogEntries"`
}
