package configurator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Configurator struct {
	httpClient http.Client
	apiUrl     string
}

func NewConfigurator(url string) *Configurator {
	return &Configurator{httpClient: http.Client{Timeout: time.Second * 5}, apiUrl: url}
}

func (c Configurator) SetHostname(hostname string) error {
	body := []byte(`{"hostname":"` + hostname + `"}`)
	apiReq, err := http.NewRequest("POST", c.apiUrl+"/hostname", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error while creating request occurred: %w", err)
	}

	apiReq.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(apiReq)
	if err != nil {
		return fmt.Errorf("error while sending request occurred: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respBody struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			return fmt.Errorf("error getting response: %w", err)
		}
		return fmt.Errorf("error changing hostname: %s", respBody.Message)
	} else {
		var respBody struct {
			Hostname string `json:"hostname"`
		}
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			return fmt.Errorf("error getting response: %w", err)
		}
		fmt.Println("Hostname changed to " + respBody.Hostname)
	}
	return nil
}
