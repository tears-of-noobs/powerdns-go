package powerdns

import (
	"fmt"
	"strings"
)

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

// QueryRecordsByName - returns records filtered only by name
func (zone *Zone) QueryRecordsByName(
	name string,
	contains bool,
) []RRSet {
	var (
		rrSets []RRSet
	)

	for _, rrSet := range zone.RRSets {
		if contains {
			if strings.Contains(
				rrSet.Name,
				name,
			) {
				rrSets = append(rrSets, rrSet)
			}
		} else {
			if rrSet.Name == name {
				rrSets = append(rrSets, rrSet)
			}
		}

	}

	return rrSets
}

// QueryRecords - find and return records filtered by
// name or content
func (zone *Zone) QueryRecords(
	query string,
	contains bool,
) []RRSet {
	var (
		rrSets   []RRSet
		wildcard = len(query) == 0
	)

	for _, rrSet := range zone.RRSets {
		hasContent := false

		if wildcard {
			rrSets = append(rrSets, rrSet)
			continue
		}

		for _, record := range rrSet.Records {

			if contains {
				if strings.Contains(
					record.Content,
					query,
				) {
					rrSets = append(rrSets, rrSet)
					hasContent = true
				}
			} else {
				if query == record.Content {
					rrSets = append(rrSets, rrSet)
					hasContent = true
				}
			}

			if hasContent {
				break
			}

		}

		if contains {
			if strings.Contains(
				rrSet.Name,
				query,
			) {
				if !hasContent {
					rrSets = append(rrSets, rrSet)
				}
			}
		} else {
			if query == rrSet.Name {
				rrSets = append(rrSets, rrSet)
			}
		}

	}

	return rrSets
}

// GetSOARecord - returns SOA record of the zone
func (zone *Zone) GetSOARecord() (RRSet, error) {
	var (
		rrSets []RRSet
	)

	for _, rrSet := range zone.RRSets {
		if rrSet.Type == TypeSOA {
			rrSets = append(rrSets, rrSet)
		}
	}

	if len(rrSets) == 0 {
		return RRSet{}, fmt.Errorf(
			"can't find any %s record for zone %s",
			TypeSOA,
			zone.Name,
		)
	}

	if len(rrSets) > 1 {
		return RRSet{}, fmt.Errorf(
			"must be only the one %s record for zone %s, but %d exists",
			TypeSOA,
			zone.Name,
			len(rrSets),
		)
	}

	return rrSets[0], nil

}
