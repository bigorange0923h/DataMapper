package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type MapperConfigModel struct {
	Mappings       map[string]string `json:"mappings"`
	FilePath       string            `json:"filePath"`
	FileType       string            `json:"fileType"`
	SplitCharacter string            `json:"splitCharacter"`
	RequestUrl     string            `json:"requestUrl"`
	RequestType    string            `json:"requestType"`
}

// LoadMapperConfig 加载映射配置
func LoadMapperConfig(configName string) (*MapperConfigModel, error) {
	file, err := os.Open(configName)
	if err != nil {
		return nil, fmt.Errorf("获取配置文件失败:%v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := &MapperConfigModel{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败:%v", err)
	}
	return config, nil
}
