package detect

import (
	"os"
	"strings"
)

type LinuxBareMetalDetector struct{}

func (l *LinuxBareMetalDetector) Detect() bool {
	if _, err := os.Stat("/sys/class/dmi/id/product_name"); err != nil {
		return false
	}
	data, err := os.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		return false
	}
	vmSignatures := []string{"VMware", "VirtualBox", "KVM", "QEMU", "Xen"}
	for _, sig := range vmSignatures {
		if strings.Contains(string(data), sig) {
			return false
		}
	}
	return true
}

func (l *LinuxBareMetalDetector) Info() string {
	return "Running on bare-metal hardware (Linux)."
}