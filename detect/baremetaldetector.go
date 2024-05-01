package detect

import (
	"runtime"
)

// BareMetalDetector checks for indications of a non-virtualized environment
type BareMetalDetector struct{}

func (BareMetalDetector) Detect() bool {
	return getBareMetalDetector().Detect()
}

func getBareMetalDetector() Detector {
	switch runtime.GOOS {
	case "linux":
		return &LinuxBareMetalDetector{}
	case "windows":
		return &WindowsBareMetalDetector{}
	case "darwin":
		return &MacOSBareMetalDetector{}

	}
	return nil
}

func (b BareMetalDetector) Info() string {
	return getBareMetalDetector().Info()
}
