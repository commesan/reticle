package reticle

import (
	"encoding/json"
	"fmt"
)

// UnmarshalNTP parses the JSON-encoded data and stores the result in the value pointed to by n.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalNTP(b []byte, fw int, ntp *NTPMeasurement) error {
	var err error

	switch {
	case fw >= 4750:
		err = parseNTP(b, ntp)
	default:
		err = fmt.Errorf("Unsupported firmware version %d", fw)
	}
	return err
}

func parseNTP(b []byte, ntp *NTPMeasurement) error {
	err := json.Unmarshal(b, &ntp)
	return err
}

// NTPMeasurement measurement result
type NTPMeasurement struct {
	AddrFamily      int         `json:"af"`
	Bundle          int         `json:"bundle,omitempty"`
	DestAddr        string      `json:"dst_addr"`
	DestName        string      `json:"dst_name"`
	FromAddr        string      `json:"from"`
	Fw              int         `json:"fw"`
	GroupID         int         `json:"group_id,omitempty"`
	LeapIndicator   string      `json:"li,omitempty"`
	Lts             int         `json:"lts"`
	Mode            string      `json:"mode,omitempty"`
	MeasurementID   int         `json:"msm_id"`
	MeasurementName string      `json:"msm_name"`
	Poll            float64     `json:"poll,omitempty"`
	ProbeID         int         `json:"prb_id"`
	Precision       float64     `json:"precision,omitempty"`
	Proto           string      `json:"proto"`
	RefID           string      `json:"ref-id,omitempty"`
	RefTimestamp    float64     `json:"ref-ts,omitempty"`
	Result          []NTPResult `json:"result"`
	RootDelay       float64     `json:"root-delay,omitempty"`
	RootDispersion  float64     `json:"root-dispersion,omitempty"`
	SrcAddr         string      `json:"src_addr"`
	StoredTimestamp int         `json:"stored_timestamp,omitempty"`
	Stratum         int         `json:"stratum,omitempty"`
	Timestamp       int         `json:"timestamp"`
	TTR             float64     `json:"ttr,omitempty"`
	Type            string      `json:"type"`
	Version         int         `json:"version,omitempty"`
}

type NTPResult struct {
	Timeout           string  `json:"x,omitempty"`
	FinalTimestamp    float64 `json:"final-ts,omitempty"`
	LeapIndicator     string  `json:"li,omitempty"`
	Offset            float64 `json:"offset,omitempty"`
	OriginTimestamp   float64 `json:"origin-ts,omitempty"`
	Precision         float64 `json:"precision,omitempty"`
	ReceiveTimestamp  float64 `json:"receive-ts,omitempty"`
	RefID             string  `json:"ref-id,omitempty"`
	RefTimestamp      float64 `json:"ref-ts,omitempty"`
	RootDelay         float64 `json:"root-delay,omitempty"`
	RootDispersion    float64 `json:"root-dispersion,omitempty"`
	RTT               float64 `json:"rtt,omitempty"`
	Stratum           string  `json:"stratum,omitempty"`
	TransmitTimestamp float64 `json:"transmit-ts,omitempty"`
}
