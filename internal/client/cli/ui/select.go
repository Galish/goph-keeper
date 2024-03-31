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

func (ui *UI) Select(label string, items []*SelectOption) {
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
		// HideSelected: true,
		Stdin:  ui.r,
		Stdout: ui.w,
	}

	index, _, err := prompt.Run()
	if errors.Is(err, promptui.ErrInterrupt) {
		os.Exit(0)
	}

	if err != nil || items[index].Run == nil {
		return
	}

	items[index].Run()
}
