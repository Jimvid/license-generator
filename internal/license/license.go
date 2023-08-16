package license

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jimvid/license-generator/internal/helpers"
	"net/http"
	"os"
	"strings"
)

type Licenses struct {
	licenses []License
}

type License struct {
	Key, Name, Spdx_id, Url, Node_id string
}

type LicenseContent struct {
	Key         string
	Description string
	Permissions []string
	Conditions  []string
	Limitations []string
	Body        string
}

func New() *Licenses {
	return &Licenses{}
}

// Get licenses
func (l *Licenses) GetLicenses(url string) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	var licenses []License
	err = json.NewDecoder(response.Body).Decode(&licenses)

	if err != nil {
		return err
	}

	l.licenses = licenses

	return nil
}

// Get license names
func (l *Licenses) GetLicenseNames() []string {
	if len(l.licenses) == 0 {
		fmt.Print("No licenses found \n")
	}

	names := make([]string, 0, len(l.licenses))

	for _, license := range l.licenses {
		names = append(names, license.Name)
	}

	return names
}

// Fetch license content
func (l *Licenses) fetchLicenseContent(url string) (LicenseContent, error) {

	response, err := http.Get(url)

	if err != nil {
		return LicenseContent{}, err
	}

	defer response.Body.Close()

	licenseContent := LicenseContent{}
	err = json.NewDecoder(response.Body).Decode(&licenseContent)

	if err != nil {
		return LicenseContent{}, err
	}

	return licenseContent, err
}

// Get license content by name
func (l *Licenses) GetLicenseContent(name string) (LicenseContent, error) {
	for _, license := range l.licenses {
		if license.Name == name {
			return l.fetchLicenseContent(license.Url)
		}
	}

	return LicenseContent{}, errors.New("Not found")
}

// Write license to file
func (l *Licenses) WriteLicense(licenseContent LicenseContent) error {
	name, err := helpers.GetName()
	if err != nil {
		return err
	}

	year, err := helpers.GetYear()
	if err != nil {
		return err
	}

	replacer := strings.NewReplacer(
		"[year]", year,
		"[yyyy]", year,
		"[fullname]", name,
		"[name of copyright owner]", name,
		"<year>", year,
		"<name of author>", name,
	)

	body := replacer.Replace(licenseContent.Body)

	// Create file
	file, err := os.Create("LICENSE")
	if err != nil {
		return err
	}

	defer file.Close()

	// Write to file
	_, err = file.WriteString(body)
	if err != nil {
		return err
	}

	fmt.Println(
		helpers.TermColors.Green +
			"✔ " + helpers.TermColors.Reset + "License created successfully!\n" +
			helpers.TermColors.Reset + "Make sure to check the LICENSE and make changes accordingly if needed.\n",
	)

	return nil
}
