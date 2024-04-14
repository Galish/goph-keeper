package cli

import (
	"errors"

	"github.com/manifoldco/promptui"
)

type SelectOption struct {
	Label string
	Run   func()
}

func (o *SelectOption) String() string {
	return o.Label
}

type selectOptions struct {
	HideSelected bool
}

func (c *Cli) Select(label string, items []*SelectOption) {
	index := c.promptSelect(
		label,
		items,
		nil,
	)

	if index >= 0 && items[index].Run != nil {
		items[index].Run()
	}
}

func (c *Cli) Confirm(label string) bool {
	index := c.promptSelect(
		label,
		[]*SelectOption{
			{
				Label: "Yes",
			},
			{
				Label: "No",
			},
		},
		nil,
	)

	return index == 0
}

func (c *Cli) Retry(err error) bool {
	if err == nil {
		return false
	}

	c.Error(err)
	c.Break()

	return c.Confirm("Want to try again")
}

func (c *Cli) promptSelect(label string, items []*SelectOption, opts *selectOptions) int {
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
		Stdin:    c.r,
		Stdout:   c.w,
	}

	if opts != nil {
		prompt.HideSelected = opts.HideSelected
	}

	index, _, err := prompt.Run()

	if errors.Is(err, promptui.ErrInterrupt) {
		c.Exit()

		return -1
	}

	if err != nil {
		return -1
	}

	return index
}
