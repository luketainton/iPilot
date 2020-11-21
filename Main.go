package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getLocalIP() string {
	resp, _ := http.Get("https://api.ipify.org")
	// defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body[:])
}

func main() {
	var ipaddress string
	localIPAddress := getLocalIP()
	flag.StringVar(&ipaddress, "i", localIPAddress, "Specify IP address. Defaults to current public IP address.")
	flag.Parse()
	apiEndpoint := "http://ip-api.com/json/" + ipaddress
	resp, _ := http.Get(apiEndpoint)
	// defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
