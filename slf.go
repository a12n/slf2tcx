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
	IsPause boolean `xml:"IsPause"`
	PauseTime int `xml:"PauseTime"`
	IsWaypoint boolean `xml:"IsWaypoint"`
	Title string `xml:"Title"`
	Description string `xml:"Description"`
	Rotations int `xml:"Rotations"`
	RelativeRotations int `xml:"RelativeRotations"`
	Speed double `xml:"Speed"`
	Heartrate int `xml:"Heartrate"`
	Altitude int `xml:"Altitude"`
	Temperature int `xml:"Temperature"`
	RideTime double `xml:"RideTime"`
	Distance double `xml:"Distance"`
	Incline double `xml:"Incline"`
	RiseRate double `xml:"RiseRate"`
	DistanceDownhill double `xml:"DistanceDownhill"`
	DistanceUphill double `xml:"DistanceUphill"`
	AltitudeDifferenceDownhill double `xml:"AltitudeDifferenceDownhill"`
	AltitudeDifferenceUphill double `xml:"AltitudeDifferenceUphill"`
	RideTimeUphill double `xml:"RideTimeUphill"`
	RideTimeDownhill double `xml:"RideTimeDownhill"`
}

type LogValues struct {
	LogEntry []*LogEntry `xml:"LogEntry"`
}

type Log struct {
	Revision int `xml:"revision,attr"`
	GeneralInformation GeneralInformation `xml:"GeneralInformation"`
	MetaInformation MetaInformation `xml:"MetaInformation"`
	LogValues LogValues `xml:"LogValues"`
}
