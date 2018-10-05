package goqrz

// QRZDatabase is documented in the standard.
type QRZDatabase struct {
	Xmlns    string   `xml:"xmlns,attr"`   // XML Namespace
	Version  string   `xml:"version,attr"` // QRZ API Version
	Session  Session  `xml:"Session"`      // User Session data
	Callsign Callsign `xml:"Callsign"`     // Requested Callsign data
	DXCC     DXCC     `xml:"DXCC"`         // Requested DXCC Zone data
}
