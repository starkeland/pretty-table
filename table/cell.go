package table

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/starkeland/pretty-table/style"
)

type Cell struct {
	text  string
	style *style.CellStyle
}

func NewCell(text string, cellStyles ...*style.CellStyle) Cell {
	if len(cellStyles) == 0 {
		return Cell{
			text: fmt.Sprintf(" %s ", text),
		}
	}
	return Cell{
		text:  fmt.Sprintf(" %s ", text),
		style: cellStyles[len(cellStyles)-1],
	}
}

func (c Cell) Width() int {
	return runewidth.StringWidth(c.text)
}

func (c Cell) Render(width int) string {
	var builder strings.Builder
	builder.WriteString(c.text)
	for i := 0; i < width-runewidth.StringWidth(c.text); i++ {
		builder.WriteString(" ")
	}

	if c.style != nil {
		return c.style.Apply(builder.String())
	}
	return builder.String()
}
