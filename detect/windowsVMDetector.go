package detect

import (
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
	return "Detected virtual machine running on Windows."
}
