package ui

import "os"

func (ui *UI) ReadFile(label string, isRequired bool) []byte {
	for {
		filePath := ui.Input(label, isRequired)

		if filePath == "" {
			return nil
		}

		b, err := os.ReadFile(filePath)
		if err != nil {
			ui.Error(err)
			continue
		}

		return b
	}
}

func (ui *UI) WriteFile(label string, data []byte, isRequired bool) {
	for {
		filePath := ui.Input(label, isRequired)

		if filePath == "" {
			return
		}

		if err := os.WriteFile(filePath, data, 0666); err != nil {
			ui.Error(err)
			continue
		}

		break
	}
}
