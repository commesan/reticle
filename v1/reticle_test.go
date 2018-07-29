package v1

import (
	"testing"
	"bytes"
)

const checkMark = "\u2713"
const crossMark = "\u2717"

var parseTest = []struct {
	fw  int
	mType string
	json string
}{
	{4910, "dns",`{"af":4,"dst_addr":"216.239.34.106","dst_port":"53","from":"195.130.61.208","fw":4910,"group_id":15314087,"lts":243,"msm_id":15314087,"msm_name":"Tdig","prb_id":6366,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":19295,"NSCOUNT":0,"QDCOUNT":1,"abuf":"S1+EAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":34.738,"size":58},"src_addr":"195.130.61.208","stored_timestamp":1532169612,"timestamp":1532169601,"type":"dns"}`},
	{4945, "traceroute",`{"af":6,"dst_addr":"2a01:5a8::1:4","dst_name":"2a01:5a8::1:4","endtime":1532168937,"from":"2001:983:ba7e:1:fad1:11ff:fea9:f090","fw":4945,"group_id":15314075,"lts":24,"msm_id":15314075,"msm_name":"Traceroute","paris_id":1,"prb_id":10001,"proto":"ICMP","result":[{"hop":1,"result":[{"from":"2001:983:ba7e:1:ca0e:14ff:fe74:7476","rtt":1.547,"size":96,"ttl":64},{"from":"2001:983:ba7e:1:ca0e:14ff:fe74:7476","rtt":0.723,"size":96,"ttl":64},{"from":"2001:983:ba7e:1:ca0e:14ff:fe74:7476","rtt":0.71,"size":96,"ttl":64}]},{"hop":2,"result":[{"from":"2001:888:1:1002::1","rtt":5.659,"size":96,"ttl":63},{"from":"2001:888:1:1002::1","rtt":5.204,"size":96,"ttl":63},{"from":"2001:888:1:1002::1","rtt":5.192,"size":96,"ttl":63}]},{"hop":3,"result":[{"from":"2001:888:1:4031::1","rtt":12.446,"size":96,"ttl":62},{"from":"2001:888:1:4031::1","rtt":10.913,"size":96,"ttl":62},{"from":"2001:888:1:4031::1","rtt":5.868,"size":96,"ttl":62}]},{"hop":4,"result":[{"from":"2001:888:1:4005::1","rtt":5.66,"size":96,"ttl":61},{"from":"2001:888:1:4005::1","rtt":5.889,"size":96,"ttl":61},{"from":"2001:888:1:4005::1","rtt":5.677,"size":96,"ttl":61}]},{"hop":5,"result":[{"from":"2001:7f8:1::a500:6939:1","rtt":8.866,"size":96,"ttl":60},{"x":"*"},{"x":"*"}]},{"hop":6,"result":[{"x":"*"},{"x":"*"},{"x":"*"}]},{"hop":7,"result":[{"from":"2a01:5a8::1:4","rtt":44.407,"size":48,"ttl":58},{"from":"2a01:5a8::1:4","rtt":44.414,"size":48,"ttl":58},{"from":"2a01:5a8::1:4","rtt":44.358,"size":48,"ttl":58}]}],"size":48,"src_addr":"2001:983:ba7e:1:fad1:11ff:fea9:f090","stored_timestamp":1532168950,"timestamp":1532168916,"type":"traceroute"}`},
	{4910, "ping",`{"af":4,"avg":-1,"dst_addr":"46.234.34.8","dst_name":"46.234.34.8","dup":0,"from":"220.233.45.200","fw":4910,"group_id":15314073,"lts":82,"max":-1,"min":-1,"msm_id":15314073,"msm_name":"Ping","prb_id":29005,"proto":"ICMP","rcvd":0,"result":[{"x":"*"},{"x":"*"},{"x":"*"}],"sent":3,"size":64,"src_addr":"192.168.1.22","step":0,"stored_timestamp":1532168496,"timestamp":1532168477,"type":"ping"}`},
}

func TestParse(t *testing.T) {
	t.Log("Parse should read a []byte and create a Measurement")

	for _, test :=range parseTest {
		m, _ := Parse([]byte(test.json))
		if m.Firmware != test.fw {
			t.Logf("\tFirmware does not match, expected %d got %d %s", test.fw, m.Firmware, crossMark)
			t.Fail()
		}
		if m.Type != test.mType {
			t.Logf("\tType does not match, expected %s got %s %s", test.mType, m.Type, crossMark)
			t.Fail()
		}
		if !bytes.Equal(m.Raw, []byte(test.json))  {
			t.Log("\tJson does not match, ")
			t.Log("\t\texpected: ", test.json)
			t.Log("/t/tgot:      ", string(m.Raw))
			t.Fail()
		}

		t.Logf("\tParsing %s %s", test.mType, checkMark)
	}
}

func TestMeasurement_TraceRoute(t *testing.T) {
	t.Log("TraceRoute should return a TraceRoute struct from the Raw field ofthe Measurement or throw error if the measurement is not a TraceRoute measurement")

	for _, test :=range parseTest {
		m, _ := Parse([]byte(test.json))

		tr, err := m.TraceRoute()
		if m.Type != "traceroute" && err == nil {
			t.Logf("\tGetting a traceroute from type %s should result in error, but didn't %s", m.Type, crossMark)
			t.Fail()
		}

		if m.Type == tr.Type && err == nil {
			t.Logf("\tReturning TraceRoute %s", checkMark)
		}
	}
}

func TestMeasurement_Ping(t *testing.T) {
	t.Log("Ping should return a Ping struct from the Raw field ofthe Measurement or throw error if the measurement is not a Ping measurement")

	for _, test :=range parseTest {
		m, _ := Parse([]byte(test.json))

		p, err := m.Ping()
		if m.Type != "ping" && err == nil {
			t.Logf("\tGetting a traceroute from type %s should result in error, but didn't %s", m.Type, crossMark)
			t.Fail()
		}

		if m.Type == p.Type && err == nil {
			t.Logf("\tReturning Ping %s", checkMark)
		}
	}
}

func TestMeasurement_DNS(t *testing.T) {
	t.Log("DNS should return a DNS struct from the Raw field of the Measurement or throw error if the measurement is not a DNS measurement")

	for _, test :=range parseTest {
		m, _ := Parse([]byte(test.json))

		d, err := m.DNS()
		if m.Type != "dns" && err == nil {
			t.Logf("\tGetting a dns from type %s should result in error, but didn't %s", m.Type, crossMark)
			t.Fail()
		}

		if m.Type == d.Type && err == nil {
			t.Logf("\tReturning DNS %s", checkMark)
		}
	}
}