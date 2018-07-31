package reticle

import (
	"encoding/json"
	"fmt"
)

// UnmarshalWifi parses the JSON-encoded data and stores the result in the value pointed to by w.
// It requires you to know the firmware version which this measurement was generated in.
func UnmarshalWifi(b []byte, fw int, w *WifiMeasurement) error {
	var err error

	switch {
	case fw >= 4740:
		err = parseWifi(b, w)
	default:
		err = fmt.Errorf("Unsupported firmware version %d", fw)
	}
	return err
}

func parseWifi(b []byte, w *WifiMeasurement) error {
	err := json.Unmarshal(b, &w)
	return err
}

// Wifi measurment result
type WifiMeasurement struct {
	BSSIDList       []WifiBSSID        `json:"bssid_list,omitempty"`
	Bundle          int                `json:"bundle,omitempty"`
	Error           string             `json:"error,omitempty"`
	FromIP          string             `json:"from"`
	Fw              int                `json:"fw"`
	GroupID         int                `json:"group_id"`
	MeasurementID   int                `json:"msm_id"`
	MeasurementName string             `json:"msm_name"`
	ProbeID         int                `json:"prb_id"`
	StoredTimestamp int                `json:"stored_timestamp,omitempty"`
	Timestamp       int                `json:"timestamp"`
	Type            string             `json:"type"`
	WPASupplicant   *WifiWPASupplicant `json:"wpa_supplicant,omitempty"`
}

type WifiBSSID struct {
	Auth     string  `json:"auth,omitempty"`
	BSSID    string  `json:"bssid,omitempty"`
	Freq     string  `json:"freq,omitempty"`
	SSID     string  `json:"ssid,omitempty"`
	Strength float64 `json:"strength,omitempty"`
}

type WifiWPASupplicant struct {
	EAPTLSCipher         string `json:"EAP TLS cipher,omitempty"`
	EAPState             string `json:"EAP state,omitempty"`
	EAPTTLSv0Phase2      string `json:"EAP-TTLSv0 Phase2 method,omitempty"`
	EAPPEAPHv0Phase2     string `json:"EAP-PEAPv0 Phase2 method"`
	SupplicantPAEState   string `json:"Supplicant PAE state,omitempty"`
	Address              string `json:"address,omitempty"`
	BSSID                string `json:"bssid,omitempty"`
	ConnectTime          string `json:"connect-time,omitempty"`
	GroupCipher          string `json:"group_cipher,omitempty"`
	ID                   string `json:"id,omitempty"`
	IPAddr               string `json:"ip_address,omitempty"`
	KeyManagement        string `json:"key_mgmt,omitempty"`
	Mode                 string `json:"mode,omitempty"`
	PairwiseCipher       string `json:"pairwise_cipher,omitempty"`
	SelectedMethod       string `json:"selectedMethod,omitempty"`
	SSID                 string `json:"ssid,omitempty"`
	SupplicantPortStatus string `json:"suppPortStatus,omitempty"`
	WPAState             string `json:"wpa_state,omitempty"`
}

/*
"bssid_list" -- array of objects with the following fields:

    "auth" -- authentication options (string)
    "bssid" -- MAC address of base station (string)
    "freq" -- channel frequency in MHz (string)
    "ssid" -- SSID (string)
    "strength" -- signal strength (dB)

"bundle" -- [optional] instance ID for a collection of related measurement results (int)
"error" -- reason for failure (string)
"fw" -- probe's firmware version (string)
"from" -- IP address of the probe as known by controller (string)
"group_id" -- [optional] If the measurement belongs to a group of measurements, the identifier of the group (int)
"msm_id" -- measurement identifier (int)
"msm_name" -- "Wifi" (string)
"prb_id" -- probe identifier (int)
"timestamp" -- Unix timestamp (int)
"type" -- "wifi" (string)
"wpa_supplicant" -- object with the following fields:

    "EAP TLS cipher" -- (string)
    "EAP state" -- (string)
    "EAP-TTLSv0 Phase2 method" -- (string)
    "EAP-PEAPv0 Phase2 method" -- (string)
    "Supplicant PAE state" -- (string)
    "address" -- local MAC address (string)
    "bssid" -- MAC address of base station (string)
    "connect-time" -- time in seconds spend trying to connect (integer)
    "group_cipher" -- multicast cipher (string)
    "id" -- (string)
    "key_mgmt" -- key management protocol (string)
    "mode" -- (string)
    "pairwise_cipher" -- point-to-point cipher (string)
    "selectedMethod" -- (string)
    "ssid" -- SSID (string)
    "suppPortStatus" -- (string)
	"wpa_state" -- final state (string)

*/
