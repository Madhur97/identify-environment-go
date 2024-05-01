package main

import (
	"fmt"

	detect "github.com/karl/identify-environment/detect"
)

func main() {
	detectors := []detect.Detector{
		detect.DockerDetector{},
		detect.VMDetector{},
		detect.BareMetalDetector{},
	}

	for _, detectorss := range detectors {
		if detectorss.Detect() {
			fmt.Println(detectorss.Info())
		}
	}
}
