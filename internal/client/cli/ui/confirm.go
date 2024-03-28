package ui

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
