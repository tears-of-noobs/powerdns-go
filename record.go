package powerdns

// RecordType - special type for represents type of record
type RecordType string

const (
	// TypeMX - MX record type
	TypeMX RecordType = "MX"

	// TypeA - A record type
	TypeA RecordType = "A"

	// TypeAAAA - AAAA record type
	TypeAAAA RecordType = "AAAA"

	// TypeNS - NS record type
	TypeNS RecordType = "NS"

	// TypeSOA - SOA record type
	TypeSOA RecordType = "SOA"

	// TypeCNAME - CNAME record type
	TypeCNAME RecordType = "CNAME"

	// TypeTXT - TXT record type
	TypeTXT RecordType = "TXT"

	// TypePTR - PTR record type
	TypePTR RecordType = "PTR"

	// TypeSRV - SRV record type
	TypeSRV RecordType = "SRV"
)

// Comment - represents commet block in RRSet definition
// https://doc.powerdns.com/md/httpapi/api_spec/#zones
type Comment struct {
	Content    string `json:"content"`
	Account    string `json:"account"`
	ModifiedAt int64  `json:"modified_at"`
}

// Record - represents record in  RRSet definition
// https://doc.powerdns.com/md/httpapi/api_spec/#zones
type Record struct {
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
}

// RRSet - represents RRSet which make up zone
// https://doc.powerdns.com/md/httpapi/api_spec/#zones
type RRSet struct {
	Name       string     `json:"name"`
	Type       RecordType `json:"type"`
	TTL        int64      `json:"ttl"`
	Records    []Record   `json:"records"`
	Comments   []Comment  `json:"comments"`
	Changetype changetype `json:"changetype,omitempty"`
}
