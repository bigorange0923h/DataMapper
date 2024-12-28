package _import

// DataImport 数据导入接口
type DataImport interface {
	Importer(filepath string) ([]map[string]string, error)
}
