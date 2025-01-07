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
		{"etcd", "etcd://10.59.15.13:2381,10.59.15.14:2381,10.59.15.15:2381/da.skylar"},
		{"db", "mongo://admin:CVCPF7JUEbnTLr.1_ZAa1@10.59.15.13:27017,10.59.15.14:27017,10.59.15.15:27017/da_skylar_moss?alias=moss&authdb=admin&create_table=true"},
		{"cache", "redis://:u39DHyOqkP23nt2P_ZAa1@10.59.15.13:16379,10.59.15.13:16380,10.59.15.14:16379,10.59.15.14:16380,10.59.15.15:16379,10.59.15.15:16380"},
	}...)
	table.Render()
}
