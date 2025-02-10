package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const GlobalConfigFile = "global_config.json"

// GlobalConfig 全局配置结构体
type GlobalConfig struct {
	Config string     `json:"config"`
	Menus  []MenuItem `json:"menus"`
}

// MenuItem 菜单项结构体
type MenuItem struct {
	Key      string     `json:"key"`
	Title    string     `json:"title"`
	Children []MenuItem `json:"children,omitempty"`
}

var (
	GlobalConfigModel *GlobalConfig // 全局配置
)

// initConfig 加载配置文件
func initConfig() error {
	dir, err := getCurrentDirectory() // 获取当前文件目录
	if err != nil {
		return err
	}
	// 拼接路径 读取文件
	data, err := os.ReadFile(filepath.Join(dir, GlobalConfigFile))
	if err != nil {
		return err
	}
	// 解析 JSON
	err = json.Unmarshal(data, &GlobalConfigModel)
	if err != nil {
		return fmt.Errorf("解析配置文件失败:%v", err)
	}
	return nil
}

// GetMenus 获取菜单项
func GetMenus() []MenuItem {
	if GlobalConfigModel == nil {
		err := initConfig()
		if err != nil {
			return nil
		}
	}
	return GlobalConfigModel.Menus
}

func updateConfig(config GlobalConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化配置文件失败:%v", err)
	}
	dir, err := getCurrentDirectory() // 获取当前文件目录
	if err != nil {
		return err
	}
	// 拼接路径 写入文件
	err = os.WriteFile(filepath.Join(dir, GlobalConfigFile), data, 0644)
	if err != nil {
		return fmt.Errorf("写入配置文件失败:%v", err)
	}
	return nil
}

// 获取当前所在的目录
func getCurrentDirectory() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("无法获取当前文件路径")
	}
	dir := filepath.Dir(filename) // 获取目录
	return dir, nil
}
