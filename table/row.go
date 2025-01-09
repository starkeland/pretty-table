package table

import (
	"strconv"

	"github.com/starkeland/pretty-table/style"
)

type Row struct {
	cells []*Cell
	style *style.CellStyle
}

type RowOption func(*Row)

func WithRowStyle(s *style.CellStyle) func(*Row) {
	return func(r *Row) {
		if s != nil {
			r.style = s
		}
	}
}

func NewRow(cells []*Cell, options ...RowOption) *Row {
	if len(cells) == 0 {
		return nil
	}
	return &Row{cells: cells}
}

func NewStringRow(cells []string, options ...RowOption) *Row {
	if len(cells) == 0 {
		return nil
	}
	var c []*Cell
	for _, cell := range cells {
		c = append(c, NewCell(cell))
	}
	return &Row{cells: c}
}

func (r *Row) InsertSeqColumn(seqNum int) {
	r.cells = append([]*Cell{NewCell(strconv.Itoa(seqNum))}, r.cells...)
}
