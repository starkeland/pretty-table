package pretty_table

import (
	"github.com/starkeland/pretty-table/style"
	"github.com/starkeland/pretty-table/table"
)

type Table interface {
	SetCaption(caption string)
	SetHeader(header []string)
	AddRow(row []string)
	SetFooter(footer []string)
	SetStyle(s *style.TableStyle)
	Render()
}

func NewTable(options ...table.Option) Table {
	t := &table.Table{}
	t.SetStyle(style.DefaultTableStyle)
	for _, option := range options {
		option(t)
	}

	return t
}
