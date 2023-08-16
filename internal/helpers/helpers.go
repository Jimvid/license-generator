package helpers

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
	"time"
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

// Select licenses
func Select(options []string) (string, error) {
	prompt := promptui.Select{
		Label:        "Choose a license",
		Items:        options,
		Size:         20,
		HideSelected: true,
	}
	_, result, err := prompt.Run()

	if err == nil {
		fmt.Print(TermColors.Green + "✔ " + TermColors.Reset + result + "\n")
	}

	return result, err
}

func GetName() (string, error) {
	prompt := promptui.Prompt{
		Label:       "Enter name",
		HideEntered: true,
	}

	result, err := prompt.Run()

	if err == nil {
		fmt.Printf(TermColors.Green + "✔ " + TermColors.Reset + result + "\n")
	}

	return result, err
}

func GetYear() (string, error) {

	currentTime := time.Now()
	currentYear := currentTime.Year()
	currentYearAsString := strconv.Itoa(currentYear)

	validate := func(input string) error {
		// If nothing is provided return nil
		if len(input) <= 0 {
			return nil
		}

		// If we have some input, make sure its an int
		_, err := strconv.ParseInt(input, 8, 64)
		if err != nil {
			return errors.New("✘ Invalid number")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:       "Enter year",
		Validate:    validate,
		HideEntered: true,
	}

	result, err := prompt.Run()

	if err == nil {
		if len(result) <= 0 {
			fmt.Printf(TermColors.Green + "✔ " + TermColors.Reset + currentYearAsString + "\n")
		} else {
			fmt.Printf(TermColors.Green + "✔ " + TermColors.Reset + result + "\n")
		}
	}

	return currentYearAsString, err
}
