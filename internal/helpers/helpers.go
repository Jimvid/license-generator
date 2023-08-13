package helpers

import (
	"bytes"
	"fmt"
	"github.com/manifoldco/promptui"
	"os/exec"
	"strings"
)

// Colors used in the terminal
type termColors struct {
	Green, Yellow, Red, Reset string
}

var TermColors = termColors{
	Green:  "\033[32m",
	Yellow: "\033[33m",
	Red:    "\033[31m",
	Reset:  "\033[0m",
}

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
		Label:        "Select a license: ",
		Items:        options,
		Size:         20,
		HideSelected: true,
	}
	_, result, err := prompt.Run()
	fmt.Print(TermColors.Green + "✔ " + result + "\n")
	return result, err
}
