package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// MappingsModel 字段映射结构体
type MappingsModel struct {
	SourceName string `json:"sourceName"`
	SourceType string `json:"sourceType"`
	TargetName string `json:"targetName"`
	TargetType string `json:"targetType"`
}

// MappingConfigModel 配置结构体
type MappingConfigModel struct {
	Mappings       []MappingsModel `json:"mappings"`
	FilePath       string          `json:"filePath"`
	FileType       string          `json:"fileType"`
	SplitCharacter string          `json:"splitCharacter"`
	RequestUrl     string          `json:"requestUrl"`
	RequestType    string          `json:"requestType"`
}

func (m MappingConfigModel) GetMappings() map[string]string {
	if m.Mappings == nil {
		return nil
	}
	mappings := make(map[string]string)
	for _, mappingsModels := range m.Mappings {
		mappings[mappingsModels.SourceName] = mappingsModels.TargetName
	}
	return mappings
}

// LoadMapperConfig 加载映射配置
func LoadMapperConfig(configName string) (*MappingConfigModel, error) {
	file, err := os.Open(configName)
	if err != nil {
		return nil, fmt.Errorf("获取配置文件失败:%v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := &MappingConfigModel{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败:%v", err)
	}
	return config, nil
}
