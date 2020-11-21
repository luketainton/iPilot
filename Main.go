package main

import (
	"flag"
	"fmt"
)

func main() {
	printHeader()

	var input string
	localIPAddress := getLocalIP()
	flag.StringVar(&input, "i", localIPAddress, "Specify IP address or domain name.")
	flag.Parse()
	var isIPCorrect bool = checkIPSyntax(input)
	if isIPCorrect == true {
		printIPInfo(input)
	} else {
		// fmt.Println(ipaddress, "is not a valid IP address.")
		fmt.Println("Domain Name:	", input)
		ipaddress := resolveDNSHostname(input)
		printIPInfo(ipaddress)
	}

}
