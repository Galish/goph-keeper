package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

type SelectOption struct {
	Label string
	Run   func()
}

func (o *SelectOption) String() string {
	return o.Label
}

type UI struct {
}

func NewUI() *UI {
	return &UI{}
}

func (ui *UI) Display(str string) {
	fmt.Println(str)
	ui.LineBreak()
}

func (ui *UI) LineBreak() {
	fmt.Println("")
}

func (ui *UI) Select(label string, items []*SelectOption) {
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
		// HideSelected: true,
	}

	index, _, _ := prompt.Run()

	if items[index].Run == nil {
		return
	}

	items[index].Run()
}

func (ui *UI) Input(label string) string {
	var input string

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, label+": ")

		input, _ = r.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "" {
			break
		}
	}

	return input
}

func (ui *UI) Confirm(label string) bool {
	var res bool

	ui.Select(
		label,
		[]*SelectOption{
			{
				Label: "Yes",
				Run: func() {
					res = true
				},
			},
			{
				Label: "No",
			},
		},
	)

	return res
}
