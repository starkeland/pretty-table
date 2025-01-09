package style

var DefaultTableStyle = &TableStyle{
	HideBorder: true,
	Border:     BorderASCII,
	Caption: &CellStyle{
		Align:   TextAlignCenter,
		FgColor: FgWhite,
		BgColor: BgYellow,
		Bold:    true,
	},
	Header: &CellStyle{
		Align:   TextAlignLeft,
		FgColor: FgHiBlue,
		Bold:    true,
	},
	Columns: map[int]*CellStyle{
		1: {
			Align: TextAlignRight,
		},
	},
}

type TableStyle struct {
	HideBorder bool
	Border     Border
	Caption    *CellStyle
	Header     *CellStyle
	Columns    map[int]*CellStyle
	Footer     *CellStyle
}

type CellStyle struct {
	Align     TextAlign
	FgColor   Color
	BgColor   Color
	Bold      bool
	Italic    bool
	Underline bool
}

// MergeCellStyles merges multiple cell styles into a single one.
// The order of the styles is important. The last style in the list will be applied first.
func MergeCellStyles(styles ...*CellStyle) *CellStyle {
	if len(styles) == 0 {
		return nil
	}
	if len(styles) == 1 {
		return styles[0]
	}
	merged := &CellStyle{}
	for _, style := range styles {
		if style == nil {
			continue
		}
		if style.Align != TextAlignLeft {
			merged.Align = style.Align
		}
		if style.FgColor != "" {
			merged.FgColor = style.FgColor
		}
		if style.BgColor != "" {
			merged.BgColor = style.BgColor
		}
		if style.Bold != false {
			merged.Bold = true
		}
		if style.Italic != false {
			merged.Italic = true
		}
		if style.Underline != false {
			merged.Underline = true
		}
	}
	return merged
}

func (s *CellStyle) Apply(text string) string {
	if s.FgColor != "" {
		text = string(s.FgColor) + text + string(FgReset)
	}
	if s.BgColor != "" {
		text = string(s.BgColor) + text + string(BgReset)
	}
	if s.Bold {
		text = string(TextBold) + text + string(TextReset)
	}
	if s.Italic {
		text = string(TextItalic) + text + string(TextReset)
	}
	if s.Underline {
		text = string(TextUnderline) + text + string(TextReset)
	}

	return text
}

type TextAlign int

const (
	TextAlignLeft TextAlign = iota
	TextAlignCenter
	TextAlignRight
)
