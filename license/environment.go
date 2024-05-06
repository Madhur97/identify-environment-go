package license

import (
	"errors"

	detect "github.com/karl/identify-environment/detect"
)

func validateEnvironmentIdentifier(data map[string]string) error {
	envType, ok := data["envtype"]
	if !ok {
		return errors.New("invalid license: environment type missing")
	}
	detector, err := detect.GetEnvDetector(envType)
	if err != nil {
		return err
	}
	identifiers, err := detector.GetIdentifiers()
	if err != nil {
		return err
	}
	isValid, err := validate(envType, data, identifiers)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("Environment verification for license failed")
	}
	return nil
}

func validate(envType string, data map[string]string, identifiers []string) (bool, error) {
	var licenseValue string
	switch envType {
	case "azure":
		licenseValue = data["azure_subs_key"]
	case "baremetal":
		licenseValue = data["mac_address"]
	default:
		return false, errors.New("unknown environment type: " + envType)
	}
	for _, id := range identifiers {
		if id == licenseValue {
			return true, nil
		}
	}
	return false, nil
}
