package detect

import (
	"fmt"
	"strings"

	network "github.com/karl/identify-environment/network"
	"github.com/yusufpapurcu/wmi"
)

type WindowsBareMetalDetector struct{}

func (w *WindowsBareMetalDetector) Detect() bool {
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
				return false
			}
		}
		return true
	}
	return false
}

func (w *WindowsBareMetalDetector) Info() string {
	fmt.Println(network.GetAdapterInfos())
	return "Running on bare-metal hardware (Windows)."
}
