package text

import "github.com/charmbracelet/x/ansi"

func TruncateText(text string, maxChars int, tails string) string {
	truncatedText := ansi.Truncate(text, maxChars-len(tails), "")
	if text != truncatedText {
		return truncatedText + tails
	}

	return text
}