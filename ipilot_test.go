package main

import "testing"

func TestGetCurrentIP(t *testing.T) {
	myip := getLocalIP()
	if myip == "" {
		t.Log("could not retrieve current IP")
		t.Fail()
	}
}

func TestIsIPAddress(t *testing.T) {
	ipaddress := "192.168.0.1"
	isIP := isIPAddress(ipaddress)
	if isIP == false {
		t.Log("could not verify " + ipaddress + " is an IP address")
		t.Fail()
	}
}

func TestResolveDNSHostname(t *testing.T) {
	hostname := "one.one.one.one"
	ipaddress := resolveDNSHostname(hostname)
	if ipaddress != "1.1.1.1" {
		t.Log("could not resolve IP for " + hostname)
		t.Fail()
	}
}
