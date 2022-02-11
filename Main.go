package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var input string
	var wantPrefixes bool
	var wantHeader bool

	flag.StringVar(&input, "i", "", "IP address or domain")
	flag.BoolVar(&wantPrefixes, "p", false, "print BGP prefixes")
	flag.BoolVar(&wantHeader, "b", true, "enable/disable header")

	flag.Usage = func() {
		fmt.Printf("Usage of iPilot: \n")
		fmt.Printf("Example: iPilot -b=false -i=me -p \n")
		fmt.Printf("  -b	bool	enable/disable header (default true)\n")
		fmt.Printf("  -h	bool	view help\n")
		fmt.Printf("  -i	string	IP address or domain\n")
		fmt.Printf("  -p	bool	print BGP prefixes (default false)\n")
	}

	flag.Parse()

	if wantHeader {
		printHeader()
	}

	if input == "" {
		fmt.Println("FATAL: No IP address or domain name was specified.")
		os.Exit(1)
	}

	if input == "me" {
		input = getLocalIP()
	}
	if isIPAddress(input) {
		printIPInfo(input, wantPrefixes)
	} else {
		ipaddress := resolveDNSHostname(input)
		if isIPAddress(ipaddress) {
			fmt.Println("Domain Name:	", input)
			printIPInfo(ipaddress, wantPrefixes)
		} else {
			fmt.Println("Invalid query.")
		}
	}

}
