package ui

import "io"

func (ui *UI) Print(str string) {
	io.WriteString(ui.w, str)
}

func (ui *UI) LineBreak() {
	io.WriteString(ui.w, "\n")
}

func (ui *UI) Error(err error) {
	io.WriteString(ui.e, "An error occured: "+err.Error())
}
