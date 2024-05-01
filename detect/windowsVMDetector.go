package detect

import (
	"fmt"
	"net"
	"strings"

	"github.com/yusufpapurcu/wmi"
)

type WindowsVMDetector struct{}

func (det *WindowsVMDetector) Detect() bool {
	var dst []struct {
		Manufacturer string
		Model        string
	}
	query := "SELECT Manufacturer, Model FROM Win32_ComputerSystem"
	if err := wmi.Query(query, &dst); err == nil {
		for _, system := range dst {
			if strings.Contains(system.Manufacturer, "VMware") ||
				(strings.Contains(system.Manufacturer, "Microsoft Corporation") && strings.Contains(system.Model, "Virtual")) ||
				strings.Contains(system.Model, "VirtualBox") {
				return true
			}
		}
	}
	return false
}

func (det *WindowsVMDetector) Info() string {
	fmt.Println(GetAdapterInfos())
	return "Detected virtual machine running on Windows."
}

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
