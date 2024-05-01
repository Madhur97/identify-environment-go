package detect

import (
	"runtime"
)

// VMDetector checks for a common VM signature
type VMDetector struct{}

func (VMDetector) Detect() bool {
	return getVMDetector().Detect()
}

func getVMDetector() Detector {
	switch runtime.GOOS {
	case "linux":
		return &LinuxVMDetector{}
	case "windows":
		return &WindowsVMDetector{}
	case "darwin":
		return &MacOSVMDetector{}

	}
	return nil
}

func (VMDetector) Info() string {
	return getVMDetector().Info()
}
