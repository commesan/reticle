package reticle

import (
	"encoding/json"
	"fmt"
)

// UnmarshallTraceRoute parses the JSON-encoded data and stores the result in the value pointed to by tr.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalTraceRoute(b []byte, fw int, tr *TraceRouteMeasurement) error {
	var err error

	switch {
	case fw >= 4460:
		err = parseTraceRoute(b, tr)
	default:
		err = fmt.Errorf("Untested firmware version %d", fw)
	}
	return err
}

func parseTraceRoute(b []byte, tr *TraceRouteMeasurement) error {
	err := json.Unmarshal(b, &tr)
	return err
}

// TraceRouteMeasurement measurement result
type TraceRouteMeasurement struct {
	AddrFamily      int                `json:"af"`
	Bundle          int                `json:"bundle,omitempty"`
	DestAddr        string             `json:"dst_addr"`
	DestName        string             `json:"dst_name"`
	Endtime         int                `json:"endtime"`
	FromAddr        string             `json:"from"`
	Fw              int                `json:"fw"`
	GroupId         int                `json:"group_id,omitempty"`
	Lts             int                `json:"lts,omitempty"`
	MeasurementId   int                `json:"msm_id"`
	MeasurementName string             `json:"msm_name,omitempty"`
	ParisId         int                `json:"paris_id"`
	ProbeID         int                `json:"prb_id"`
	Proto           string             `json:"proto"`
	Result          []TraceRouteResult `json:"result"`
	Size            int                `json:"size"`
	SrcAddr         string             `json:"src_addr"`
	StoredTimestamp int                `json:"stored_timestamp,omitempty"`
	Timestamp       int                `json:"timestamp"`
	TTR             float64            `json:"ttr,omitempty"`
	Type            string             `json:"type"`
}

type TraceRouteResult struct {
	Hop    int                   `json:"hop"`
	Error  string                `json:"error,omitempty"`
	Result []TraceRouteHopResult `json:"result,omitempty"`
}

type TraceRouteHopResult struct {
	Timeout         string   `json:"x,omitempty"`
	Error           string   `json:"err,omitempty"`
	From            string   `json:"from,omitempty"`
	ICMPExt         *ICMPExt `json:"icmpext,omitempty"`
	ITTL            int      `json:"ittl,omitempty"`
	DestAddrErr     string   `json:"dest,omitempty"`
	Late            int      `json:"late,omitempty"`
	Mtu             int      `json:"mtu,omitempty"`
	RTT             float64  `json:"rtt,omitempty"`
	Size            int      `json:"size,omitempty"`
	TTL             int      `json:"ttl,omitempty"`
	Flags           string   `json:"flags,omitempty"`
	DestOptSize     int      `json:"dstoptsize,omitempty"`
	HopByHopOptSize int      `json:"hbhoptsize,omitempty"`

}

type ICMPExt struct {
	Object  []IcmpObject `json:"obj"`
	RFC4884 int         `json:"rfc4884"`
	Version int         `json:"version"`
}

type IcmpObject struct {
	Class int    `json:"class"`
	MPLS  []MPLS `json:"mpls,omitempty"`
	Type  int    `json:"type"`
}

type MPLS struct {
	Exp   int `json:"exp"`
	Label int `json:"label"`
	S     int `json:"s"`
	TTL   int `json:"ttl"`
}
