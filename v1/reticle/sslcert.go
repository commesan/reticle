package reticle

import (
	"encoding/json"
	"fmt"
)

// UnmarshalSSLGetCert parses the JSON-encoded data and stores the result in the value pointed to by s.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalSSLGetCert(b []byte, fw int, s *SSLGetCertMeasurement) error {
	var err error

	switch {
	case fw >= 4750:
		err = parseSSLGetCert(b, s)
	default:
		err = fmt.Errorf("Unsupported firmware version %d", fw)
	}
	return err
}

func parseSSLGetCert(b []byte, s *SSLGetCertMeasurement) error {
	err := json.Unmarshal(b, &s)
	return err
}

// SSLCert measurement result
type SSLGetCertMeasurement struct {
	AddrFamily      int              `json:"af"`
	Bundle          int              `json:"bundle,omitempty"`
	Alert           *SSLGetCertAlert `json:"alert,omitempty"`
	Cert            []string         `json:"cert,omitempty"`
	DestAddr        string           `json:"dst_addr"`
	DestName        string           `json:"dst_name"`
	DestPort        string           `json:"dst_port"`
	FromAddr        string           `json:"from"`
	Fw              int              `json:"fw"`
	GroupID         int              `json:"group_id,omitempty"`
	Lts             int              `json:"lts"`
	Mode            string           `json:"mode,omitempty"`
	Method          string           `json:"method"`
	MeasurementID   int              `json:"msm_id"`
	MeasurementName string           `json:"msm_name"`
	ProbeID         int              `json:"prb_id"`
	ResponseTime    float64          `json:"rt,omitempty"`
	ServerCipher    string           `json:"server_cipher,omitempty"`
	SrcAddr         string           `json:"src_addr"`
	StoredTimestamp int              `json:"stored_timestamp,omitempty"`
	Timestamp       int              `json:"timestamp"`
	TTC             float64          `json:"ttc,omitempty"`
	TTR             float64          `json:"ttr,omitempty"`
	Type            string           `json:"type"`
	Version         string           `json:"ver,omitempty"`
}

type SSLGetCertAlert struct {
	Level       int `json:"level,omitempty"`
	Description int `json:"description,omitempty"`
}
