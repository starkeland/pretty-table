package table

import (
	"fmt"
	"strings"
)

type Cell struct {
	text string
}

func NewCell(text string) Cell {
	return Cell{
		text: fmt.Sprintf(" %s ", text),
	}
}

func (c Cell) Width() int {
	return len(c.text)
}

func (c Cell) Text(width int) string {
	var builder strings.Builder
	builder.WriteString(c.text)
	for i := 0; i < width-len(c.text); i++ {
		builder.WriteString(" ")
	}
	return builder.String()
}
