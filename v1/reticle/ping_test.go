package reticle

import (
	"encoding/json"
	"testing"
)

var pingFirmwareTests = []struct {
	fw   int
	json string
}{
	{4780, `{"af":4,"avg":309.163397,"dst_addr":"185.15.245.163","dst_name":"185.15.245.163","dup":0,"from":"203.119.0.195","fw":4780,"group_id":9200642,"lts":818,"max":309.243942,"min":309.027327,"msm_id":9200642,"msm_name":"PingMeasurement","prb_id":6270,"proto":"ICMP","rcvd":3,"result":[{"rtt":309.243942},{"rtt":309.218922},{"rtt":309.027327}],"sent":3,"size":64,"src_addr":"203.119.0.195","step":240,"timestamp":1500411495,"ttl":50,"type":"ping"}`},
	{4910, `{"af":4,"avg":-1,"dst_addr":"46.234.34.8","dst_name":"46.234.34.8","dup":0,"from":"220.233.45.200","fw":4910,"group_id":15314073,"lts":82,"max":-1,"min":-1,"msm_id":15314073,"msm_name":"PingMeasurement","prb_id":29005,"proto":"ICMP","rcvd":0,"result":[{"x":"*"},{"x":"*"},{"x":"*"}],"sent":3,"size":64,"src_addr":"192.168.1.22","step":0,"stored_timestamp":1532168496,"timestamp":1532168477,"type":"ping"}`},
}

func TestParsePing(t *testing.T) {
	t.Log("Testing ping measurement parser for different firmware versions")

	for _, test := range pingFirmwareTests {
		var ping PingMeasurement
		UnmarshalPing([]byte(test.json), test.fw, &ping)
		trJson, _ := json.Marshal(ping)

		if test.json != string(trJson) {
			t.Logf("Error parsing version %d\n", test.fw)
			t.Log("Expected: ", test.json)
			t.Log("Got:      ", string(trJson))
			t.Fail()
		} else {
			t.Logf("\tVersion %d %s\n", test.fw, checkMark)
		}
	}
}
