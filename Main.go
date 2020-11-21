package main

import (
	"flag"
	"fmt"
)

func main() {
	printHeader()

	var input string
	var wantPrefixes bool
	localIPAddress := getLocalIP()
	flag.StringVar(&input, "i", localIPAddress, "Specify IP address or domain name.")
	flag.BoolVar(&wantPrefixes, "p", false, "Enable printing of advertised BGP prefixes.")
	flag.Parse()
	var isIPCorrect bool = checkIPSyntax(input)
	if isIPCorrect == true {
		printIPInfo(input, wantPrefixes)
	} else {
		ipaddress := resolveDNSHostname(input)
		if checkIPSyntax(ipaddress) == true {
			fmt.Println("Domain Name:	", input)
			printIPInfo(ipaddress, wantPrefixes)
		} else {
			fmt.Println("Invalid query.")
		}
	}

}
