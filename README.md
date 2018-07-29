## Reticle
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

For now Reticle only has support for the most recent firmware versions (>4650). So if you need support for older versions 
definitly use Sagan. Support for older firmware versions will be added in time. See below for the availible versions.    

### Usage

#### Install 
```
    go get github.com/commesan/reticle
```

### Tested firmware versions
##### Ping
4780,4910
##### Traceroute
4650-4790, 4900-4940
##### DNS Lookup
4740,4770-4790, 4900, 4910
##### HTTP
4650, 4680, 4770-4790, 4900-4940
##### SSL Cert
Not yet implemented
##### NTP
Not yet implemented
##### WiFi
Not yet implemented


### Version overview
- [ ] Version 4750 is currently the most recent version of the datastructure documentation. At the moment any value greather than 4750 conforms to the 4750 documentation. An upper limit to this version will added with the release of a firmware version that changes the datastructures.
  - [x] Ping
  - [x] Traceroute
  - [x] DNS Lookup
  - [X] HTTP
  - [ ] SSL Cert
  - [ ] NTP
  - [ ] WIFI
- [ ] Version 4610 is identified by a value of between "4610" and "4749".
  - [ ] Ping
  - [x] Traceroute
  - [ ] DNS Lookup
  - [X] HTTP
  - [ ] SSL Cert
  - [ ] NTP
- [ ] Version 4570 is identified by a value of between "4570" and "4609".
  - [ ] Ping
  - [] Traceroute
  - [ ] DNS Lookup
  - [ ] HTTP
  - [ ] SSL Cert
- [ ] Version 4540 is identified by a value of between "4540" and "4569".
  - [ ] Ping
  - [ ] Traceroute
  - [ ] DNS Lookup
  - [ ] HTTP
  - [ ] SSL Cert
- [ ] Version 4460 is identified by a value of between "4460" and "4539".
  - [ ] Ping
  - [ ] Traceroute
  - [ ] DNS Lookup
  - [ ] HTTP
  - [ ] SSL Cert
- [ ] Version 4400 is identified by a value of between "4400" and "4459".
  - [ ] Ping
  - [ ] Traceroute
  - [ ] DNS Lookup
  - [ ] HTTP
  - [ ] SSL Cert
- [ ] Version 1 is identified by the value "1".
  - [ ] Ping
  - [ ] Traceroute


### Used test data
- SSL: https://atlas.ripe.net/api/v2/measurements/15285836/results/?format=txt
- Ping: https://atlas.ripe.net/api/v2/measurements/9200642/results/?format=txt
- Traceroute: https://atlas.ripe.net/api/v2/measurements/15314075/results/?format=txt
- DNS: https://atlas.ripe.net/api/v2/measurements/15314087/results/?format=txt
- NTP:https://atlas.ripe.net/api/v2/measurements/14872252/results/?format=txt
- HTTP: https://atlas.ripe.net/api/v2/measurements/15313917/results/?start=1532131200&stop=1532217599&format=txt
- WiFi: https://atlas.ripe.net/api/v2/measurements/7948824/results/?format=txt