package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	printHeader()

	var input string
	var wantPrefixes bool

	flag.StringVar(&input, "i", "", "Specify IP address or domain name.")
	flag.BoolVar(&wantPrefixes, "p", false, "Enable printing of advertised BGP prefixes.")
	flag.Parse()

	if input == "" {
		fmt.Println("FATAL: No IP address or domain name was specified.")
		os.Exit(1)
	} else {
		if input == "me" {
			input = getLocalIP()
		}
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

}
