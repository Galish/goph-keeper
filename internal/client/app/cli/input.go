package cli

import (
	"errors"

	"github.com/manifoldco/promptui"
)

var ErrIsRequired = errors.New("field is required")

type inputOptions struct {
	AllowEdit bool
	Default   string
	Mask      rune
	Validate  func(input string) error
}

func (c *Cli) Input(label string, isRequired bool) string {
	return c.promptInput(
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

func (c *Cli) InputPassword(label string, isRequired bool) string {
	return c.promptInput(
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

func (c *Cli) Edit(label, value string, isRequired bool) string {
	return c.promptInput(
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

func (c *Cli) promptInput(label string, opts *inputOptions) string {
	prompt := promptui.Prompt{
		Label:  label,
		Stdin:  c.r,
		Stdout: c.w,
	}

	if opts != nil {
		prompt.AllowEdit = opts.AllowEdit
		prompt.Default = opts.Default
		prompt.Mask = opts.Mask
		prompt.Validate = opts.Validate
	}

	result, err := prompt.Run()
	if errors.Is(err, promptui.ErrInterrupt) {
		c.Exit()

		return ""
	}

	if err != nil {
		return ""
	}

	return result
}
