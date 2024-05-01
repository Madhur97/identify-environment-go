package network

import "net"

// AdapterInfo holds information about a network adapter
type AdapterInfo struct {
	Name          string
	IPv4Addresses []string
	MACAddress    string
}

// GetAdapterInfos retrieves information about all network adapters
func GetAdapterInfos() ([]AdapterInfo, error) {
	var adapterInfos []AdapterInfo

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err // Handle error, similar to LOG_WARN in C++
	}

	for _, iface := range interfaces {
		// Skip loopback interfaces similar to the check (ifa->ifa_flags & IFF_LOOPBACK)
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue // Optionally handle the error
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
