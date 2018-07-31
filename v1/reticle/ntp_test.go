package reticle

import (
	"encoding/json"
	"testing"
)

var ntpFirmwareTests = []struct {
	fw   int
	json string
}{
	{4790, `{"af":4,"dst_addr":"68.105.28.12","dst_name":"68.105.28.12","from":"174.68.164.233","fw":4790,"group_id":15489173,"lts":89,"msm_id":15489173,"msm_name":"Ntp","prb_id":28,"proto":"UDP","result":[{"x":"*"},{"x":"*"},{"x":"*"}],"src_addr":"192.168.0.41","stored_timestamp":1532944687,"timestamp":1532944638,"type":"ntp"}`},
	{4940, `{"af":6,"dst_addr":"2001:470:702b:23::1","dst_name":"2001:470:702b:23::1","from":"2001:470:702b:23:ee08:6bff:fe73:4392","fw":4940,"group_id":15489178,"li":"no","lts":7,"mode":"server","msm_id":15489178,"msm_name":"Ntp","poll":8,"prb_id":29640,"precision":9.537e-7,"proto":"UDP","ref-id":"c11e78f5","ref-ts":3741933342.4067,"result":[{"final-ts":3741933440.9683,"offset":0.041477,"origin-ts":3741933440.88476,"receive-ts":3741933440.885048,"rtt":0.083514,"transmit-ts":3741933440.885075},{"final-ts":3741933440.96927,"li":"unknown","offset":0.000263,"origin-ts":3741933440.96874,"precision":1,"receive-ts":3741933440.9687,"ref-id":"RATE","ref-ts":0.0001,"root-delay":0.0001,"root-dispersion":0.0001,"rtt":0.000527,"stratum":"invalid","transmit-ts":3741933440.96874},{"x":"*"}],"root-delay":0.047564,"root-dispersion":0.05335,"src_addr":"2001:470:702b:23:ee08:6bff:fe73:4392","stored_timestamp":1532944665,"stratum":3,"timestamp":1532944640,"type":"ntp","version":4}`},
}

func TestUnmarshalNTP(t *testing.T) {
	t.Log("Testing ntp parser for different firmware versions")

	for _, test := range ntpFirmwareTests {
		var ntp NTPMeasurement
		UnmarshalNTP([]byte(test.json), test.fw, &ntp)
		hJson, _ := json.Marshal(ntp)

		if test.json != string(hJson) {
			t.Logf("Error parsing version %d\n", test.fw)
			t.Log("Expected: ", test.json)
			t.Log("Got:      ", string(hJson))
			t.Fail()
		} else {
			t.Logf("\tVersion %d %s\n", test.fw, checkMark)
		}
	}
}
