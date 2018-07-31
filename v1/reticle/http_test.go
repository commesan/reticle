package reticle

import (
	"encoding/json"
	"testing"
)

var httpFirmwareTests = []struct {
	fw   int
	json string
}{
	{4650, `{"from":"2a00:8740:0:2071:ea94:f6ff:fee3:57d8","fw":4650,"group_id":15313915,"lts":97,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":23136,"result":[{"af":6,"bsize":12624,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":144.45,"src_addr":"2a00:8740:0:2071:ea94:f6ff:fee3:57d8","ver":"1.1"}],"stored_timestamp":1532169140,"timestamp":1532169106,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4680, `{"from":"2406:3000:11:1022:eade:27ff:fec9:68d8","fw":4680,"group_id":15313915,"lts":2,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":25050,"result":[{"af":6,"bsize":16836,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":383.385,"src_addr":"2406:3000:11:1022:eade:27ff:fec9:68d8","ver":"1.1"}],"stored_timestamp":1532169157,"timestamp":1532169026,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4770, `{"from":"2a01:e34:ed4e:d840:1ad6:c7ff:fef3:1346","fw":4770,"group_id":15313915,"lts":16,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":33386,"result":[{"af":6,"bsize":16840,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":46.179505,"src_addr":"2a01:e34:ed4e:d840:1ad6:c7ff:fef3:1346","ver":"1.1"}],"stored_timestamp":1532169285,"timestamp":1532169198,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4780, `{"from":"2001:470:6d:e64:c225:e9ff:fe9a:1aaa","fw":4780,"group_id":15313915,"lts":53,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":35468,"result":[{"af":6,"bsize":16828,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":95.795315,"src_addr":"2001:470:6d:e64:c225:e9ff:fe9a:1aaa","ver":"1.1"}],"stored_timestamp":1532169397,"timestamp":1532169324,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4790, `{"from":"2001:660:4701:f011:220:4aff:fec6:cb0c","fw":4790,"group_id":15313915,"lts":17,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":758,"result":[{"af":6,"bsize":12627,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":45.461002,"src_addr":"2001:660:4701:f011:220:4aff:fec6:cb0c","ver":"1.1"}],"stored_timestamp":1532169515,"timestamp":1532169318,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4900, `{"from":"2001:1af8:4020:a035:1::1","fw":4900,"group_id":15313915,"lts":30,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":25686,"result":[{"af":6,"bsize":12588,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":22.891695,"src_addr":"2001:1af8:4020:a035:1::1","ver":"1.1"}],"stored_timestamp":1532169956,"timestamp":1532169880,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4910, `{"from":"2a01:9f40:a000:0:a62b:b0ff:fee0:8b0","fw":4910,"group_id":15313915,"lts":13,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":31702,"result":[{"af":6,"bsize":12621,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":51.959065,"src_addr":"2a01:9f40:a000:0:a62b:b0ff:fee0:8b0","ver":"1.1"}],"stored_timestamp":1532169924,"timestamp":1532169871,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4930, `{"from":"2001:910:1167:42:1:c2ff:fef6:cc24","fw":4930,"group_id":15313915,"lts":2,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":50019,"result":[{"af":6,"bsize":12615,"dst_addr":"2a04:9140:3003:1101::a","hsize":131,"method":"GET","res":200,"rt":71.326833,"src_addr":"2001:910:1167:42:1:c2ff:fef6:cc24","ver":"1.1"}],"stored_timestamp":1532169257,"timestamp":1532169216,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4940, `{"from":"2a05:7ac0:101:c000:fa1a:67ff:fe4d:7227","fw":4940,"group_id":15313915,"lts":71,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":11296,"result":[{"af":6,"dst_addr":"2a04:9140:3003:1101::a","err":"timeout reading chunk: state 6 linelen 0 lineoffset 0","method":"GET","src_addr":"2a05:7ac0:101:c000:fa1a:67ff:fe4d:7227"}],"stored_timestamp":1532169426,"timestamp":1532169302,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
}

func TestUnmarshalHTTP(t *testing.T) {
	t.Log("Testing http parser for different firmware versions")

	for _, test := range httpFirmwareTests {
		var h HTTPMeasurement
		UnmarshalHTTP([]byte(test.json), test.fw, &h)
		hJson, _ := json.Marshal(h)

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
