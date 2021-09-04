package xfile

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Write struct {
	Name    string
	Tables  []string
	Schemas map[string][]string
	Datas   map[string][]map[string]string
}

func (w *Write) Dump() (string, bool) {
	ext := w.Ext()

	switch ext {
	case "txt":
		w.Txt()
	case "csv":
		w.Csv()
	case "xlsx":
		w.Xlsx()
	case "xls":
		w.Xlsx()
	}
	return "", false
}

func (w *Write) Ext() string {
	filenames := strings.Split(w.Name, ".")
	ext := filenames[len(filenames)-1]
	return ext
}

func (w *Write) Csv() {
	file, err := os.OpenFile(w.Name, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	wf := csv.NewWriter(file)
	// 写入UTF-8 BOM，防止中文乱码
	for _, row := range w.Schemas {
		wf.Write(row)
	}
	for _, data := range w.Datas {
		for _, rows := range data {
			tmp := []string{}
			for _, schema := range w.Schemas {
				for _, sch := range schema {
					if content, ok := rows[sch]; ok {
						tmp = append(tmp, content)
					} else {
						tmp = append(tmp, "")
					}
				}
			}
			wf.Write(tmp)
		}
	}
	// 写文件需要flush，不然缓存满了，后面的就写不进去了，只会写一部分
	wf.Flush()
}

func (w *Write) Xlsx() {
	f := excelize.NewFile()
	for _, table := range w.Tables {
		if schema, ok := w.Schemas[table]; ok {
			index := f.NewSheet(table)
			if data, ok2 := w.Datas[table]; ok2 {
				err := f.SetSheetRow(table, "A1", &schema)
				fmt.Println(err)
				for index2, row := range data {
					tmp := []string{}
					for _, sch := range schema {
						if content, ok3 := row[sch]; ok3 {
							tmp = append(tmp, content)
						} else {
							tmp = append(tmp, "")
						}
					}

					err2 := f.SetSheetRow(table, fmt.Sprintf("A%d", index2+2), &tmp)
					fmt.Println(err2)
				}
			}
			if index == 0 {
				f.SetActiveSheet(index)
			}
		}
	}

	fmt.Println("w.Name:", w.Name)
	if err := f.SaveAs(w.Name); err != nil {
		fmt.Println(err)
	}
}

func (w *Write) Txt() {
	f, err := os.Create(w.Name)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, schema := range w.Schemas {
			if _, err := f.Write([]byte(strings.Join(schema, ","))); err != nil {
				fmt.Println(err)
			}
		}
		for _, rows := range w.Datas {
			for _, data := range rows {
				tmp := []string{}
				for _, schema := range w.Schemas {
					for _, sch := range schema {
						if content, ok := data[sch]; ok {
							tmp = append(tmp, content)
						} else {
							tmp = append(tmp, "")
						}
					}
				}
				_, err2 := f.Write([]byte(strings.Join(tmp, ",")))
				fmt.Println(err2)
			}
		}
		fmt.Println(err)
	}
}
