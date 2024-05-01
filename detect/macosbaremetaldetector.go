package detect

import (
	"os/exec"
	"strings"
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
	return "Running on bare-metal hardware (macOS)."
}
