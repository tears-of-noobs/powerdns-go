package powerdns

const (
	TypeMX    string = "MX"
	TypeA     string = "A"
	TypeAAAA  string = "AAAA"
	TypeNS    string = "NS"
	TypeSOA   string = "SOA"
	TypeCNAME string = "CNAME"
	TypeTXT   string = "TXT"
	TypePTR   string = "PTR"
)

type Comment struct {
	Content    string `json:"content"`
	Account    string `json:"account"`
	ModifiedAt int64  `json:"modified_at"`
}

type Record struct {
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
}

type RRSet struct {
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	TTL        int64      `json:"ttl"`
	Records    []Record   `json:"records"`
	Comments   []Comment  `json:"comments"`
	Changetype changetype `json:"changetype,omitempty"`
}
