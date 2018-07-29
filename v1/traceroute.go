package v1

import (
	"encoding/json"
	"errors"
	"fmt"
)

// UnmarshallTraceRoute parses the JSON-encoded data and stores the result in the value pointed to by tr.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalTraceRoute(b []byte, fw int, tr *TraceRoute) (error) {
	var err error

	switch {
	case fw >= 4650:
		err = parseTraceRoute(b, tr)
	default:
		err = errors.New(fmt.Sprintf("Unsupported firmware version %d", fw))
	}
	return err
}

// Works for vfw versions:
// Version 4750 or greater.
// Version 4610 is identified by a value between "4610" and "4749" for the key fw in the result.
func parseTraceRoute(b []byte, tr *TraceRoute) (error) {
	err := json.Unmarshal(b, &tr)
	return err
}

// TraceRoute a struct reprensentation for a RIPE traceroute measurement
type TraceRoute struct {
	AddrFamily      int                `json:"af"`
	Bundle          int                `json:"bundle,omitempty"`
	DestAddr        string             `json:"dst_addr"`
	DestName        string             `json:"dst_name"`
	Endtime         int                `json:"endtime"`
	FromAddr        string             `json:"from"`
	Fw              int                `json:"fw"`
	GroupId         int                `json:"group_id"`
	Lts             int                `json:"lts"`
	MeasurementId   int                `json:"msm_id"`
	MeasurementName string             `json:"msm_name"`
	ParisId         int                `json:"paris_id"`
	SrcProbeId      int                `json:"prb_id"`
	Proto           string             `json:"proto"`
	Result          []TraceRouteResult `json:"result"`
	Size            int                `json:"size"`
	SrcAddr         string             `json:"src_addr"`
	StoredTimestamp int 			   `json:"stored_timestamp"`
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
	TTLInPacket     int      `json:"ittl,omitempty"`
	DestAddrErr     string   `json:"edst,omitempty"`
	Late            int      `json:"late,omitempty"`
	Mtu             int      `json:"mtu,omitempty"`
	RTT             float64  `json:"rtt,omitempty"`
	Size            int      `json:"size,omitempty"`
	TTL             int      `json:"ttl,omitempty"`
	Flags           string   `json:"flags,omitempty"`
	DestOptSize     int      `json:"dstoptsize,omitempty"`
	HopByHopOptSize int      `json:"hbhoptsize,omitempty"`
	ICMPExt         *ICMPExt `json:"icmpext,omitempty"`
}

type ICMPExt struct {
	Version int         `json:"version"`
	RFC4884 int         `json:"frc4884"`
	Object  *IcmpObject `json:"obj"`
}

type IcmpObject struct {
	Class int    `json:"class"`
	Type  int    `json:"type"`
	MPLS  []MPLS `json:"mpls,omitempty"`
}

type MPLS struct {
	Exp   int `json:"exp"`
	Label int `json:"label"`
	S     int `json:"s"`
	TTL   int `json:"ttl"`
}
