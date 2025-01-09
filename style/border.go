package style

var BorderASCII = Border{
	TopLeft:         "+",
	Top:             "-",
	TopSeparator:    "+",
	TopRight:        "+",
	HeaderBottom:    "=",
	HrLeft:          "+",
	Hr:              "-",
	HrSeparator:     "+",
	HrRight:         "+",
	ColumnSeparator: "|",
	BottomLeft:      "+",
	Bottom:          "-",
	BottomSeparator: "+",
	BottomRight:     "+",
}

var BorderLight = Border{
	TopLeft:         "┌",
	Top:             "─",
	TopSeparator:    "┬",
	TopRight:        "┐",
	HeaderBottom:    "─",
	HrLeft:          "├",
	Hr:              "─",
	HrSeparator:     "┼",
	HrRight:         "┤",
	ColumnSeparator: "│",
	BottomLeft:      "└",
	Bottom:          "─",
	BottomSeparator: "┴",
	BottomRight:     "┘",
}

var BorderRound = &Border{
	TopLeft:         "╭",
	Top:             "─",
	TopSeparator:    "┬",
	TopRight:        "╮",
	HeaderBottom:    "─",
	HrLeft:          "├",
	Hr:              "─",
	HrSeparator:     "┼",
	HrRight:         "┤",
	ColumnSeparator: "│",
	BottomLeft:      "╰",
	Bottom:          "─",
	BottomSeparator: "┴",
	BottomRight:     "╯",
}

var BorderBold = Border{
	TopLeft:         "┏",
	Top:             "━",
	TopSeparator:    "┳",
	TopRight:        "┓",
	HeaderBottom:    "━",
	HrLeft:          "┣",
	Hr:              "━",
	HrSeparator:     "╋",
	HrRight:         "┫",
	ColumnSeparator: "┃",
	BottomLeft:      "┗",
	Bottom:          "━",
	BottomRight:     "┛",
	BottomSeparator: "┻",
}

var BorderDouble = Border{
	TopLeft:         "╔",
	Top:             "═",
	TopSeparator:    "╦",
	TopRight:        "╗",
	HeaderBottom:    "═",
	HrLeft:          "╠",
	Hr:              "═",
	HrSeparator:     "╬",
	HrRight:         "╣",
	ColumnSeparator: "║",
	BottomLeft:      "╚",
	Bottom:          "═",
	BottomSeparator: "╩",
	BottomRight:     "╝",
}

type Border struct {
	TopLeft      string
	Top          string
	TopSeparator string
	TopRight     string

	HeaderBottom string

	HrLeft      string
	Hr          string
	HrSeparator string
	HrRight     string

	ColumnSeparator string

	BottomLeft      string
	Bottom          string
	BottomSeparator string
	BottomRight     string
}
