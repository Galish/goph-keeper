package ui

import (
	"errors"
	"os"

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

func (ui *UI) Select(label string, items []*SelectOption) {
	index := ui.promptSelect(
		label,
		items,
		nil,
	)

	if index >= 0 && items[index].Run != nil {
		items[index].Run()
	}
}

func (ui *UI) Confirm(label string) bool {
	index := ui.promptSelect(
		label,
		[]*SelectOption{
			{
				Label: "Yes",
			},
			{
				Label: "No",
			},
		},
		&selectOptions{
			HideSelected: true,
		},
	)

	return index == 0
}

func (ui *UI) promptSelect(label string, items []*SelectOption, opts *selectOptions) int {
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
		Stdin:    ui.r,
		Stdout:   ui.w,
	}

	if opts != nil {
		prompt.HideSelected = opts.HideSelected
	}

	index, _, err := prompt.Run()
	if errors.Is(err, promptui.ErrInterrupt) {
		os.Exit(0)
	}

	if err != nil {
		return -1
	}

	return index
}
