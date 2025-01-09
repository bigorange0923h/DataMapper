package main

import (
	"DataMapper/config"
	"context"
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

// LoadMappingConfig 加载字段配置
func (a *App) LoadMappingConfig(title string) []config.MappingsModel {
	fmt.Println("title:" + title)
	mapperConfig, err := config.LoadMapperConfig(title)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return mapperConfig.Mappings
}

func (a *App) SelectExcelFile() string {
	options := runtime.OpenDialogOptions{
		Title: "选择 Excel 文件",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Excel Files",
				Pattern:     "*.xlsx;*.xls",
			},
		},
	}

	file, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil || file == "" {
		return ""
	}
	return file
}
