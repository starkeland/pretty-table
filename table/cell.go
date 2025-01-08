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

type CellOption func(*Cell)

func WithCellStyle(cs *style.CellStyle) CellOption {
	return func(c *Cell) {
		c.style = cs
	}
}

func NewCell(text string, options ...CellOption) *Cell {
	c := &Cell{
		text: fmt.Sprintf(" %s ", text),
	}
	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Cell) Width() int {
	return runewidth.StringWidth(c.text)
}

func (c *Cell) Render(width int) string {
	// If cell has no style and there is a default style for the row or column.
	var builder strings.Builder

	if c.style != nil {
		var paddingLeft, paddingRight int
		switch c.style.Align {
		case style.TextAlignLeft:
			paddingRight = width - runewidth.StringWidth(c.text)
		case style.TextAlignCenter:
			paddingLeft = (width - runewidth.StringWidth(c.text)) / 2
			paddingRight = width - runewidth.StringWidth(c.text) - paddingLeft
		case style.TextAlignRight:
			paddingLeft = width - runewidth.StringWidth(c.text)
		}

		for i := 0; i < paddingLeft; i++ {
			builder.WriteString(" ")
		}
		builder.WriteString(c.text)
		for i := 0; i < paddingRight; i++ {
			builder.WriteString(" ")
		}

		return c.style.Apply(builder.String())
	}

	builder.WriteString(c.text)
	for i := 0; i < width-runewidth.StringWidth(c.text); i++ {
		builder.WriteString(" ")
	}

	return builder.String()
}
