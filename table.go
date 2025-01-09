package pretty_table

import (
	"github.com/starkeland/pretty-table/style"
	"github.com/starkeland/pretty-table/table"
)

type Table interface {
	SetCaption(caption string)
	SetHeader(header []string)
	AddStringRow(row []string)
	AddStringRows(rows ...[]string)
	AddRow(row *table.Row)
	AddRows(rows ...*table.Row)
	SetFooter(footer []string)
	SetTableStyle(tableStyle *style.TableStyle)
	Render()
}

func NewTable(options ...table.Option) Table {
	t := &table.Table{}
	t.SetTableStyle(style.DefaultTableStyle)
	for _, option := range options {
		option(t)
	}

	return t
}
