package app

import "net/http"

// Data whoami information.
type Data struct {
	Hostname   string            `json:"hostname,omitempty"`
	IP         []string          `json:"ip,omitempty"`
	Headers    http.Header       `json:"headers,omitempty"`
	URL        string            `json:"url,omitempty"`
	Host       string            `json:"host,omitempty"`
	Method     string            `json:"method,omitempty"`
	Name       string            `json:"name,omitempty"`
	RemoteAddr string            `json:"remoteAddr,omitempty"`
	Environ    map[string]string `json:"environ,omitempty"`
}
