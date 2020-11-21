package main

import (
	"io/ioutil"
	"net"
	"net/http"
)

func getLocalIP() string {
	resp, _ := http.Get("https://api.ipify.org")
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body[:])
}

func checkIPSyntax(ipaddress string) bool {
	addr := net.ParseIP(ipaddress)
	if addr == nil {
		return false
	}
	return true
}

func resolveDNSHostname(hostname string) string {
	address, _ := net.LookupHost(hostname)
	return address[0]
}
