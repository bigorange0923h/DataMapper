package _import

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// CSVImporter csv文件导入处理器
type CSVImporter struct {
}

func (c CSVImporter) Importer(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '|'
	// 获取首行作为表头
	headers, err := reader.Read()
	var data []map[string]string
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
		for i, value := range record {
			row[headers[i]] = value
		}
		data = append(data, row)
	}
	return nil, nil
}
