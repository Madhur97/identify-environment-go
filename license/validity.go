package license

import (
	"errors"
	"time"
)

func validateLicenseValidity(data map[string]string) error {
	validToDate, ok := data["valid-to-date"]
	if !ok {
		return errors.New("invalid license: validity missing")
	}

	const layout = "2006-01-02"
	expirationDate, err := time.Parse(layout, validToDate)
	if err != nil {
		return errors.New("invalid date format")
	}
	if expirationDate.Before(time.Now()) {
		return errors.New("license has expired")
	}

	return nil
}
