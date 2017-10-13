package powerdns

type changetype string

const (
	ZoneKindNative string = "Native"
	ZoneKindMaster string = "Master"
	ZoneKindSlave  string = "Slave"

	ChangetypeReplace changetype = "REPLACE"
	ChangetypeDelete  changetype = "DELETE"
)

type BasicZoneInfo struct {
	Account        string   `json:"account,omitempty"`
	DNSSec         bool     `json:"dnssec,omitempty"`
	ID             string   `json:"id,omitempty"`
	Kind           string   `json:"kind,omitempty"`
	LastCheck      int64    `json:"last_check,omitempty"`
	Masters        []string `json:"masters,omitempty"`
	Name           string   `json:"string,omitempty"`
	NotifiedSerial int64    `json:"notified_serial,omitempty"`
	Serial         int64    `json:"serial,omitempty"`
	URL            string   `json:"url,omitempty"`
}

type Zone struct {
	BasicZoneInfo
	Nameservers []string `json:"nameservers,omitempty"`
	Servers     []string `json:"servers,omitempty"`
	RRSSets     []RRSet  `json:"rrsets,omitempty"`
	SOAEdit     string   `json:"soa_edit,omitempty"`
	SOAEditAPI  string   `json:"soa_edit_api,omitempty"`
}
