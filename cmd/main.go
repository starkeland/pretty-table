package main

import (
	prettyTable "github.com/starkeland/pretty-table"
	"github.com/starkeland/pretty-table/style"
	"github.com/starkeland/pretty-table/table"
)

func main() {
	// t := prettyTable.NewTable()
	t := prettyTable.NewTable(table.ShowSeqColumn())
	t.SetCaption("配置信息")
	t.SetHeader([]string{"配置项", "值"})
	t.AddRows(
		table.NewStringRow([]string{"mysql", "mysql://root:12345@127.0.0.1:3306"}),
		table.NewRow([]*table.Cell{
			table.NewCell("redis", table.WithCellStyle(&style.CellStyle{
				FgColor: style.FgCyan,
			})),
			table.NewCell("redis://root:123456@127.0.0.1:6379"),
		}),
		table.NewStringRow([]string{"etcd", "etcd://root:3acf823AICnqsLQc29ac:2379"}),
	)
	t.SetTableStyle(style.DefaultTableStyle)
	t.Render()
}
