package main

import (
	"fmt"
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
