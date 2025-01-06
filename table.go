package pretty_table

import (
	"github.com/starkeland/pretty-table/table"
)

type Table interface {
	SetCaption(caption string)
	SetHeader(header []string)
	AddRows(rows ...[]string)
	SetFooter(footer []string)
	Render()
}

func NewTable(options ...table.Option) Table {
	t := &table.Table{}
	for _, option := range options {
		option(t)
	}

	return t
}
