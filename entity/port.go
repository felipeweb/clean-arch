package entity

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// PortInfo is the information of a port.
type PortInfo struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

// Port is a port.
type Port struct {
	Key      string `json:"key,omitempty"`
	PortInfo `json:",inline"`
}

// UnmarshalJSON unmarshals a port.
func (p *Port) UnmarshalJSON(b []byte) error {
	m := make(map[string]PortInfo)
	if err := json.Unmarshal(b, &m); err != nil {
		return fmt.Errorf("failed to unmarshal port json: %w", err)
	}
	for k, v := range m {
		p.Key = k
		p.PortInfo = v
		break
	}
	return nil
}

// MarshalJSON marshals a port.
func (p Port) MarshalJSON() ([]byte, error) {
	m := map[string]PortInfo{p.Key: p.PortInfo}
	return json.Marshal(m)
}
