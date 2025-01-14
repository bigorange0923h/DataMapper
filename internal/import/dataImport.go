package _import

// DataImport 数据导入接口
type DataImport interface {
	Importer(filePath string) ([]map[string]string, error)

	GetFiled(headerRow int, filePath string) ([]string, error)
}
