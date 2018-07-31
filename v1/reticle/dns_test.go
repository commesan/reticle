package reticle

import (
	"encoding/json"
	"testing"
)

var dnsFirmwareTests = []struct {
	fw   int
	json string
}{
	{4740, `{"af":4,"dst_addr":"216.239.34.106","from":"208.70.31.51","fw":4740,"group_id":15314087,"lts":8,"msm_id":15314087,"msm_name":"Tdig","prb_id":30251,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":43149,"NSCOUNT":0,"QDCOUNT":1,"abuf":"qI2EAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":150.274,"size":58},"src_addr":"208.70.31.51","stored_timestamp":1532169612,"timestamp":1532169602,"type":"dns"}`},
	{4770, `{"af":4,"dst_addr":"216.239.34.106","from":"208.92.192.4","fw":4770,"group_id":15314087,"lts":29,"msm_id":15314087,"msm_name":"Tdig","prb_id":33151,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":18424,"NSCOUNT":0,"QDCOUNT":1,"abuf":"R/iEAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":97.531,"size":58},"src_addr":"208.92.192.4","stored_timestamp":1532169612,"timestamp":1532169601,"type":"dns"}`},
	{4780, `{"af":4,"dst_addr":"216.239.34.106","from":"198.105.231.114","fw":4780,"group_id":15314087,"lts":42,"msm_id":15314087,"msm_name":"Tdig","prb_id":35296,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":27713,"NSCOUNT":0,"QDCOUNT":1,"abuf":"bEGEAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":105.48,"size":58},"src_addr":"198.105.231.114","stored_timestamp":1532169623,"timestamp":1532169601,"type":"dns"}`},
	{4790, `{"af":4,"dst_addr":"216.239.34.106","from":"213.202.36.42","fw":4790,"group_id":15314087,"lts":68,"msm_id":15314087,"msm_name":"Tdig","prb_id":704,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":15602,"NSCOUNT":0,"QDCOUNT":1,"abuf":"PPKEAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":79.069,"size":58},"src_addr":"213.202.36.42","stored_timestamp":1532169614,"timestamp":1532169600,"type":"dns"}`},
	{4900, `{"af":4,"dst_addr":"216.239.34.106","from":"161.22.123.106","fw":4900,"group_id":15314087,"lts":1,"msm_id":15314087,"msm_name":"Tdig","prb_id":30071,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":8707,"NSCOUNT":0,"QDCOUNT":1,"abuf":"IgOEAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":195.351,"size":58},"src_addr":"161.22.123.106","stored_timestamp":1532169612,"timestamp":1532169601,"type":"dns"}`},
	{4910, `{"af":4,"dst_addr":"216.239.34.106","dst_port":"53","from":"195.130.61.208","fw":4910,"group_id":15314087,"lts":243,"msm_id":15314087,"msm_name":"Tdig","prb_id":6366,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":19295,"NSCOUNT":0,"QDCOUNT":1,"abuf":"S1+EAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":34.738,"size":58},"src_addr":"195.130.61.208","stored_timestamp":1532169612,"timestamp":1532169601,"type":"dns"}`},
}

func TestParseDNS(t *testing.T) {
	t.Log("Testing dns measurement parser for different firmware versions")

	for _, test := range dnsFirmwareTests {
		var dns DNSMeasurement
		UnmarshalDNS([]byte(test.json), test.fw, &dns)
		trJson, _ := json.Marshal(dns)

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
