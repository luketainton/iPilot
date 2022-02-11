package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sort"
	"strings"
)

func getLocalIP() string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		fmt.Println("FATAL: Cannot get local IP.")
		os.Exit(2)
		return ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func isIPAddress(ipaddress string) bool {
	addr := net.ParseIP(ipaddress)
	return addr != nil
}

func resolveDNSHostname(hostname string) string {
	address, _ := net.LookupHost(hostname)
	return address[0]
}

func getIPInfo(ipaddress string) IPAddressInfo {
	apiEndpoint := "http://ip-api.com/json/" + ipaddress
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		fmt.Println("FATAL: Cannot contact IP address information API.")
		os.Exit(3)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	infoString := string(body)
	var info IPAddressInfo
	err = json.Unmarshal([]byte(infoString), &info)
	if err != nil {
		fmt.Println("FATAL: Cannot serialize recieved IP address data.")
		os.Exit(4)
	}
	return info
}

func printBGPPrefixes(as string) {
	apiEndpoint := "https://api.hackertarget.com/aslookup/?q=" + as
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		fmt.Println("FATAL: Cannot contact BGP Prefixes API.")
		os.Exit(5)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	prefixesString := string(body)
	var prefixes = strings.Split(prefixesString, "\n")[1:]
	sort.Strings(prefixes)
	for index := range prefixes {
		fmt.Println(prefixes[index])
	}
}
