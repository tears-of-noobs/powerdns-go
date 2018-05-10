package powerdns

// APIError - represents a error which returned from
// PowerDNS API
// https://doc.powerdns.com/md/httpapi/api_spec/#errors
type APIError struct {
	Error  string   `json:"error"`
	Errors []string `json:"errors"`
}
