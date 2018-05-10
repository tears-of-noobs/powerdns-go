package powerdns

// Server - represent `Servers` resource
// https://doc.powerdns.com/md/httpapi/api_spec/#servers
type Server struct {
	ConfigURL  string `json:"config_url"`
	DaemonType string `json:"daemon_type"`
	ID         string `json:"id"`
	Type       string `json:"type"`
	URL        string `json:"url"`
	Version    string `json:"version"`
	ZonesURL   string `json:"zones_url"`
}
