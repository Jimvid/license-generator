package main

import (
	"fmt"
	"github.com/jimvid/license-generator/internal/helpers"
	"github.com/jimvid/license-generator/internal/license"
)

func main() {
	l := license.New()

	// fetch licenses
	err := l.GetLicenses("https://api.github.com/licenses")
	if err != nil {
		fmt.Println("Error fetching licenses: ", err)
		return
	}

	// select from all licenses
	licenseNames := l.GetLicenseNames()
	selectedLicense, err := helpers.Select(licenseNames)

	if err != nil {
		fmt.Println("Error selecting a license: ", err)
		return
	}

	// get license content
	licenseContent, err := l.GetLicenseContent(selectedLicense)
	if err != nil {
		fmt.Println("Error selecting a license: ", err)
		return
	}

	// write license to file
	err = l.WriteLicense(licenseContent)
	if err != nil {
		fmt.Println("Error writing license to file: ", err)
		return
	}
}
