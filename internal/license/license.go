package main

type LicenseGenerator struct {
	licenses []License
}

type License struct {
	key, name, spdx_id, url, node_id string
}

type LicenseContent struct {
	key         string
	description string
	permissions []string
	conditions  []string
	limitations []string
	body        string
}

func New() *LicenseGenerator {
	return &LicenseGenerator{}
}

func (l *LicenseGenerator) fetchLicenses(url string) {

}

func (l *LicenseGenerator) getLicenseNames() []string {
	return []string{""}
}

func (l *LicenseGenerator) getLicenseFromName(name string) LicenseContent {
	return LicenseContent{}
}
