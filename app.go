package main

import (
	"DataMapper/config"
	_import "DataMapper/internal/import"
	"context"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello1 %s, It's show time!", name)
}

func getDataImporter(fileType string) (_import.DataImport, error) {
	switch fileType {
	case "csv":
		return &_import.CSVImporter{}, nil
	case "excel":
		return nil, errors.New("暂不支持该类型")
	case "json":
		return nil, errors.New("暂不支持该类型")
	default:
		return nil, errors.New("文件类型不存在")
	}

}

// LoadMappingConfig 加载映射配置
func (a *App) LoadMappingConfig(fileType string, filePath string) ([]config.MappingsModel, error) {
	fmt.Println("filePath:" + filePath)
	// todo 后面开发基于配置优先判定配置中是否存在,再是读取文件来解析 mapperConfig, err := config.LoadMapperConfig(title)
	importer, err := getDataImporter(fileType)
	if err != nil {
		return nil, err
	}

	// 如果没有读取到配置则直接读取文件指定行
	filedList, err := importer.GetFiled(0, filePath)
	if err != nil {
		return nil, err
	}
	var mappingsModels []config.MappingsModel
	for i, filed := range filedList {
		mappingsModel := config.MappingsModel{
			Id:         i + 1,
			SourceName: filed,
		}
		mappingsModels = append(mappingsModels, mappingsModel)
	}
	return mappingsModels, nil
}

func (a *App) importFile(fileType string, filePath string, mappings []config.MappingsModel) {
	// todo 导入文件
}

// SelectExcelFile 选择文件
func (a *App) SelectExcelFile() string {
	options := runtime.OpenDialogOptions{
		Title: "选择 Excel 文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Excel Files",
				Pattern:     "*.xlsx;*.xls;*.csv",
			},
		},
	}

	file, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil || file == "" {
		return ""
	}
	return file
}
