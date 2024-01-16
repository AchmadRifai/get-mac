package getmac

import (
	"fmt"
	"strings"
	"testing"
)

func TestAllMac(t *testing.T) {
	allMac := GetMacAddr()
	for _, mac := range allMac {
		fmt.Println("Mac :", mac.Mac)
		fmt.Println("Name :", mac.Name)
		fmt.Println("Flags :", strings.Join(mac.Flags, ", "))
		fmt.Println("Ip Addr :")
		for _, addr := range mac.IpAddrs {
			fmt.Println("\t", addr.Addr, ":", addr.Network)
		}
		fmt.Println("Multicast Addr :")
		for _, addr := range mac.MulticastAddrs {
			fmt.Println("\t", addr.Addr, ":", addr.Network)
		}
		fmt.Println()
	}
	if allMac == nil {
		t.Log("Mac not found")
	}
}
