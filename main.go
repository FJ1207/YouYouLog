package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// 打开Excel文件
	f, err := excelize.OpenFile("your_file.xlsx")
	if err != nil {
		log.Fatal("无法打开Excel文件:", err)
	}
	defer f.Close()

	// 选择要读取的工作表
	sheetName := "Sheet1"

	// 要读取的列号 (例如读取B列)
	columnLetter := "B"

	// 获取指定列的数据
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal("获取工作表数据失败:", err)
	}

	// 创建一个字符串切片存储读取到的列数据
	var columnData []string
	for _, row := range rows[1:] { // 跳过标题行 (如果没有标题行可以从rows[0:]开始)
		if len(row) >= 2 { // 确保行中有足够的列
			columnData = append(columnData, row[1]) // B列为索引1
		}
	}

	// 将所有数据拼接成一个字符串
	result := strings.Join(columnData, "")

	// 输出结果
	fmt.Println("拼接后的字符串:")
	fmt.Println(result)
}
