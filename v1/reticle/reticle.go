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
package reticle

import (
	"encoding/json"
	"fmt"
)

type Measurement struct {
	Firmware int    `json:"fw"`
	Type     string `json:"type,omitempty"`
	Raw      []byte `json:"raw,omitempty"`
}

func errorIncorrectType(expected, got string) error {
	return fmt.Errorf("incorrect type: measurement is of type %s, expected %s", got, expected)
}

// Parse takes json as a []byte and returns a Measurement with the firmware version, measurement type and json byte array
func Parse(b []byte) (Measurement, error) {
	var m Measurement
	err := json.Unmarshal(b, &m)
	m.Raw = b

	return m, err
}

// ParseString takes json as a string and returns a Measurement with the firmware version, measurement type and json byte array
func ParseString(jsonString string) (Measurement, error) {
	jsonBytes := []byte(jsonString)
	m, err := Parse(jsonBytes)
	return m, err
}

// TraceRouteMeasurement returns a TraceRouteMeasurement struct based on the data in the Raw field. Results in error if the measurement type
// is not traceroute.
func (m *Measurement) TraceRoute() (TraceRouteMeasurement, error) {
	var tr TraceRouteMeasurement
	if m.Type != "traceroute" {
		return tr, errorIncorrectType("traceroute", m.Type)
	}
	err := UnmarshalTraceRoute(m.Raw, m.Firmware, &tr)
	return tr, err
}

// PingMeasurement returns a PingMeasurement struct based on the data in the Raw field. Results in error if the measurement type is not ping.
func (m *Measurement) Ping() (PingMeasurement, error) {
	var p PingMeasurement
	if m.Type != "ping" {
		return p, errorIncorrectType("ping", m.Type)
	}
	err := UnmarshalPing(m.Raw, m.Firmware, &p)
	return p, err
}

// DNSMeasurement returns a DNSMeasurement struct based on the data in the Raw field. Results in error if the measurement type is not dns.
func (m *Measurement) DNS() (DNSMeasurement, error) {
	var dns DNSMeasurement

	if m.Type != "dns" {
		return dns, errorIncorrectType("dns", m.Type)
	}

	err := UnmarshalDNS(m.Raw, m.Firmware, &dns)
	return dns, err
}

// HTTPMeasurement returns a HTTPMeasurement struct based on the data in the Raw field. Results in error if the measurement type is not http.
func (m *Measurement) HTTP() (HTTPMeasurement, error) {
	var h HTTPMeasurement

	if m.Type != "http" {
		return h, errorIncorrectType("http", m.Type)
	}

	err := UnmarshalHTTP(m.Raw, m.Firmware, &h)
	return h, err
}

// NTPMeasurement returns a NTPstruct based on the data in the Raw field. Results in error if the measurement type is not ntp.
func (m *Measurement) NTP() (NTPMeasurement, error) {
	var n NTPMeasurement

	if m.Type != "ntp" {
		return n, errorIncorrectType("ntp", m.Type)
	}

	err := UnmarshalNTP(m.Raw, m.Firmware, &n)
	return n, err
}

// SSLGetCertMeasurement returns a SSLGetCertMeasurement struct based on the data in the Raw field. Results in error if the measurement type is not sslcert.
func (m *Measurement) SSLGetCert() (SSLGetCertMeasurement, error) {
	var s SSLGetCertMeasurement

	if m.Type != "sslcert" {
		return s, errorIncorrectType("sslcert", m.Type)
	}

	err := UnmarshalSSLGetCert(m.Raw, m.Firmware, &s)
	return s, err
}

// Wifi returns a Wifi struct based on the data in the Raw field. Results in error if the measurement type is not Wifi.
func (m *Measurement) Wifi() (WifiMeasurement, error) {
	var w WifiMeasurement

	if m.Type != "wifi" {
		return w, errorIncorrectType("wifi", m.Type)
	}

	err := UnmarshalWifi(m.Raw, m.Firmware, &w)
	return w, err
}
