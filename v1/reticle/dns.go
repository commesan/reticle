package reticle

import (
	"encoding/json"
	"fmt"
)

// UnmarshalDNS parses the JSON-encoded data and stores the result in the value pointed to by d.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalDNS(b []byte, fw int, d *DNSMeasurement) error {
	var err error

	switch {
	case fw >= 4740:
		err = parseDNS(b, d)
	default:
		err = fmt.Errorf("Unsupported firmware version %d", fw)
	}
	return err
}

func parseDNS(b []byte, d *DNSMeasurement) error {
	err := json.Unmarshal(b, &d)
	return err
}

// DNSMeasurement measurement result
type DNSMeasurement struct {
	AddrFamily      int              `json:"af"`
	Bundle          int              `json:"bundle,omitempty"`
	DestAddr        string           `json:"dst_addr,omitempty"`
	DestName        string           `json:"dst_name,omitempty"`
	DestPort        string           `json:"dst_port,omitempty"`
	Error           *DNSError        `json:"error,omitempty"`
	FromAddr        string           `json:"from,omitempty"`
	Fw              int              `json:"fw"`
	GroupId         int              `json:"group_id"`
	Lts             int              `json:"lts"`
	MeasurementId   int              `json:"msm_id"`
	MeasurementName string           `json:"msm_name"`
	ProbeID         int              `json:"prb_id"`
	Proto           string           `json:"proto"`
	QBuf            string           `json:"qbuf,omitempty"`
	Result          *DNSResult       `json:"result"`
	Resultset       []DNSMeasurement `json:"resultset,omitempty"`
	Retry           int              `json:"retry,omitempty"`
	SrcAddr         string           `json:"src_addr,omitempty"`
	SubID           int              `json:"subid,omitempty"`
	SubMax          int              `json:"submax,omitempty"`
	StoredTimestamp int              `json:"stored_timestamp"`
	Timestamp       int              `json:"timestamp"`
	Type            string           `json:"type"`
}

type DNSError struct {
	Timeout     int    `json:"timeout"`
	GetAddrInfo string `json:"getaddrinfo"`
}

type DNSResult struct {
	ANCOUNT int               `json:"ANCOUNT"`
	ARCOUNT int               `json:"ARCOUNT"`
	ID      int               `json:"ID"`
	NSCOUNT int               `json:"NSCOUNT"`
	QDCOUNT int               `json:"QDCOUNT"`
	ABuf    string            `json:"abuf"`
	Answers []DNSResultAnswer `json:"answers,omitempty"`
	RT      float64           `json:"rt,omitempty"`
	Size    int               `json:"size,omitempty"`
}
type DNSResultAnswer struct {
	MNAME  string   `json:"MNAME"`
	NAME   string   `json:"NAME"`
	RDATA  []string `json:"RDATA"`
	RNAME  string   `json:"RNAME"`
	SERIAL int      `json:"SERIAL"`
	TTL    int      `json:"TTL"`
	TYPE   string   `json:"TYPE"`
}

// TODO: Add decoder  for QBuf and ABuf
