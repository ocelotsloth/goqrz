package goqrz

// DXCC is defined by the following standard:
// https://www.qrz.com/XML/current_spec.html
type DXCC struct {
	Dxcc      string `xml:"dxcc"`      // DXCC entity number for this record
	Cc        string `xml:"cc"`        // 2-letter country code (ISO-3166)
	Ccc       string `xml:"ccc"`       // 3-letter country code (ISO-3166)
	Name      string `xml:"name"`      // long name
	Continent string `xml:"continent"` // 2-letter continent designator
	Ituzone   string `xml:"ituzone"`   // ITU Zone
	Cqzone    string `xml:"cqzone"`    // CQ Zone
	Timezone  string `xml:"timezone"`  // UTC timezone offset +/-
	Lat       string `xml:"lat"`       // Latitude (approx.)
	Lon       string `xml:"lon"`       // Longitude (approx.)
	Notes     string `xml:"notes"`     // Special notes and/or exceptions
}
