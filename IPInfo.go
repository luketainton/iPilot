package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// IPAddressInfo is the IP Address information from API
type IPAddressInfo struct {
	Status       string  `json:"status"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"countryCode"`
	Region       string  `json:"region"`
	RegionName   string  `json:"regionName"`
	City         string  `json:"city"`
	ZIP          string  `json:"zip"`
	Latitude     float32 `json:"lat"`
	Longitude    float32 `json:"lon"`
	Timezone     string  `json:"timezone"`
	ISP          string  `json:"isp"`
	Organisation string  `json:"org"`
	AS           string  `json:"as"`
	IPAddress    string  `json:"query"`
}

func getIPInfo(ipaddress string) IPAddressInfo {
	apiEndpoint := "http://ip-api.com/json/" + ipaddress
	resp, _ := http.Get(apiEndpoint)
	body, _ := ioutil.ReadAll(resp.Body)
	infoString := string(body)
	var info IPAddressInfo
	err := json.Unmarshal([]byte(infoString), &info)
	if err != nil {
		fmt.Println(err)
	}
	return info
}

func getBGPPrefixes(as string) {
	apiEndpoint := "https://api.hackertarget.com/aslookup/?q=" + as
	resp, _ := http.Get(apiEndpoint)
	body, _ := ioutil.ReadAll(resp.Body)
	prefixesString := string(body)
	var prefixes = strings.Split(prefixesString, "\n")[1:]
	sort.Strings(prefixes)
	for i := 0; i < len(prefixes); i++ {
		fmt.Println(prefixes[i])
	}
}

func printIPInfo(input string, wantPrefixes bool) {
	var IPInfo IPAddressInfo = getIPInfo(input)
	var location string = IPInfo.Country + "/" + IPInfo.RegionName + "/" + IPInfo.City
	var bgpAS string = strings.Fields(IPInfo.AS)[0]
	fmt.Println("IP Address:	", IPInfo.IPAddress)
	fmt.Println("Location:	", location)
	fmt.Println("Timezone:	", IPInfo.Timezone)
	fmt.Println("ISP:		", IPInfo.ISP)
	fmt.Println("BGP AS:		", bgpAS)
	if wantPrefixes == true {
		fmt.Println("\nBGP Prefixes:")
		getBGPPrefixes(bgpAS)
	}
}
