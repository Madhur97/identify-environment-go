package detect

import (
	"errors"
	"fmt"

	network "github.com/karl/identify-environment/network"
)

// BareMetalDetector checks for indications of a non-virtualized environment
type BareMetalDetector struct{}

func (BareMetalDetector) GetIdentifiers() ([]string, error) {
	adapters, err := network.GetAdapterInfos()
	var macAdresses []string
	if err != nil {
		return macAdresses, err
	}
	for _, adapter := range adapters {
		if adapter.MACAddress == "" {
			fmt.Println("No MAC address found for interface - " + adapter.Name)
			continue
		}
		macAdresses = append(macAdresses, adapter.MACAddress)
	}
	if len(macAdresses) == 0 {
		return macAdresses, errors.New("No adapters found")
	}
	fmt.Println(macAdresses)
	return macAdresses, nil
}
