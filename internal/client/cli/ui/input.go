package ui

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
)

var ErrIsRequired = errors.New("field is required")

type inputOptions struct {
	AllowEdit bool
	Default   string
	Mask      rune
	Validate  func(input string) error
}

func (ui *UI) Input(label string, isRequired bool) string {
	return ui.promptInput(
		label,
		&inputOptions{
			Validate: func(input string) error {
				if isRequired && input == "" {
					return ErrIsRequired
				}

				return nil
			},
		},
	)
}

func (ui *UI) InputPassword(label string, isRequired bool) string {
	return ui.promptInput(
		label,
		&inputOptions{
			Mask: '*',
			Validate: func(input string) error {
				if isRequired && input == "" {
					return ErrIsRequired
				}

				return nil
			},
		},
	)
}

func (ui *UI) Edit(label, value string, isRequired bool) string {
	return ui.promptInput(
		label,
		&inputOptions{
			AllowEdit: true,
			Default:   value,
			Validate: func(input string) error {
				if isRequired && input == "" {
					return ErrIsRequired
				}

				return nil
			},
		},
	)
}

func (ui *UI) promptInput(label string, opts *inputOptions) string {
	prompt := promptui.Prompt{
		Label:  label,
		Stdin:  ui.r,
		Stdout: ui.w,
	}

	if opts != nil {
		prompt.AllowEdit = opts.AllowEdit
		prompt.Default = opts.Default
		prompt.Mask = opts.Mask
		prompt.Validate = opts.Validate
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
