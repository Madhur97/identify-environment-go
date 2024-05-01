package detect

import (
	"os/exec"
	"strings"
)

type MacOSVMDetector struct{}

func (det *MacOSVMDetector) Detect() bool {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err == nil {
		if strings.Contains(string(output), "VMware") ||
			strings.Contains(string(output), "Parallels") ||
			strings.Contains(string(output), "VirtualBox") {
			return true
		}
	}
	return false
}

func (det *MacOSVMDetector) Info() string {
	return "Detected virtual machine running on macOS."
}
