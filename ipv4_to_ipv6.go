package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ipv4tohex(octets []int) string {
	var res [2]string
	for i := 0; i < 2; i++ {
		hexString1 := strconv.FormatInt(int64(octets[i]), 16)
		hexString2 := strconv.FormatInt(int64(octets[i+1]), 16)
		res[i] = hexString1 + hexString2
	}
	result := res[0] + ":" + res[1]
	return result
}
func StringtoInt(octets []string) []int {
	var intArray []int

	// Convert strings to integers and store in the array
	for _, strNumber := range octets {
		intNumber, err := strconv.Atoi(strNumber)
		if err != nil {
			fmt.Printf("Error converting %s to integer: %v\n", strNumber, err)
			continue
		}
		intArray = append(intArray, intNumber)
	}
	return intArray
}
func splitIP(ipAddress string) ([]string, error) {
	// Parse the IP address
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return nil, fmt.Errorf("Invalid IP address: %s", ipAddress)
	}

	// Check if it's an IPv4 address
	if ip.To4() != nil {
		// For IPv4, return the four octets
		return strings.Split(ip.String(), "."), nil
	}

	// For IPv6, return each hexadecimal part
	return strings.Split(strings.Trim(ip.String(), "[]"), ":"), nil
}

// Convert IPv4 to IPv6-mapped IPv6 address
func convertIPv4ToIPv6(ipv4 net.IP) string {
	ipv6 := "0000:0000:0000:0000:0000:ffff"
	result := ipv6 + ":" + ipv4.String()
	return result
}

func convertIPv4ToIPv6to4(ipv4 net.IP) string {

	octets, err := splitIP(ipv4.String())
	if err != nil {
		fmt.Printf("Error breaking %s into octets \n", ipv4, err)
		return " "
	}
	intArray := StringtoInt(octets)

	result := ipv4tohex(intArray)

	prefix := "2002"
	postfix := "0000:0000:0000:0000:0000"
	ipv6 := prefix + ":" + result + ":" + postfix

	return ipv6
}

func convertIPv4ToIPv6_to_4(ipv4 net.IP) string {

	octets, err := splitIP(ipv4.String())
	if err != nil {
		fmt.Printf("Error breaking %s into octets \n", ipv4, err)
		return " "
	}
	intArray := StringtoInt(octets)

	result := ipv4tohex(intArray)

	prefix := "2002"
	postfix := "0000:0000:0000"
	ipv6 := prefix + ":" + result + ":" + postfix + ":" + result

	return ipv6
}

func main() {
	// Prompt the user to enter an IP address
	fmt.Print("Enter an IP address: ")

	// Read the IP address from the user
	var ipAddressString string

	fmt.Scan(&ipAddressString)

	// Parse the entered string into an IP address
	ipv4Address := net.ParseIP(ipAddressString)

	if ipv4Address == nil {
		fmt.Println("Invalid IP address entered.")
	} else {
		// Display the parsed IP address
		ipv6Address := convertIPv4ToIPv6(ipv4Address)
		ipv6Address6to4 := convertIPv4ToIPv6to4(ipv4Address)
		ipv6Address6_to_4 := convertIPv4ToIPv6_to_4(ipv4Address)
		fmt.Printf("IPv6  address in IPv4 Mapped notation is : %s\n", ipv6Address)
		fmt.Printf("IPv6  address in 6 to 4 notation is : %s\n", ipv6Address6to4)
		fmt.Printf("IPv6  address in 6 to 4 notation is : %s\n", ipv6Address6_to_4)

	}
}
