package style

var DefaultTableStyle = &TableStyle{
	Border: BorderRound,
}

type TableStyle struct {
	Border Border // 边框样式
}

type ColumnStyle struct {
	Color      string
	Background string
	Align      TextAlign
}

type TextAlign int

const (
	TextAlignLeft TextAlign = iota
	TextAlignCenter
	TextAlignRight
)
