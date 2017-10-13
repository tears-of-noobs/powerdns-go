package powerdns

type PowerDNSAPIError struct {
	Error  string   `json:"error"`
	Errors []string `json:"errors"`
}
