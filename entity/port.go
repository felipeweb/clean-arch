package entity

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// PortInfo is the information of a port.
type PortInfo struct {
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates" gorm:"type:numeric[]"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias" gorm:"type:varchar[]"`
	Regions     []string  `json:"regions" gorm:"type:varchar[]"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs" gorm:"type:varchar[]"`
	Code        string    `json:"code"`
}

// Port is a port.
type Port struct {
	Key      string `json:"key,omitempty" gorm:"primaryKey"`
	PortInfo `json:",inline" gorm:"embedded"`
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
