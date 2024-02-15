package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tealeg/xlsx"
)

func GetArticle(file string) string {
	filename := filepath.Base(file)
	parts := strings.Split(filename, "_")
	if len(parts) > 0 {
		return strings.TrimSpace(parts[0])
	}
	return ""
}

func main() {
	root := "\\\\nas\\databank\\MP_Design_Contractors\\"

	data := make(map[string][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".jpg" || ext == ".png" {
				article := GetArticle(path)
				if article != "" {
					path := strings.Replace(path, root, "", -1)
					citilux_path := "https://citilux.ru/upload/nas/MP_Design_Contractors/" + filepath.ToSlash(path)
					data[article] = append(data[article], citilux_path)
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error reading folder: %v\n", err)
		return
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf("Error creating sheet: %v\n", err)
		return
	}

	for key, values := range data {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetString(key)

		for _, value := range values {
			cell = row.AddCell()
			cell.SetString(value)
		}
	}

	err = file.Save("output.xlsx")
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Println("File saved successfully")
}
