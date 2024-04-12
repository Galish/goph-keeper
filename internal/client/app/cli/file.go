package cli

import "os"

func (c *Cli) ReadFile(label string, isRequired bool) []byte {
	for {
		filePath := c.Input(label, isRequired)

		if filePath == "" {
			return nil
		}

		b, err := os.ReadFile(filePath)
		if err != nil {
			c.Error(err)

			continue
		}

		return b
	}
}

func (c *Cli) WriteFile(label string, data []byte, isRequired bool) {
	for {
		filePath := c.Input(label, isRequired)

		if filePath == "" {
			return
		}

		if err := os.WriteFile(filePath, data, 0666); err != nil {
			c.Error(err)

			continue
		}

		break
	}
}
