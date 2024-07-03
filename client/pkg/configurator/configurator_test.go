package configurator

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		path := strings.Split(fmt.Sprintf("%v", req.URL), "/")
		if len(path) == 2 && path[1] == "hostname" {
			if req.Method == http.MethodGet {
				resp := hostResponse{Hostname: "hostname"}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
			if req.Method == http.MethodPost {
				var reqBody hostResponse
				_ = json.NewDecoder(req.Body).Decode(&reqBody)
				resp := hostResponse{Hostname: reqBody.Hostname}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
		}
		if len(path) == 2 && path[1] == "dns-servers" {
			if req.Method == http.MethodGet {
				resp := dnsResponse{Servers: []string{"1.1.1.1", "2.2.2.2"}}
				_ = json.NewEncoder(w).Encode(resp)
				return
			}
		}
		if len(path) == 3 && path[1] == "dns-servers" {
			if req.Method == http.MethodPut || req.Method == http.MethodDelete {
				if net.ParseIP(path[2]) != nil {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotImplemented)
					resp := errResponse{Message: "invalid DNS server"}
					_ = json.NewEncoder(w).Encode(resp)
				}
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		resp := errResponse{Message: "Not Found"}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	return server
}

func TestConfigurator_AddServer(t *testing.T) {
	server := newTestServer()
	client := http.DefaultClient
	url := server.URL

	tests := []struct {
		name     string
		address  string
		apiUrl   string
		expected string
	}{
		{
			name:     "OK",
			address:  "1.1.1.1",
			apiUrl:   url,
			expected: "<nil>",
		},
		{
			name:     "Invalid address",
			address:  "1.1.1..1",
			apiUrl:   url,
			expected: "error getting servers: invalid DNS server",
		},
		{
			name:     "No URL",
			address:  "1.1.1.1",
			expected: `error while sending request occurred: Put "/dns-servers/1.1.1.1": unsupported protocol scheme ""`,
		},
		{
			name:     "Invalid URL",
			address:  "1.1.1.1",
			apiUrl:   url + "/url",
			expected: "error getting servers: Not Found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Configurator{
				httpClient: *client,
				apiUrl:     tt.apiUrl,
			}
			err := c.AddServer(tt.address)
			if (err != nil && fmt.Sprint(err) != tt.expected) || err == nil && tt.expected != "<nil>" {
				t.Errorf("AddServer() error = '%v', expected '%v'", err, tt.expected)
			}
		})
	}
}

func TestConfigurator_DeleteServer(t *testing.T) {
	server := newTestServer()
	client := http.DefaultClient
	url := server.URL

	tests := []struct {
		name     string
		address  string
		apiUrl   string
		expected string
	}{
		{
			name:     "OK",
			address:  "1.1.1.1",
			apiUrl:   url,
			expected: "<nil>",
		},
		{
			name:     "Invalid address",
			address:  "1.1.1..1",
			apiUrl:   url,
			expected: "error getting servers: invalid DNS server",
		},
		{
			name:     "No URL",
			address:  "1.1.1.1",
			expected: `error while sending request occurred: Delete "/dns-servers/1.1.1.1": unsupported protocol scheme ""`,
		},
		{
			name:     "Invalid URL",
			address:  "1.1.1.1",
			apiUrl:   url + "/url",
			expected: "error getting servers: Not Found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Configurator{
				httpClient: *client,
				apiUrl:     tt.apiUrl,
			}
			err := c.DeleteServer(tt.address)
			if (err != nil && fmt.Sprint(err) != tt.expected) || err == nil && tt.expected != "<nil>" {
				t.Errorf("AddServer() error = '%v', expected '%v'", err, tt.expected)
			}
		})
	}
}
