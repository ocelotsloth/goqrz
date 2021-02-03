package goqrz

// Callsign is documented in the standard.
type Callsign struct {
	Callsign         string `xml:"call"`      // callsign
	CrossRef         string `xml:"xref"`      // Cross reference: the query callsign that returned this record
	Aliases          string `xml:"aliases"`   // Other callsigns that resolve to this record
	DXCCID           string `xml:"dxcc"`      // DXCC entity ID (country code) for the callsign
	FirstName        string `xml:"fname"`     // first name
	LastName         string `xml:"name"`      // last name
	Address1         string `xml:"addr1"`     // address line 1 (i.e. house # and street)
	Address2         string `xml:"addr2"`     // address line 2 (i.e, city name)
	State            string `xml:"state"`     // state (USA Only)
	Zip              string `xml:"zip"`       // Zip/postal code
	CountryName      string `xml:"country"`   // country name for the QSL mailing address
	DXCCEntityCode   string `xml:"ccode"`     // dxcc entity code for the mailing address country
	Lat              string `xml:"lat"`       // lattitude of address (signed decimal) S < 0 > N
	Lon              string `xml:"lon"`       // longitude of address (signed decimal) W < 0 > E
	GridLocator      string `xml:"grid"`      // grid locator
	County           string `xml:"county"`    // county name (USA)
	FIPSIdentifier   string `xml:"fips"`      // FIPS county identifier (USA)
	DXCCCountryName  string `xml:"land"`      // DXCC country name of the callsign
	EffectiveDate    string `xml:"efdate"`    // license effective date (USA)
	ExpirationDate   string `xml:"expdate"`   // license expiration date (USA)
	PreviousCall     string `xml:"p_call"`    // previous callsign
	Class            string `xml:"class"`     // license class
	Codes            string `xml:"codes"`     // license type codes (USA)
	QSLManager       string `xml:"qslmgr"`    // QSL manager info
	Email            string `xml:"email"`     // email address
	WebURL           string `xml:"url"`       // web page address
	QRZPageViews     string `xml:"u_views"`   // QRZ web page views
	BioByteLength    string `xml:"bio"`       // approximate length of the bio HTML in bytes
	LastBioUpdate    string `xml:"biodate"`   // date of the last bio update
	ImageURL         string `xml:"image"`     // full URL of the callsign's primary image
	ImageInfo        string `xml:"imageinfo"` // height:width:size in bytes, of the image file
	Serial           string `xml:"serial"`    // QRZ db serial number
	LastModDate      string `xml:"moddate"`   // QRZ callsign last modified date
	MetroServiceArea string `xml:"MSA"`       // Metro Service Area (USPS)
	TelAreaCode      string `xml:"AreaCode"`  // Telephone Area Code (USA)
	TimeZone         string `xml:"TimeZone"`  // Time Zone (USA)
	GMTOffset        string `xml:"GMTOffset"` // GMT Time Offset
	DSTObserved      string `xml:"DST"`       // Daylight Saving Time Observed
	EQSLAccepted     string `xml:"eqsl"`      // Will accept e-qsl (0/1 or blank if unknown)
	PaperQSLAccepted string `xml:"mqsl"`      // Will return paper QSL (0/1 or blank if unknown)
	CQZone           string `xml:"cqzone"`    // CQ Zone identifier
	ITUZone          string `xml:"ituzone"`   // ITU Zone identifier
	BirthDate        string `xml:"born"`      // operator's year of birth
	QRZCallManager   string `xml:"user"`      // User who manages this callsign on QRZ
	LOTWAccepted     string `xml:"lotw"`      // Will accept LOTW (0/1 or blank if unknown)
	IOTADesignator   string `xml:"iota"`      // IOTA Designator (blank if unknown)
	GeolocSource     string `xml:"geoloc"`    // Describes source of lat/long data
	AttentionAddress string `xml:"attn"`      // Attention address line, this line should be prepended to the address
	Nickname         string `xml:"nickname"`  // A different or shortened name used on the air
	NameFormat       string `xml:"name_fmt"`  // Combined full name and nickname in the format used by QRZ. This fortmat is subject to change.
}
