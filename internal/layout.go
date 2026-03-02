package internal

type panelPlacement string

const (
	placementRight  panelPlacement = "right"
	placementBottom panelPlacement = "bottom"
	placementNone   panelPlacement = "none"
)

type viewLayout struct {
	placement  panelPlacement
	listWidth  int
	listHeight int
	infoWidth  int
	infoHeight int
}

func resolveLayout(termWidth, termHeight int) viewLayout {
	if termWidth <= 0 || termHeight <= 0 {
		return viewLayout{placement: placementNone}
	}

	const (
		minListWidth  = 56
		minInfoWidth  = 36
		minListHeight = 10
		minInfoHeight = 11
	)

	infoWidth := maxInt(minInfoWidth, termWidth/3)
	listWidth := termWidth - infoWidth
	if listWidth >= minListWidth {
		return viewLayout{
			placement:  placementRight,
			listWidth:  listWidth,
			listHeight: termHeight,
			infoWidth:  infoWidth,
			infoHeight: termHeight,
		}
	}

	infoHeight := maxInt(minInfoHeight, termHeight/3)
	listHeight := termHeight - infoHeight
	if listHeight >= minListHeight {
		return viewLayout{
			placement:  placementBottom,
			listWidth:  termWidth,
			listHeight: listHeight,
			infoWidth:  termWidth,
			infoHeight: infoHeight,
		}
	}

	return viewLayout{
		placement:  placementNone,
		listWidth:  termWidth,
		listHeight: termHeight,
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
