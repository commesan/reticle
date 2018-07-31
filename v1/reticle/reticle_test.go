package reticle

import (
	"bytes"
	"testing"
)

const checkMark = "\u2713"
const crossMark = "\u2717"

var parseTest = []struct {
	fw    int
	mType string
	json  string
}{
	{4910, "dns", `{"af":4,"dst_addr":"216.239.34.106","dst_port":"53","from":"195.130.61.208","fw":4910,"group_id":15314087,"lts":243,"msm_id":15314087,"msm_name":"Tdig","prb_id":6366,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":19295,"NSCOUNT":0,"QDCOUNT":1,"abuf":"S1+EAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":34.738,"size":58},"src_addr":"195.130.61.208","stored_timestamp":1532169612,"timestamp":1532169601,"type":"dns"}`},
	{4945, "traceroute", `{"af":6,"dst_addr":"2a01:5a8::1:4","dst_name":"2a01:5a8::1:4","endtime":1532168937,"from":"2001:983:ba7e:1:fad1:11ff:fea9:f090","fw":4945,"group_id":15314075,"lts":24,"msm_id":15314075,"msm_name":"Traceroute","paris_id":1,"prb_id":10001,"proto":"ICMP","result":[{"hop":1,"result":[{"from":"2001:983:ba7e:1:ca0e:14ff:fe74:7476","rtt":1.547,"size":96,"ttl":64},{"from":"2001:983:ba7e:1:ca0e:14ff:fe74:7476","rtt":0.723,"size":96,"ttl":64},{"from":"2001:983:ba7e:1:ca0e:14ff:fe74:7476","rtt":0.71,"size":96,"ttl":64}]},{"hop":2,"result":[{"from":"2001:888:1:1002::1","rtt":5.659,"size":96,"ttl":63},{"from":"2001:888:1:1002::1","rtt":5.204,"size":96,"ttl":63},{"from":"2001:888:1:1002::1","rtt":5.192,"size":96,"ttl":63}]},{"hop":3,"result":[{"from":"2001:888:1:4031::1","rtt":12.446,"size":96,"ttl":62},{"from":"2001:888:1:4031::1","rtt":10.913,"size":96,"ttl":62},{"from":"2001:888:1:4031::1","rtt":5.868,"size":96,"ttl":62}]},{"hop":4,"result":[{"from":"2001:888:1:4005::1","rtt":5.66,"size":96,"ttl":61},{"from":"2001:888:1:4005::1","rtt":5.889,"size":96,"ttl":61},{"from":"2001:888:1:4005::1","rtt":5.677,"size":96,"ttl":61}]},{"hop":5,"result":[{"from":"2001:7f8:1::a500:6939:1","rtt":8.866,"size":96,"ttl":60},{"x":"*"},{"x":"*"}]},{"hop":6,"result":[{"x":"*"},{"x":"*"},{"x":"*"}]},{"hop":7,"result":[{"from":"2a01:5a8::1:4","rtt":44.407,"size":48,"ttl":58},{"from":"2a01:5a8::1:4","rtt":44.414,"size":48,"ttl":58},{"from":"2a01:5a8::1:4","rtt":44.358,"size":48,"ttl":58}]}],"size":48,"src_addr":"2001:983:ba7e:1:fad1:11ff:fea9:f090","stored_timestamp":1532168950,"timestamp":1532168916,"type":"traceroute"}`},
	{4910, "ping", `{"af":4,"avg":-1,"dst_addr":"46.234.34.8","dst_name":"46.234.34.8","dup":0,"from":"220.233.45.200","fw":4910,"group_id":15314073,"lts":82,"max":-1,"min":-1,"msm_id":15314073,"msm_name":"PingMeasurement","prb_id":29005,"proto":"ICMP","rcvd":0,"result":[{"x":"*"},{"x":"*"},{"x":"*"}],"sent":3,"size":64,"src_addr":"192.168.1.22","step":0,"stored_timestamp":1532168496,"timestamp":1532168477,"type":"ping"}`},
	{4940, "http", `{"from":"2a05:7ac0:101:c000:fa1a:67ff:fe4d:7227","fw":4940,"group_id":15313915,"lts":71,"msm_id":15313917,"msm_name":"HTTPGet","prb_id":11296,"result":[{"af":6,"dst_addr":"2a04:9140:3003:1101::a","err":"timeout reading chunk: state 6 linelen 0 lineoffset 0","method":"GET","src_addr":"2a05:7ac0:101:c000:fa1a:67ff:fe4d:7227"}],"stored_timestamp":1532169426,"timestamp":1532169302,"type":"http","uri":"http://fr-ysd-as201958.anchors.atlas.ripe.net:80/4096"}`},
	{4940, "sslcert", `{"af":4,"cert":["-----BEGIN CERTIFICATE-----\nMIIFSzCCBDOgAwIBAgIQAjEKkojC6VJa+iIYFMMHhDANBgkqhkiG9w0BAQsFADCB\nkDELMAkGA1UEBhMCR0IxGzAZBgNVBAgTEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4G\nA1UEBxMHU2FsZm9yZDEaMBgGA1UEChMRQ09NT0RPIENBIExpbWl0ZWQxNjA0BgNV\nBAMTLUNPTU9ETyBSU0EgRG9tYWluIFZhbGlkYXRpb24gU2VjdXJlIFNlcnZlciBD\nQTAeFw0xODAyMDQwMDAwMDBaFw0yMTAyMDMyMzU5NTlaMFExITAfBgNVBAsTGERv\nbWFpbiBDb250cm9sIFZhbGlkYXRlZDEUMBIGA1UECxMLUG9zaXRpdmVTU0wxFjAU\nBgNVBAMTDXJldml3aWtpLmluZm8wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\nAoIBAQDSgqUfWeFmNXXEFYtJNFQOho+MtzPu2WfcbVU/u7JtWLg6Wj8QPzR+oOrP\nbQfd369Zxj9AZdQ0jv22phNwdEpp7WMB8lj+zLr2U35BxqGYFbtPug95jbTvRAaO\ncW0bSadCX+5m5yHCrjrlkJGN2GdbLiZei91SjkJOCkFWGPHllPeHUOFIb1M0pq5l\nIx9uO7WOdggEL9wpzWzIY932JMN+ypCuMXQxraBl0XFwTVgbk0R3e1GZ2b2TrH/v\nrDng5xE/y+KvkYBj/TiaTyBCfCgj+JKBCsr2eOVdabIc1M/x9pCG47y6sAY1SKIW\nlJzJSC/K6yYVkUs99seXHB2s/TDTAgMBAAGjggHdMIIB2TAfBgNVHSMEGDAWgBSQ\nr2o6lFoL2JDqElZz30O0Oija5zAdBgNVHQ4EFgQUxxmTshuOb1W3PgmvdfZMIGIM\nEvowDgYDVR0PAQH/BAQDAgWgMAwGA1UdEwEB/wQCMAAwHQYDVR0lBBYwFAYIKwYB\nBQUHAwEGCCsGAQUFBwMCME8GA1UdIARIMEYwOgYLKwYBBAGyMQECAgcwKzApBggr\nBgEFBQcCARYdaHR0cHM6Ly9zZWN1cmUuY29tb2RvLmNvbS9DUFMwCAYGZ4EMAQIB\nMFQGA1UdHwRNMEswSaBHoEWGQ2h0dHA6Ly9jcmwuY29tb2RvY2EuY29tL0NPTU9E\nT1JTQURvbWFpblZhbGlkYXRpb25TZWN1cmVTZXJ2ZXJDQS5jcmwwgYUGCCsGAQUF\nBwEBBHkwdzBPBggrBgEFBQcwAoZDaHR0cDovL2NydC5jb21vZG9jYS5jb20vQ09N\nT0RPUlNBRG9tYWluVmFsaWRhdGlvblNlY3VyZVNlcnZlckNBLmNydDAkBggrBgEF\nBQcwAYYYaHR0cDovL29jc3AuY29tb2RvY2EuY29tMCsGA1UdEQQkMCKCDXJldml3\naWtpLmluZm+CEXd3dy5yZXZpd2lraS5pbmZvMA0GCSqGSIb3DQEBCwUAA4IBAQBa\nm5h69CCDj/7KvqMZktnWvLRKqBcsR8xa9wGCGT6WVPF+uGZ0m7UGW7YPLmfU/lsO\ncxZLjXI/OoWclN9VZp2zOp6G5f6TywuurYvcX7xH2wNu78Z1t5tiNvnNTFwQPZHs\nLbf4sEIprNtor5QehRpm8A6HZaIa8dL9XIPiUwnYfqiTx3/u71LxdBsfMTQb6cXS\nnstQ1ROdySF4DPkwdV55HMvsViOzak6/cjIrFW+EeMH/1em0UqyZ0gpM+g59ZR3b\nAZW/xQNMC+PWPVtJbzNLX23sHotxnYsZ2dQuFw0Exxyogh2xmq3TOUtpfCzYvJi7\nGJ9wjDVkN9YJ8VphlDaj\n-----END CERTIFICATE-----","-----BEGIN CERTIFICATE-----\nMIIGCDCCA/CgAwIBAgIQKy5u6tl1NmwUim7bo3yMBzANBgkqhkiG9w0BAQwFADCB\nhTELMAkGA1UEBhMCR0IxGzAZBgNVBAgTEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4G\nA1UEBxMHU2FsZm9yZDEaMBgGA1UEChMRQ09NT0RPIENBIExpbWl0ZWQxKzApBgNV\nBAMTIkNPTU9ETyBSU0EgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMTQwMjEy\nMDAwMDAwWhcNMjkwMjExMjM1OTU5WjCBkDELMAkGA1UEBhMCR0IxGzAZBgNVBAgT\nEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4GA1UEBxMHU2FsZm9yZDEaMBgGA1UEChMR\nQ09NT0RPIENBIExpbWl0ZWQxNjA0BgNVBAMTLUNPTU9ETyBSU0EgRG9tYWluIFZh\nbGlkYXRpb24gU2VjdXJlIFNlcnZlciBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEP\nADCCAQoCggEBAI7CAhnhoFmk6zg1jSz9AdDTScBkxwtiBUUWOqigwAwCfx3M28Sh\nbXcDow+G+eMGnD4LgYqbSRutA776S9uMIO3Vzl5ljj4Nr0zCsLdFXlIvNN5IJGS0\nQa4Al/e+Z96e0HqnU4A7fK31llVvl0cKfIWLIpeNs4TgllfQcBhglo/uLQeTnaG6\nytHNe+nEKpooIZFNb5JPJaXyejXdJtxGpdCsWTWM/06RQ1A/WZMebFEh7lgUq/51\nUHg+TLAchhP6a5i84DuUHoVS3AOTJBhuyydRReZw3iVDpA3hSqXttn7IzW3uLh0n\nc13cRTCAquOyQQuvvUSH2rnlG51/ruWFgqUCAwEAAaOCAWUwggFhMB8GA1UdIwQY\nMBaAFLuvfgI9+qbxPISOre44mOzZMjLUMB0GA1UdDgQWBBSQr2o6lFoL2JDqElZz\n30O0Oija5zAOBgNVHQ8BAf8EBAMCAYYwEgYDVR0TAQH/BAgwBgEB/wIBADAdBgNV\nHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwGwYDVR0gBBQwEjAGBgRVHSAAMAgG\nBmeBDAECATBMBgNVHR8ERTBDMEGgP6A9hjtodHRwOi8vY3JsLmNvbW9kb2NhLmNv\nbS9DT01PRE9SU0FDZXJ0aWZpY2F0aW9uQXV0aG9yaXR5LmNybDBxBggrBgEFBQcB\nAQRlMGMwOwYIKwYBBQUHMAKGL2h0dHA6Ly9jcnQuY29tb2RvY2EuY29tL0NPTU9E\nT1JTQUFkZFRydXN0Q0EuY3J0MCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5jb21v\nZG9jYS5jb20wDQYJKoZIhvcNAQEMBQADggIBAE4rdk+SHGI2ibp3wScF9BzWRJ2p\nmj6q1WZmAT7qSeaiNbz69t2Vjpk1mA42GHWx3d1Qcnyu3HeIzg/3kCDKo2cuH1Z/\ne+FE6kKVxF0NAVBGFfKBiVlsit2M8RKhjTpCipj4SzR7JzsItG8kO3KdY3RYPBps\nP0/HEZrIqPW1N+8QRcZs2eBelSaz662jue5/DJpmNXMyYE7l3YphLG5SEXdoltMY\ndVEVABt0iN3hxzgEQyjpFv3ZBdRdRydg1vs4O2xyopT4Qhrf7W8GjEXCBgCq5Ojc\n2bXhc3js9iPc0d1sjhqPpepUfJa3w/5Vjo1JXvxku88+vZbrac2/4EjxYoIQ5QxG\nV/Iz2tDIY+3GH5QFlkoakdH368+PUq4NCNk+qKBR6cGHdNXJ93SrLlP7u3r7l+L4\nHyaPs9Kg4DdbKDsx5Q5XLVq4rXmsXiBmGqW5prU5wfWYQ//u+aen/e7KJD2AFsQX\nj4rBYKEMrltDR5FL1ZoXX/nUh8HCjLfn4g8wGTeGrODcQgPmlKidrv0PJFGUzpII\n0fxQ8ANAe4hZ7Q7drNJ3gjTcBpUC2JD5Leo31Rpg0Gcg19hCC0Wvgmje3WYkN5Ap\nlBlGGSW4gNfL1IYoakRwJiNiqZ+Gb7+6kHDSVneFeO/qJakXzlByjAA6quPbYzSf\n+AZxAeKCINT+b72x\n-----END CERTIFICATE-----","-----BEGIN CERTIFICATE-----\nMIIFdDCCBFygAwIBAgIQJ2buVutJ846r13Ci/ITeIjANBgkqhkiG9w0BAQwFADBv\nMQswCQYDVQQGEwJTRTEUMBIGA1UEChMLQWRkVHJ1c3QgQUIxJjAkBgNVBAsTHUFk\nZFRydXN0IEV4dGVybmFsIFRUUCBOZXR3b3JrMSIwIAYDVQQDExlBZGRUcnVzdCBF\neHRlcm5hbCBDQSBSb290MB4XDTAwMDUzMDEwNDgzOFoXDTIwMDUzMDEwNDgzOFow\ngYUxCzAJBgNVBAYTAkdCMRswGQYDVQQIExJHcmVhdGVyIE1hbmNoZXN0ZXIxEDAO\nBgNVBAcTB1NhbGZvcmQxGjAYBgNVBAoTEUNPTU9ETyBDQSBMaW1pdGVkMSswKQYD\nVQQDEyJDT01PRE8gUlNBIENlcnRpZmljYXRpb24gQXV0aG9yaXR5MIICIjANBgkq\nhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAkehUktIKVrGsDSTdxc9EZ3SZKzejfSNw\nAHG8U9/E+ioSj0t/EFa9n3Byt2F/yUsPF6c947AEYe7/EZfH9IY+Cvo+XPmT5jR6\n2RRr55yzhaCCenavcZDX7P0N+pxs+t+wgvQUfvm+xKYvT3+Zf7X8Z0NyvQwA1onr\nayzT7Y+YHBSrfuXjbvzYqOSSJNpDa2K4Vf3qwbxstovzDo2a5JtsaZn4eEgwRdWt\n4Q08RWD8MpZRJ7xnw8outmvqRsfHIKCxH2XeSAi6pE6p8oNGN4Tr6MyBSENnTnIq\nm1y9TBsoilwie7SrmNnu4FGDwwlGTm0+mfqVF9p8M1dBPI1R7Qu2XK8sYxrfV8g/\nvOldxJuvRZnio1oktLqpVj3Pb6r/SVi+8Kj/9Lit6Tf7urj0Czr56ENCHonYhMsT\n8dm74YlguIwoVqwUHZwK53Hrzw7dPamWoUi9PPevtQ0iTMARgexWO/bTouJbt7IE\nIlKVgJNp6I5MZfGRAy1wdALqi2cVKWlSArvX31BqVUa/oKMoYX9w0MOiqiwhqkfO\nKJwGRXa/ghgntNWutMtQ5mv0TIZxMOmm3xaG4Nj/QN370EKIf6MzOi5cHkERgWPO\nGHFrK+ymircxXDpqR+DDeVnWIBqv8mqYqnK8V0rSS527EPywTEHl7R09XiidnMy/\ns1Hap0flhFMCAwEAAaOB9DCB8TAfBgNVHSMEGDAWgBStvZh6NLQm9/rEJlTvA73g\nJMtUGjAdBgNVHQ4EFgQUu69+Aj36pvE8hI6t7jiY7NkyMtQwDgYDVR0PAQH/BAQD\nAgGGMA8GA1UdEwEB/wQFMAMBAf8wEQYDVR0gBAowCDAGBgRVHSAAMEQGA1UdHwQ9\nMDswOaA3oDWGM2h0dHA6Ly9jcmwudXNlcnRydXN0LmNvbS9BZGRUcnVzdEV4dGVy\nbmFsQ0FSb290LmNybDA1BggrBgEFBQcBAQQpMCcwJQYIKwYBBQUHMAGGGWh0dHA6\nLy9vY3NwLnVzZXJ0cnVzdC5jb20wDQYJKoZIhvcNAQEMBQADggEBAGS/g/FfmoXQ\nzbihKVcN6Fr30ek+8nYEbvFScLsePP9NDXRqzIGCJdPDoCpdTPW6i6FtxFQJdcfj\nJw5dhHk3QBN39bSsHNA7qxcS1u80GH4r6XnTq1dFDK8o+tDb5VCViLvfhVdpfZLY\nUspzgb8c8+a4bmYRBbMelC1/kZWSWfFMzqORcUx8Rww7Cxn2obFshj5cqsQugsv5\nB5a6SE2Q8pTIqXOi6wZ7I53eovNNVZ96YUWYGGjHXkBrI/V5eu+MtWuLt29G9Hvx\nPUsE2JOAWVrgQSQdso8VYFhH2+9uRv0V9dlfmrPb2LjkQLPNlzmuhbsdjrzch5vR\npu/xO28QOG8=\n-----END CERTIFICATE-----"],"dst_addr":"172.104.111.8","dst_name":"reviwiki.info","dst_port":"443","from":"43.230.155.140","fw":4940,"group_id":15418551,"lts":12,"method":"TLS","msm_id":15418551,"msm_name":"SSLCert","prb_id":29296,"rt":176.52825,"server_cipher":"0xcca8","src_addr":"43.230.155.140","stored_timestamp":1532751386,"timestamp":1532751370,"ttc":85.514185,"ttr":495.4986,"type":"sslcert","ver":"1.2"}`},
	{4790, "ntp", `{"af":4,"dst_addr":"68.105.28.12","dst_name":"68.105.28.12","from":"174.68.164.233","fw":4790,"group_id":15489173,"lts":89,"msm_id":15489173,"msm_name":"Ntp","prb_id":28,"proto":"UDP","result":[{"x":"*"},{"x":"*"},{"x":"*"}],"src_addr":"192.168.0.41","stored_timestamp":1532944687,"timestamp":1532944638,"type":"ntp"}`},
	{4765, "wifi", `{"bundle":1490610346,"error":"IPv6 RA timeout","from":"2001:610:148:f00d:6666:b3ff:fed1:2e3a","fw":4765,"group_id":7967499,"msm_id":7967499,"msm_name":"WiFi","prb_id":17616,"timestamp":1490610346,"type":"wifi","wpa_supplicant":{"EAP TLS cipher":"AES-128-SHA","EAP state":"SUCCESS","EAP-PEAPv0 Phase2 method":"MSCHAPV2","Supplicant PAE state":"AUTHENTICATED","address":"66:66:b3:d1:2e:3a","bssid":"00:26:3e:8c:ad:02","connect-time":"1","group_cipher":"CCMP","id":"0","ip_address":"192.87.38.111","key_mgmt":"WPA2/IEEE 802.1X/EAP","mode":"station","pairwise_cipher":"CCMP","selectedMethod":"25 (EAP-PEAP)","ssid":"eduroam","suppPortStatus":"Authorized","wpa_state":"COMPLETED"}}`},
}

func TestParse(t *testing.T) {
	t.Log("Parse should read a []byte and create a Measurement")

	for _, test := range parseTest {
		m, _ := ParseString(test.json)
		if m.Firmware != test.fw {
			t.Logf("\tFirmware does not match, expected %d got %d %s", test.fw, m.Firmware, crossMark)
			t.Fail()
		}
		if m.Type != test.mType {
			t.Logf("\tType does not match, expected %s got %s %s", test.mType, m.Type, crossMark)
			t.Fail()
		}
		if !bytes.Equal(m.Raw, []byte(test.json)) {
			t.Log("\tJson does not match, ")
			t.Log("\t\texpected: ", test.json)
			t.Log("/t/tgot:      ", string(m.Raw))
			t.Fail()
		}

		t.Logf("\tParsing %s %s", test.mType, checkMark)
	}
}

func TestMeasurement_TraceRoute(t *testing.T) {
	t.Log("TraceRouteMeasurement should return a TraceRouteMeasurement struct from the Raw field ofthe Measurement or throw error if the measurement is not a TraceRouteMeasurement measurement")

	for _, test := range parseTest {
		m, _ := Parse([]byte(test.json))

		tr, err := m.TraceRoute()
		if m.Type != "traceroute" && err == nil {
			t.Logf("\tNon matching types should result in error, but didn't %s", crossMark)
			t.Fail()
		}

		if m.Type == tr.Type && err == nil {
			t.Logf("\tReturning TraceRouteMeasurement %s", checkMark)
		}
	}
}

func TestMeasurement_Ping(t *testing.T) {
	t.Log("PingMeasurement should return a PingMeasurement struct from the Raw field ofthe Measurement or throw error if the measurement is not a PingMeasurement measurement")

	for _, test := range parseTest {
		m, _ := Parse([]byte(test.json))

		p, err := m.Ping()
		if m.Type != "ping" && err == nil {
			t.Logf("\tNon matching types should result in error, but didn't %s", crossMark)
			t.Fail()
		}

		if m.Type == p.Type && err == nil {
			t.Logf("\tReturning PingMeasurement %s", checkMark)
		}
	}
}

func TestMeasurement_DNS(t *testing.T) {
	t.Log("DNSMeasurement should return a DNSMeasurement struct from the Raw field of the Measurement or throw error if the measurement is not a DNSMeasurement measurement")

	for _, test := range parseTest {
		m, _ := Parse([]byte(test.json))

		d, err := m.DNS()
		if m.Type != "dns" && err == nil {
			t.Logf("\tNon matching types should result in error, but didn't %s", crossMark)
			t.Fail()
		}

		if m.Type == d.Type && err == nil {
			t.Logf("\tReturning DNSMeasurement %s", checkMark)
		}
	}
}

func TestMeasurement_HTTP(t *testing.T) {
	t.Log("HTTPMeasurement should return a HTTPMeasurement struct from the Raw field of the Measurement or throw error if the measurement is not a HTTPMeasurement measurement")

	for _, test := range parseTest {
		m, _ := Parse([]byte(test.json))

		d, err := m.HTTP()
		if m.Type != "http" && err == nil {
			t.Logf("\tNon matching types should result in error, but didn't %s", crossMark)
			t.Fail()
		}

		if m.Type == d.Type && err == nil {
			t.Logf("\tReturning HTTPMeasurement %s", checkMark)
		}
	}
}

func TestMeasurement_SSLGetCert(t *testing.T) {
	t.Log("SSLGetCertMeasurement should return a SSLGetCertMeasurement struct from the Raw field of the Measurement or throw error if the measurement is not a SSLGetCertMeasurement measurement")

	for _, test := range parseTest {
		m, _ := Parse([]byte(test.json))

		d, err := m.SSLGetCert()
		if m.Type != "sslcert" && err == nil {
			t.Logf("\tNon matching types should result in error, but didn't %s", crossMark)
			t.Fail()
		}

		if m.Type == d.Type && err == nil {
			t.Logf("\tReturning SSLGetCertMeasurement %s", checkMark)
		}
	}
}
func TestMeasurement_NTP(t *testing.T) {
	t.Log("NTPMeasurement should return a NTPMeasurement struct from the Raw field of the Measurement or throw error if the measurement is not a NTPMeasurement measurement")

	for _, test := range parseTest {
		m, _ := Parse([]byte(test.json))

		d, err := m.NTP()
		if m.Type != "ntp" && err == nil {
			t.Logf("\tNon matching types should result in error, but didn't %s", crossMark)
			t.Fail()
		}

		if m.Type == d.Type && err == nil {
			t.Logf("\tReturning NTPMeasurement %s", checkMark)
		}
	}
}
func TestMeasurement_Wifi(t *testing.T) {
	t.Log("Wifi should return a Wifi struct from the Raw field of the Measurement or throw error if the measurement is not a Wifi measurement")

	for _, test := range parseTest {
		m, _ := Parse([]byte(test.json))

		d, err := m.Wifi()
		if m.Type != "wifi" && err == nil {
			t.Logf("\tNon matching types should result in error, but didn't %s", crossMark)
			t.Fail()
		}

		if m.Type == d.Type && err == nil {
			t.Logf("\tReturning Wifi %s", checkMark)
		}
	}
}
