package getmac

import (
	"log"
	"net"
	"runtime/debug"
	"strings"

	arrayutils "github.com/AchmadRifai/array-utils"
)

func GetMacAddr() []NetworkInterface {
	defer normalError()
	ifas, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	return arrayutils.Map(arrayutils.Filter(ifas, func(ifa net.Interface, i int) bool {
		return ifa.HardwareAddr.String() != ""
	}), interfaceToNetworkInterface)
}

func interfaceToNetworkInterface(ifa net.Interface, i int) NetworkInterface {
	var inter NetworkInterface
	addrs, err := ifa.Addrs()
	if err != nil {
		panic(err)
	}
	inter.IpAddrs = arrayutils.Map(addrs, netAddrToNetworkAddress)
	addrs, err = ifa.MulticastAddrs()
	if err != nil {
		panic(err)
	}
	inter.MulticastAddrs = arrayutils.Map(addrs, netAddrToNetworkAddress)
	inter.Flags = strings.Split(ifa.Flags.String(), "|")
	inter.Mac = ifa.HardwareAddr.String()
	inter.Name = ifa.Name
	return inter
}

func netAddrToNetworkAddress(addr net.Addr, i int) NetworkAddress {
	return NetworkAddress{Addr: addr.Network(), Network: addr.String()}
}

func normalError() {
	if r := recover(); r != nil {
		log.Println("Catched", r)
		log.Println("Stack", string(debug.Stack()))
	}
}

type NetworkInterface struct {
	Mac            string
	Flags          []string
	Name           string
	IpAddrs        []NetworkAddress
	MulticastAddrs []NetworkAddress
}

type NetworkAddress struct {
	Addr    string
	Network string
}
