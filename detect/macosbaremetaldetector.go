package detect

import (
	"fmt"
	"os/exec"
	"strings"

	network "github.com/karl/identify-environment/network"
)

type MacOSBareMetalDetector struct{}

func (m *MacOSBareMetalDetector) Detect() bool {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err == nil {
		if strings.Contains(string(output), "VMware") ||
			strings.Contains(string(output), "Parallels") ||
			strings.Contains(string(output), "VirtualBox") {
			return false
		}
		return true
	}
	return false
}

func (m *MacOSBareMetalDetector) Info() string {
	fmt.Println(network.GetAdapterInfos())
	return "Running on bare-metal hardware (macOS)."
}
