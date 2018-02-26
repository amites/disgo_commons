package utils

import (
	"strings"
	"net"

)
func GetLocalIP() string {
	return strings.Join(getLocalIPList(), ",")
}

func getLocalIPList() []string {
	var ipList = []string{}

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			var ipAddress = ip.String()

			// var isUnspecified = ip.IsUnspecified()
			// var isLoopback = ip.IsLoopback()
			// var isMulticast = ip.IsMulticast()
			// var isInterfaceLocalMulticast = ip.IsInterfaceLocalMulticast()
			// var isLinkLocalMulticast = ip.IsLinkLocalMulticast()
			// var isLinkLocalUnicast = ip.IsLinkLocalUnicast()
			// var isGlobalUnicast = ip.IsGlobalUnicast()

			if ip.IsGlobalUnicast() {
				ipList = append(ipList, ipAddress)
			}
		}
	}

	return ipList

	// name, err := os.Hostname()
	// if err != nil {
	// 	fmt.Printf("Oops: %v\n", err)
	// 	return ""
	// }

	// addrs, err := net.LookupHost(name)
	// if err != nil {
	// 	fmt.Printf("Oops: %v\n", err)
	// 	return ""
	// }
	// fmt.Printf("Local IP: %s\n", addrs[0])

	// return addrs[0]
}