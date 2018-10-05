package goqrz

// Session is documented in the standard.
type Session struct {
	Key     string `xml:"Key"`     // a valid user session key
	Count   int    `xml:"Count"`   // Number of lookups performed by this user in the current 24 hour period
	SubExp  string `xml:"SubExp"`  // time and date that the users subscription will expire - or - "non-subscriber"
	GMTime  string `xml:"GMTime"`  // Time stamp for this message
	Message string `xml:"Message"` // An informational message for the user
	Error   string `xml:"Error"`   // XML system error message
	user    string // Username (Generally Callsign), private variable
	pass    string // Password, private variable
	agent   string // Agent ID, private variable
}
