package main

import (
	prettyTable "github.com/starkeland/pretty-table"
	"github.com/starkeland/pretty-table/style"
	"github.com/starkeland/pretty-table/table"
)

func main() {
	t := prettyTable.NewTable(table.ShowSeqColumn(), table.ShowHr())
	t.SetCaption("配置信息")
	t.SetHeader([]string{"配置项", "值"})
	t.AddRow([]string{"mysql", "mysql://root:123456@127.0.0.1:3306"})
	t.AddRow([]string{"redis", "redis://root:123456@127.0.0.1:6379"})
	t.AddRow([]string{"etcd", "etcd://root:3acf823AICnqsLQc29ac:2379"})
	t.SetStyle(style.DefaultTableStyle)
	t.Render()
}
