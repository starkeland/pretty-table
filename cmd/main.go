package main

import (
	prettyTable "github.com/starkeland/pretty-table"
	tableOption "github.com/starkeland/pretty-table/table"
)

func main() {
	table := prettyTable.NewTable(tableOption.Options{
		// ShowSeqColumn: true,
		ShowHR: true,
	})
	// table.SetCaption(" 配置信息")
	table.SetHeader([]string{"配置项", "值"})
	table.AddRows([][]string{
		{"a", "b"},
		{"c", "d"},
		{"e", "f"},
		{"g", "h"},
	}...)
	table.Render()
}
