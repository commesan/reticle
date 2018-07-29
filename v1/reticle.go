// Reticle
//
// A parsing library for RIPE Atlas measurement results in Go
//
// RIPE Atlas generates a lot of data, and the format of that data changes over time. Often you want to do something simple
// like fetch the median RTT for each measurement result between date X and date Y. Unfortunately, there are dozens of edge
// cases to account for while parsing the JSON, like the format of errors and firmware upgrades that changed the format
// entirely.
//
// Reticle should make it easier for RIPE Atlas users to use the measurement data. For each measurement type
// (ping, traceroute, ...) it has a single struct regardless of the original firmware version. This struct will be modelled
//  after the most recent firmware versions. Older version will be mapped onto the newer versions.
package v1

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Measurement struct {
	Firmware    int         `json:"fw"`
	Type        string      `json:"type,omitempty"`
	Raw         []byte      `json:"raw,omitempty"`
}

func errorIncorrectType(expected, got string) error {
	return errors.New(fmt.Sprintf("incorrect type: measurement is of type %s, expected %s", got, expected))
}

// Parse extracts the firmware version and measurement type and stores the byte array in the Raw field
func Parse(b []byte) (Measurement, error) {
	var m Measurement
	err := json.Unmarshal(b, &m)
	m.Raw = b

	return m, err
}

// TraceRoute returns a TraceRoute struct based on the data in the Raw field. Results in error is the measurement type
// is not traceroute.
func (m *Measurement) TraceRoute() (TraceRoute, error) {
	var tr TraceRoute
	if m.Type != "traceroute" {
		return tr, errorIncorrectType("traceroute", m.Type)
	}
	err := UnmarshalTraceRoute(m.Raw, m.Firmware, &tr)
	return tr, err
}

// Ping returns a Ping struct based on the data in the Raw field. Results in error is the measurement type
// is not ping.
func (m *Measurement) Ping() (Ping, error) {
	var p Ping
	if m.Type != "ping" {
		return p, errorIncorrectType("ping", m.Type)
	}
	err := UnmarshalPing(m.Raw, m.Firmware, &p)
	return p, err
}

// DNS returns a DNS struct based on the data in the Raw field. Results in error is the measurement type
// is not dns.
func (m *Measurement) DNS() (DNS, error) {
	var dns DNS

	if m.Type != "dns" {
		return dns, errorIncorrectType("dns", m.Type)
	}

	err := UnmarshalDNS(m.Raw, m.Firmware, &dns)
	return dns, err
}
