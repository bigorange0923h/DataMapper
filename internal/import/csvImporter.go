package _import

import (
	"DataMapper/config"
	"DataMapper/internal/api"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// CSVImporter csv文件导入处理器
type CSVImporter struct {
}

func (c CSVImporter) Importer(configName string) ([]map[string]string, error) {
	mapperConfig, err := config.LoadMapperConfig(configName)
	if err != nil {
		//todo 获取配置失败日志处理
		return nil, err
	}
	file, err := os.Open(mapperConfig.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	// todo 这里未来要处理默认分隔符
	reader.Comma = rune(mapperConfig.SplitCharacter[0])
	// 获取首行作为表头
	headers, err := reader.Read()
	//var data map[string]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			// 如果到达文件末尾,停止读取
			break
		} else if err != nil {
			fmt.Println("Error reading CSV:", err)
			break
		}
		fmt.Println(record)
		row := make(map[string]string)
		mappings := mapperConfig.GetMappings()
		for i, value := range record {
			row[mappings[headers[i]]] = value
		}
		fmt.Println(row)
		//data = append(data, row)
		json, _ := json.Marshal(row)
		request := api.Request{mapperConfig.RequestUrl, nil, nil, mapperConfig.RequestType, string(json)}
		request.Do()
	}
	return nil, nil
}

func (c CSVImporter) GetFiled(headerRow int, filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("文件打开失败:%w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentLine := 0
	for scanner.Scan() {
		if currentLine == headerRow {
			// todo 这里未来要处理默认分隔符
			return strings.Split(scanner.Text(), ","), nil
		}
	}
	return nil, nil
}
