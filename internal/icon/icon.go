package icon

// Style for icons
type Style struct {
	Icon  string
	Color string
}

var (
	Space = " "

	Cursor          = "\uf054"     // Printable Rune : ""
	CheckboxEmpty   = "\U000f0131" // Printable Rune : "󰄱"
	CheckboxChecked = "\U000f0856" // Printable Rune : "󰡖"
	Error           = "\uf530"     // Printable Rune : ""
	Warn            = "\uf071"     // Printable Rune : ""
	Search          = "\ue68f"     // Printable Rune : ""

)
