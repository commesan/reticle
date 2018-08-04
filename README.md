## Reticle
[![GoDoc](https://godoc.org/github.com/commesan/reticle?status.svg)](https://godoc.org/github.com/commesan/reticle/v1/reticle)

A parsing library for RIPE Atlas measurement results in Go

RIPE Atlas generates a lot of data, and the format of that data changes over time. Often you want to do something simple 
like fetch the median RTT for each measurement result between date X and date Y. Unfortunately, there are dozens of edge 
cases to account for while parsing the JSON, like the format of errors and firmware upgrades that changed the format 
entirely.

Reticle should make it easier for RIPE Atlas users to use the measurement data. For each measurement type 
(ping, traceroute, ...) it has a single struct regardless of the original firmware version. This struct will be modelled
after the most recent firmware versions. Older version will be mapped onto the newer versions. 

### Isn't there a library for this already?
Yes! RIPE's own parsing library [Sagan](https://github.com/RIPE-NCC/ripe.atlas.sagan) is an excellent parser with support 
for all the different probe firmware versions. But it's written in Python and I needed a Go version to integrate with the
rest of my tooling. 

For now Reticle only has support for the more recent firmware versions (>4650). So if you need support for older versions
definitly use Sagan. Support for older firmware versions will be added in time. See below for the availible versions.    

### Usage

#### Install 
```
    go get github.com/commesan/reticle/v1/reticle
```

#### Example
```go
    dnsJson := `{"af":4,"dst_addr":"216.239.34.106","dst_port":"53","from":"195.130.61.208","fw":4910,"group_id":15314087,"lts":243,"msm_id":15314087,"msm_name":"Tdig","prb_id":6366,"proto":"UDP","result":{"ANCOUNT":1,"ARCOUNT":0,"ID":19295,"NSCOUNT":0,"QDCOUNT":1,"abuf":"S1+EAAABAAEAAAAACDBES1A0YzVZBHRlc3QGZ2NwZG5zA25ldAAAAQABwAwAAQABAAAOEAAEaxbr9Q==","rt":34.738,"size":58},"src_addr":"195.130.61.208","stored_timestamp":1532169612,"timestamp":1532169601,"type":"dns"}`

    msmt, _ := ParseString(dnsJson)
    dnsMsmt, _ := msmt.DNS()
    responseTime := dnsMsmt.Result.RT

    fmt.Printf("Request took: %f", responseTime)
```

### Tested for firmware
  - Ping : fw >= 4610
  - Traceroute : All
  - DNS Lookup : fw >= 4610
  - HTTP : fw >= 4750
  - SSL Cert : fw >= 4750
  - NTP : fw >= 4750
  - WIFI : fw >= 4750


### For test data
- SSL: https://atlas.ripe.net/api/v2/measurements/15364693/results/?format=txt
- Ping: https://atlas.ripe.net/api/v2/measurements/9200642/results/?format=txt
- Traceroute: https://atlas.ripe.net/api/v2/measurements/15314075/results/?format=txt
- DNS: https://atlas.ripe.net/api/v2/measurements/15314087/results/?format=txt
- NTP:https://atlas.ripe.net/api/v2/measurements/14872252/results/?format=txt
- HTTP: https://atlas.ripe.net/api/v2/measurements/15313917/results/?start=1532131200&stop=1532217599&format=txt
- WiFi: https://atlas.ripe.net/api/v2/measurements/7948824/results/?format=txt