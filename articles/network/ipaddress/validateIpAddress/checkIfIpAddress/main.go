package main

import (
	"fmt"
	"net"
)

func main() {
	validIPV4 := "10.40.210.253"
	checkIPAddressType(validIPV4)

	invalidIPV4 := "1000.40.210.253"
	checkIPAddressType(invalidIPV4)

	valiIPV6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	checkIPAddressType(valiIPV6)

	invalidIPV6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334:3445"
	checkIPAddressType(invalidIPV6)

}

func checkIPAddressType(ip string) {
	if net.ParseIP(ip) == nil {
		fmt.Printf("Invalid IP Address: %s\n", ip)
		return
	}
	for i := 0; i < len(ip); i++ {
		switch ip[i] {
		case '.':
			fmt.Printf("Given IP Address %s is IPV4 type\n", ip)
			return
		case ':':
			fmt.Printf("Given IP Address %s is IPV6 type\n", ip)
			return
		}
	}

}
