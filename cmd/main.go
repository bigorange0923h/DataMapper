package main

import (
	"DataMapper/config"
	"fmt"
)

func main() {
	conf, err := config.LoadMapperConfig("C:\\workspace\\GolangProjects\\DataMapper\\config\\mapper_config_template.json")
	if err != nil {
		fmt.Errorf("报错了")
	}
	/*	var fileImporter _import.DataImport
		fileImporter = _import.CSVImporter{}
		data, err := fileImporter.Importer("C:\\Users\\alvhua\\Documents\\数据处理\\Global_2023-07-22_Customer_GPOS_01_00.csv")
		if err != nil {
			return
		}*/
	fmt.Println(conf)
}
