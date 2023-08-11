package helpers

import (
	"bytes"
	"github.com/manifoldco/promptui"
	"os/exec"
	"strings"
)

// Get username from global git if possible
func GetGlobalUsername() (string, error) {
	command := exec.Command("git", "config", "--global", "user.name")
	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}

// Select licenses
func Select(options []string) (string, error) {
	prompt := promptui.Select{
		Label:        "Select Day",
		Items:        options,
		Size:         20,
		HideSelected: true,
	}
	_, result, err := prompt.Run()
	return result, err
}

// Create content for the license

// Write file
