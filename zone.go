package powerdns

type changetype string

const (
	// ZoneKindNative - Native kind of zone
	ZoneKindNative string = "Native"

	// ZoneKindMaster - Master kind of zone
	ZoneKindMaster string = "Master"

	// ZoneKindSlave - Slave kind of zone
	ZoneKindSlave string = "Slave"

	// ChangetypeReplace - type of operation - REPLACE
	ChangetypeReplace changetype = "REPLACE"

	// ChangetypeDelete - type of operation - DELETE
	ChangetypeDelete changetype = "DELETE"
)

// BasicZoneInfo - reduced representation of zone
// without RRSet
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

// Zone - represents zone
// https://doc.powerdns.com/md/httpapi/api_spec/#zones
type Zone struct {
	BasicZoneInfo
	Nameservers []string `json:"nameservers,omitempty"`
	Servers     []string `json:"servers,omitempty"`
	RRSets      []RRSet  `json:"rrsets,omitempty"`
	SOAEdit     string   `json:"soa_edit,omitempty"`
	SOAEditAPI  string   `json:"soa_edit_api,omitempty"`
}

// GetRecords - find and return records filtered by
// name or recordType or content
func (zone *Zone) GetRecords(
	name string,
	recordType RecordType,
	content []string,
) []RRSet {
	var rrSets []RRSet

	for _, rrSet := range zone.RRSets {
		var (
			hasName    = true
			hasType    = true
			hasContent = true
		)

		if name != "" {
			hasName = name == rrSet.Name
		}

		if recordType != "" {
			hasType = recordType == rrSet.Type
		}

		if len(content) != 0 {
			for _, record := range rrSet.Records {
				for _, c := range content {
					if c == record.Content {
						hasContent = true
						break
					}
				}

				if hasContent {
					break
				}
			}

		}

		if hasName && hasType && hasContent {

			rrSets = append(rrSets, rrSet)
		}
	}

	return rrSets
}
