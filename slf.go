package main

type GeneralInformation struct {
	SerialNumber int `xml:"serialNumber,attr"`
	Unit string `xml:"unit,attr"`
}

type MetaInformation struct {
	// TODO
}

type Log struct {
	Revision int `xml:"revision,attr"`
	GeneralInformation GeneralInformation `xml:"GeneralInformation"`
	MetaInformation MetaInformation `xml:"MetaInformation"`
	// TODO
}
