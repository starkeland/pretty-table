package main

import (
	prettyTable "github.com/starkeland/pretty-table"
	tableOption "github.com/starkeland/pretty-table/table"
)

func main() {
	table := prettyTable.NewTable(tableOption.ShowSeqColumn(), tableOption.ShowHR())
	table.SetCaption("Hello, World!")
	table.SetHeader([]string{"Name", "Age"})
	table.AddRows([][]string{
		{"Alice", "20"},
		{"Bob", "21"},
		{"Charlie", "22"},
		{"Dave", "23"},
		{"Eve", "24"},
		{"Frank", "25"},
		{"Grace", "26"},
		{"Heidi", "27"},
		{"Ivan", "28"},
		{"Judy", "29"},
		{"Kevin", "30"},
		{"Lily", "31"},
		{"Mallory", "32"},
		{"Nancy", "33"},
		{"Oscar", "34"},
		{"Peggy", "35"},
		{"Quentin", "36"},
		{"Ruth", "37"},
		{"Steve", "38"},
		{"Trent", "39"},
		{"Uma", "40"},
		{"Victor", "41"},
		{"Wendy", "42"},
		{"Xander", "43"},
		{"Yvonne", "44"},
		{"Zelda", "45"},
	}...)
	table.SetFooter([]string{"Total", "63"})
	table.Render()
}
