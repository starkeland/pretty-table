package table

import (
	"github.com/starkeland/pretty-table/style"
)

type Row struct {
	cells []*Cell
	style *style.CellStyle
}
type RowOption func(*Row)

func WithRowStyle(s *style.CellStyle) func(*Row) {
	return func(r *Row) { r.style = s }
}

func NewRow(cells []*Cell, options ...RowOption) *Row {
	return &Row{cells: cells}
}
