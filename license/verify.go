package license

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sort"
	"strings"
)

type LicenseVerifier struct{}

type LicenseInfo struct {
	Feature   string
	Signature string
	Data      map[string]string
}

func NewLicenseInfo(feature string, data map[string]string, signature string) LicenseInfo {
	return LicenseInfo{
		Feature:   feature,
		Data:      data,
		Signature: signature,
	}
}

func (LicenseVerifier) Verify(licenses []LicenseInfo, feature string, publicKey *rsa.PublicKey) error {
	var license *LicenseInfo
	for _, lic := range licenses {
		if lic.Feature == feature {
			license = &lic
			break
		}
	}

	if license == nil {
		return errors.New("license not found for the feature : " + feature)
	}

	licInfo := stringifyLicenseInfo(*license)
	err := verifySignature(publicKey, licInfo, license.Signature)
	if err != nil {
		return err
	}

	err = validateEnvironmentIdentifier(license.Data)
	if err != nil {
		return err
	}

	err = validateLicenseValidity(license.Data)
	if err != nil {
		return err
	}
	return nil
}

func verifySignature(publicKey *rsa.PublicKey, data, signatureB64 string) error {
	// Decode the base64-encoded signature.
	sigBytes, err := base64.StdEncoding.DecodeString(signatureB64)
	if err != nil {
		return err
	}
	hashedData := hashData(data)

	// Verify the signature.
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedData, sigBytes)
	if err != nil {
		return errors.New("signature verification failed")
	}
	return nil
}

func hashData(data string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hasher.Sum(nil)
}

func stringifyLicenseInfo(licInfo LicenseInfo) string {
	var buf strings.Builder
	// Assuming you want the feature name to be uppercase as in the original function
	buf.WriteString(strings.ToUpper(licInfo.Feature))

	var keys []string
	for key := range licInfo.Data {
		if key != "LICENSE_SIGNATURE" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	for _, key := range keys {
		buf.WriteString(strings.TrimSpace(key))
		buf.WriteString(strings.TrimSpace(licInfo.Data[key]))
	}
	return buf.String()
}
