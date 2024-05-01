package detect

import (
	"os"
	"strings"
)

type LinuxVMDetector struct{}

func (det *LinuxVMDetector) Detect() bool {
	if data, err := os.ReadFile("/sys/class/dmi/id/product_name"); err == nil {
		vmSignatures := []string{"VMware", "VirtualBox", "KVM", "QEMU"}
		for _, sig := range vmSignatures {
			if strings.Contains(string(data), sig) {
				return true
			}
		}
	}
	return false
}

func (det *LinuxVMDetector) Info() string {
	return "Detected virtual machine running on Linux."
}
