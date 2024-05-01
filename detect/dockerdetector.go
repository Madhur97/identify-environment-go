package detect

import (
	"os"
	"strings"
)

// DockerDetector checks if running inside Docker
type DockerDetector struct{}

func (DockerDetector) Detect() bool {
	if data, err := os.ReadFile("/proc/1/cgroup"); err == nil {
		return strings.Contains(string(data), "docker")
	}
	return false
}

func (DockerDetector) Info() string {
	return "Running inside Docker."
}
