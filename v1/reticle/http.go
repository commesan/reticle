package reticle

import (
	"encoding/json"
	"fmt"
)

// UnmarshalHTTP parses the JSON-encoded data and stores the result in the value pointed to by h.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalHTTP(b []byte, fw int, h *HTTPMeasurement) error {
	var err error

	switch {
	case fw >= 4650:
		err = parseHTTP(b, h)
	default:
		err = fmt.Errorf("Unsupported firmware version %d", fw)
	}
	return err
}

func parseHTTP(b []byte, h *HTTPMeasurement) error {
	err := json.Unmarshal(b, &h)
	return err
}

// HTTPMeasurement measurement result
type HTTPMeasurement struct {
	Bundle          int          `json:"bundle,omitempty"`
	FromAddr        string       `json:"from,omitempty"`
	Fw              int          `json:"fw"`
	GroupID         int          `json:"group_id,omitempty"`
	Lts             int          `json:"lts"`
	MeasurementID   int          `json:"msm_id"`
	MeasurementName string       `json:"msm_name"`
	ProbeID         int          `json:"prb_id"`
	Result          []HTTPResult `json:"result"`
	StoredTimestamp int          `json:"stored_timestamp"`
	Timestamp       int          `json:"timestamp"`
	Type            string       `json:"type"`
	URI             string       `json:"uri"`
}

type HTTPResult struct {
	AddrFamily int              `json:"af"`
	BodySize   int              `json:"bsize,omitempty"`
	DNSError   string           `json:"dnserr,omitempty"`
	DestAddr   string           `json:"dst_addr"`
	Error      string           `json:"err,omitempty"`
	Header     []string         `json:"header,omitempty"`
	HeaderSize int              `json:"hsize,omitempty"`
	Method     string           `json:"method"`
	ReadTiming []HTTPReadTiming `json:"readtiming,omitempty"`
	Res        int              `json:"res,omitempty"`
	RT         float64          `json:"rt,omitempty"`
	SrcAddr    string           `json:"src_addr,omitempty"`
	SubID      int              `json:"subid,omitempty"`
	SubMax     int              `json:"submax,omitempty"`
	Time       int              `json:"time,omitempty"`
	TTC        float64          `json:"ttc,omitempty"`
	TTFB       float64          `json:"ttfb,omitempty"`
	TTR        float64          `json:"ttr,omitempty"`
	Version    string           `json:"ver,omitempty"`
}

type HTTPReadTiming struct {
	Offset int     `json:"o"`
	Time   float64 `json:"t"`
}
