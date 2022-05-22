package util

import (
	"fmt"
	"log"
	"net"
	"os"
)

// get the local ip and port based on our destination ip
func LocalIPPort(dstip net.IP) (net.IP, int) {
	serverAddr, err := net.ResolveUDPAddr("udp", dstip.String()+":12345")
	if err != nil {
		log.Fatal(err)
	}

	// We don't actually connect to anything, but we can determine
	// based on our destination ip what source ip we should use.
	if con, err := net.DialUDP("udp", nil, serverAddr); err == nil {
		defer con.Close()
		if udpaddr, ok := con.LocalAddr().(*net.UDPAddr); ok {
			return udpaddr.IP, udpaddr.Port
		}
	}
	return nil, -1
}

func DomainLookUp(host string) net.IP {
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("Domain " + host + " Lookup Fail.")
		os.Exit(1)
	}

	var ipSlice = []net.IP{}
	var ipv6Flag = false

	for _, ip := range ips {
		ipSlice = append(ipSlice, ip)
		// 仅返回ipv4的ip
		// if ip.To4() != nil {
		// 	ipSlice = append(ipSlice, ip)
		// } else {
		// 	ipv6Flag = true
		// }
	}

	if ipv6Flag {
		fmt.Println("[Info] IPv6 Traceroute is not supported right now.")
		if len(ipSlice) == 0 {
			os.Exit(0)
		}
	}

	if len(ipSlice) == 1 {
		return ipSlice[0]
	} else {
		fmt.Println("Please Choose the IP You Want To TraceRoute")
		for i, ip := range ipSlice {
			fmt.Printf("%d. %s\n", i, ip)
		}
		var index int
		fmt.Printf("Your Option: ")
		fmt.Scanln(&index)
		if index >= len(ipSlice) || index < 0 {
			fmt.Println("Your Option is invalid")
			os.Exit(3)
		}
		return ipSlice[index]
	}
}
