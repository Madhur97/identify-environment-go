package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	license "github.com/karl/identify-environment/license"
	"gopkg.in/ini.v1"
)

func main() {

	publicKey, err := loadPublicKeyFromFile("./public.pem")
	if err != nil {
		fmt.Print(err)
		return
	}
	licenses, err := readLicenses("./license.lic")
	if err != nil {
		fmt.Print(err)
		return
	}

	var licenseVerifier license.LicenseVerifier
	err = licenseVerifier.Verify(licenses, "F1", publicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

func readLicenses(filePath string) ([]license.LicenseInfo, error) {
	var licenses []license.LicenseInfo
	cfg, err := ini.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("error loading license file: %w", err)
	}

	for _, section := range cfg.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}
		sectionData := make(map[string]string)
		var signature string
		for _, key := range section.Keys() {
			if key.Name() == "sig" {
				signature = key.String()
			} else {
				sectionData[key.Name()] = key.String()
			}
		}
		featureName := section.Name()
		license := license.NewLicenseInfo(featureName, sectionData, signature)
		licenses = append(licenses, license)
	}
	return licenses, nil
}

func loadPublicKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	pemBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("public key PEM decode failed")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("unexpected type of public key")
	}
}
