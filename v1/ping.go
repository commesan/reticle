package v1

import (
	"encoding/json"
	"fmt"
	"errors"
)

// UnmarshallPing parses the JSON-encoded data and stores the result in the value pointed to by p.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalPing(b []byte, fw int, p *Ping) error {
	var err error

	switch {
	case fw >= 4740:
		err = parsePing(b, p)
	default:
		err = errors.New(fmt.Sprintf("Unsupported firmware version %d", fw))
	}
	return err
}

func parsePing(b []byte, p *Ping) error {
	err := json.Unmarshal(b, &p)
	return err
}

// Ping Version 4750 or greater.  See: https://atlas.ripe.net/docs/data_struct/#v4750
type Ping struct {
	AddrFamily      int          `json:"af"`
	AvgRoundTrip    float64      `json:"avg"`
	Bundle          int          `json:"bundle,omitempty"`
	DestAddr        string       `json:"dst_addr"`
	DestName        string       `json:"dst_name"`
	Dup             int          `json:"dup"`
	FromAddr        string       `json:"from"`
	Fw              int          `json:"fw"`
	GroupId         int          `json:"group_id"`
	Lts             int          `json:"lts"`
	MaxRoundTrip    float64      `json:"max"`
	MinRoundTrip    float64      `json:"min"`
	MeasurementId   int          `json:"msm_id"`
	MeasurementName string       `json:"msm_name"`
	SrcProbeId      int          `json:"prb_id"`
	Proto           string       `json:"proto"`
	PkgReceived     int          `json:"rcvd"`
	Result          []PingResult `json:"result"`
	Sent            int          `json:"sent"`
	Size            int          `json:"size"`
	SrcAddr         string       `json:"src_addr"`
	Step            int          `json:"step"`
	StoredTimestamp int          `json:"stored_timestamp,omitempty"`
	Timestamp       int          `json:"timestamp"`
	TTL             int          `json:"ttl,omitempty"`
	TTR             float64      `json:"ttr,omitempty"`
	Type            string       `json:"type"`
}

type PingResult struct {
	Timeout       string  `json:"x,omitempty"`
	Error         string  `json:"error,omitempty"`
	RoundTripTime float64 `json:"rtt,omitempty"`
	SrcAddress    string  `json:"src_Addr,omitempty"`
	TTL           int     `json:"ttl,omitempty"`
	Duplicate     int     `json:"dup,omitempty"`
}
