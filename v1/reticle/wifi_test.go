package reticle

import (
	"encoding/json"
	"testing"
)

var wifiFirmwareTests = []struct {
	fw   int
	json string
}{
	{4765, `{"bundle":1490610350,"from":"2001:610:514:116:ea94:f6ff:fee3:745a","fw":4765,"group_id":7967499,"msm_id":7967499,"msm_name":"WiFi","prb_id":22407,"timestamp":1490610351,"type":"wifi","wpa_supplicant":{"EAP TLS cipher":"AES-128-SHA","EAP state":"SUCCESS","EAP-PEAPv0 Phase2 method":"MSCHAPV2","Supplicant PAE state":"AUTHENTICATED","address":"ea:94:f6:e3:74:5a","bssid":"d8:54:a2:52:87:94","connect-time":"5","group_cipher":"CCMP","id":"0","ip_address":"192.87.161.72","key_mgmt":"WPA2/IEEE 802.1X/EAP","mode":"station","pairwise_cipher":"CCMP","selectedMethod":"25 (EAP-PEAP)","ssid":"eduroam","suppPortStatus":"Authorized","wpa_state":"COMPLETED"}}`},
	{4765, `{"bundle":1490610346,"error":"IPv6 RA timeout","from":"2001:610:148:f00d:6666:b3ff:fed1:2e3a","fw":4765,"group_id":7967499,"msm_id":7967499,"msm_name":"WiFi","prb_id":17616,"timestamp":1490610346,"type":"wifi","wpa_supplicant":{"EAP TLS cipher":"AES-128-SHA","EAP state":"SUCCESS","EAP-PEAPv0 Phase2 method":"MSCHAPV2","Supplicant PAE state":"AUTHENTICATED","address":"66:66:b3:d1:2e:3a","bssid":"00:26:3e:8c:ad:02","connect-time":"1","group_cipher":"CCMP","id":"0","ip_address":"192.87.38.111","key_mgmt":"WPA2/IEEE 802.1X/EAP","mode":"station","pairwise_cipher":"CCMP","selectedMethod":"25 (EAP-PEAP)","ssid":"eduroam","suppPortStatus":"Authorized","wpa_state":"COMPLETED"}}`},
}

func TestParseWifi(t *testing.T) {
	t.Log("Testing wifi measurement parser for different firmware versions")

	for _, test := range wifiFirmwareTests {
		var w WifiMeasurement
		UnmarshalWifi([]byte(test.json), test.fw, &w)
		j, _ := json.Marshal(w)

		if test.json != string(j) {
			t.Logf("Error parsing version %d\n", test.fw)
			t.Log("Expected: ", test.json)
			t.Log("Got:      ", string(j))
			t.Fail()
		} else {
			t.Logf("\tVersion %d %s\n", test.fw, checkMark)
		}
	}
}
