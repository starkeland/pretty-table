package style

var (
	DefaultTableStyle = &TableStyle{
		Border: BorderDefault,
		Caption: &CellStyle{
			Align:     TextAlignCenter,
			FgColor:   FgWhite,
			BgColor:   BgBlue,
			Bold:      true,
			Italic:    true,
			Underline: true,
		},
		Header: &CellStyle{
			Align:   TextAlignLeft,
			FgColor: FgBlue,
			BgColor: BgCyan,
			Bold:    true,
		},
	}
	DefaultCellStyle = &CellStyle{}
)

type TableStyle struct {
	Border  *Border
	Caption *CellStyle
	Header  *CellStyle
	Columns map[int]*CellStyle
	Rows    map[int]*CellStyle
	Cells   map[int]map[int]*CellStyle
	Footer  *CellStyle
}

type CellStyle struct {
	Align     TextAlign
	FgColor   Color
	BgColor   Color
	Bold      bool
	Italic    bool
	Underline bool
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
