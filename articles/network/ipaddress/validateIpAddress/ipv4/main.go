package main

import (
	"fmt"
	"net"
)

func main() {
	validIPV4 := "10.40.210.253"
	checkIPAddress(validIPV4)

	invalidIPV4 := "1000.40.210.253"
	checkIPAddress(invalidIPV4)

	valiIPV6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	checkIPAddress(valiIPV6)

	invalidIPV6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334:3445"
	checkIPAddress(invalidIPV6)

}

func checkIPAddress(ip string) {
	if net.ParseIP(ip) == nil {
		fmt.Printf("IP Address: %s - Invalid\n", ip)
	} else {
		fmt.Printf("IP Address: %s - Valid\n", ip)
	}
}
