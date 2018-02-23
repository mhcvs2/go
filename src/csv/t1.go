package main

import (
	"encoding/csv"
	"os"
)

func t11() {
	f, err := os.Create("/root/github/go/src/csv/test.csv")//创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)//创建一个新的写入文件流
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)//写入数据
	w.Flush()
}

func t12() {
	f, err := os.Create("/root/github/go/src/csv/test.csv")//创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)//创建一个新的写入文件流
	defer w.Flush()

	w.Write([]string{"序号", "国家", "number"})

	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)//写入数据
}

func t13() {
	f, err := os.OpenFile("/root/github/go/src/csv/test.csv", os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)//创建一个新的写入文件流
	defer w.Flush()
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)//写入数据
}

func main() {
	t13()
}