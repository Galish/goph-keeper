package ui

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
)

var ErrIsRequired = errors.New("field is required")

func (ui *UI) Input(label string, isRequired bool) string {
	validate := func(input string) error {
		if isRequired && input == "" {
			return ErrIsRequired
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()
	if errors.Is(err, promptui.ErrInterrupt) {
		os.Exit(0)
	}

	if err != nil {
		return ""
	}

	return result
}

func (ui *UI) Edit(label, value string, isRequired bool) string {
	validate := func(input string) error {
		if isRequired && input == "" {
			return ErrIsRequired
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:     label,
		Validate:  validate,
		Default:   value,
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if errors.Is(err, promptui.ErrInterrupt) {
		os.Exit(0)
	}

	if err != nil {
		return ""
	}

	return result
}
