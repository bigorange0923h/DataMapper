package main

import _import "DataMapper/internal/import"

func main() {
	var importer = _import.CSVImporter{}
	importer.Importer("config/mapper_config_template.json")
}
