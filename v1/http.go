package v1

import (
	"encoding/json"
	"errors"
	"fmt"
)

// UnmarshallTraceRoute parses the JSON-encoded data and stores the result in the value pointed to by tr.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalHTTP(b []byte, fw int, h *HTTP) error {
	var err error

	switch {
	case fw >= 4650:
		err = parseHTTP(b, h)
	default:
		err = errors.New(fmt.Sprintf("Unsupported firmware version %d", fw))
	}
	return err
}

// Works for fw versions:
// Version 4750 or greater.
// Version 4610 is identified by a value between "4610" and "4749" for the key fw in the result.
func parseHTTP(b []byte, h *HTTP) error {
	err := json.Unmarshal(b, &h)
	return err
}

// HTTP Version 4790 or greater.  See: https://atlas.ripe.net/docs/data_struct/#v4750
type HTTP struct {
	Bundle          int          `json:"bundle,omitempty"`
	FromAddr        string       `json:"from,omitempty"`
	Fw              int          `json:"fw"`
	GroupID         int          `json:"group_id,omitempty"`
	Lts             int          `json:"lts"`
	MeasurementId   int          `json:"msm_id"`
	MeasurementName string       `json:"msm_name"`
	SrcProbeId      int          `json:"prb_id"`
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
	ResultCode int              `json:"res,omitempty"`
	RT         float64          `json:"rt,omitempty"`
	SrcAddr    string           `json:"src_addr,omitempty"`
	SubID      int              `json:"subid,omitempty"`
	SubMax     int              `json:"submax,omitempty"`
	Time       int              `json:"time,omitempty"`
	TTC        float64          `json:"ttc,omitempty"`
	TTFB       float64          `json:"ttfb,omitempty"`
	TTR        float64          `json:"ttr,omitempty"`
	Version    string           `json:"ver,omitempty"`
} // DNS Version 4750 or greater.  See: https://atlas.ripe.net/docs/data_struct/#v4750

type HTTPReadTiming struct {
	Offset int     `json:"o"`
	Time   float64 `json:"t"`
}

/*

   "bundle" -- [optional] instance ID for a collection of related measurement results (int)
   "from" -- IP address of the probe as known by controller (string)
   "group_id" -- [optional] If the measurement belongs to a group of measurements, the identifier of the group (int)
   "lts" -- last time synchronised. How long ago (in seconds) the clock of the probe was found to be in sync with that of a controller. The value -1 is used to indicate that the probe does not know whether it is in sync (int)
   "msm_id" -- measurement identifier (int)
   "msm_name" -- measurement type "HTTPGet" (string)
   "prb_id" -- source probe ID (int)
   "result" -- results of query (array of objects)
   objects have the following fields:
       "af" -- address family, 4 or 6 (integer)
       "bsize" -- size of body in octets (int)
       "dnserr" -- [optional] DNS resolution failed (string)
       "dst_addr" -- target address (string)
       "err" -- [optional] other failure (string)
       "header" -- [optional] elements are strings. The last string can be empty to indicate the end of enders or end with "[...]" to indicate truncation (array of strings)
       "hsize" -- header size in octets (int)
       "method" -- "GET", "HEAD", or "POST" (string)
       "readtiming" -- [optional] timing results for reply data (array of objects)
       objects have the following fields:
           "o" -- offset in stream of reply data (int, was string before 4790)
           "t" -- time since starting to connect when data is received (in milli seconds) (float)
       "res" -- HTTP result code (int)
       "rt" -- time to execute request excluding DNS (float)
       "src_addr" -- source address used by probe (string)
       "subid" -- [optional] sequence number of this result within a group of results, when the 'all' option is used without the 'combine' option (int)
       "submax" -- [optional] total number of results within a group (int)
       "time" -- [optional] Unix timestamp, when the 'all' option is used with the 'combine' option (int)
       "ttc" -- [optional] time to connect to the target (in milli seconds) (float)
       "ttfb" -- [optional] time to first response byte received by measurent code after starting to connect (in milli seconds) (float)
       "ttr" -- [optional] time to resolve the DNS name (in milli seconds) (float)
       "ver" -- major, minor version of http server (string)
   "timestamp" -- Unix timestamp (int)
   "type" -- "http" (string)
   "uri" -- request uri (string)

*/
