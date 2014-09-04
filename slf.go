package main

type GeneralInformation struct {
	SerialNumber string `xml:"serialNumber,attr"`
	Unit string `xml:"unit,attr"`
}

type Log struct {
	Revision int `xml:"revision,attr"`
	// TODO
}
