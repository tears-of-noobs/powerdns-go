package powerdns

// Config - represents config resource
// https://doc.powerdns.com/md/httpapi/api_spec/#config95setting95resource
type Config struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
