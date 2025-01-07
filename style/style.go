package style

var (
	DefaultTableStyle  = &TableStyle{Border: BorderDefault}
	DefaultColumnStyle = &ColumnStyle{}
	DefaultRowStyle    = &RowStyle{}
	DefaultCellStyle   = &CellStyle{}
	DefaultHeaderStyle = &HeaderStyle{
		CellStyle: CellStyle{},
		ToUpper:   true,
	}
)

type TableStyle struct {
	Border Border // 边框样式
}

type HeaderStyle struct {
	CellStyle
	ToUpper bool
}

type ColumnStyle struct {
	CellStyle
}

type RowStyle struct {
	CellStyle
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
	if s.Bold {
		text += string(TextBold) + text + string(TextReset)
	}
	if s.Italic {
		text += string(TextItalic) + text + string(TextReset)
	}
	if s.Underline {
		text += string(TextUnderline) + text + string(TextReset)
	}
	if s.FgColor != "" {
		text = string(s.FgColor) + text + string(FgReset)
	}
	if s.BgColor != "" {
		text = string(s.BgColor) + text + string(BgReset)
	}
	return text
}

type TextAlign int

const (
	TextAlignLeft TextAlign = iota
	TextAlignCenter
	TextAlignRight
)
