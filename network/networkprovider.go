package network

import (
	"fmt"
	"net"
)

// AdapterInfo holds information about a network adapter
type AdapterInfo struct {
	Name          string
	IPv4Addresses []string
	MACAddress    string
}

// GetAdapterInfos retrieves information about physical network adapters
func GetAdapterInfos() ([]AdapterInfo, error) {
	var adapterInfos []AdapterInfo

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		// Skip loopback interfaces
		// iface.Flags&net.FlagUp == 0 (this check can be added later to include only active interfaces)
		if iface.Flags&net.FlagLoopback != 0 || iface.HardwareAddr == nil {
			continue
		}

		// skipping locally administered addresses unless necessary
		if iface.HardwareAddr[0]&0x02 == 0x02 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		var ipv4Addresses []string
		for _, addr := range addrs {
			// Parse the address
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ipv4 := ipnet.IP.To4()
			if ipv4 != nil {
				ipv4Addresses = append(ipv4Addresses, ipv4.String())
			}
		}

		adapterInfo := AdapterInfo{
			Name:          iface.Name,
			IPv4Addresses: ipv4Addresses,
			MACAddress:    iface.HardwareAddr.String(),
		}

		adapterInfos = append(adapterInfos, adapterInfo)
	}

	return adapterInfos, nil
}
