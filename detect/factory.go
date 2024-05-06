package detect

import "errors"

func GetEnvDetector(envType string) (Detector, error) {
	switch envType {
	case "azure":
		return AzureDetector{}, nil
	case "baremetal":
		return BareMetalDetector{}, nil
	default:
		return nil, errors.New("unknown environment type: " + envType)
	}
}
