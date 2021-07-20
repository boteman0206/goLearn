package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
)

//func main() {
//	f := excelize.NewFile()
//	// Create a new sheet.
//	index := f.NewSheet("Sheet2")
//	// Set value of a cell.
//	f.SetCellValue("Sheet2", "A2", "Hello world.")
//	f.SetCellValue("Sheet1", "B2", 100)
//	// Set active sheet of the workbook.
//	f.SetActiveSheet(index)
//	// Save spreadsheet by the given path.
//	if err := f.SaveAs("Book1.xlsx"); err != nil {
//		fmt.Println(err)
//	}
//}

func main() {
	//创建excel文件
	xlsx := excelize.NewFile()

	//创建新表单
	index := xlsx.NewSheet("成绩表")

	//写入数据
	data := map[string]string{
		//学科
		"B1": "语文",
		"C1": "数学",
		"D1": "英语",
		"E1": "理综",

		//姓名
		"A2": "啊俊",
		"A3": "小杰",
		"A4": "老王",

		//啊俊成绩
		"B2": "112",
		"C2": "115",
		"D2": "128",
		"E2": "255",

		//小杰成绩
		"B3": "100",
		"C3": "90",
		"D3": "110",
		"E3": "200",

		//老王成绩
		"B4": "70",
		"C4": "140",
		"D4": "60",
		"E4": "265",
	}
	for k, v := range data {
		//设置单元格的值
		xlsx.SetCellValue("成绩表", k, v)
	}

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	err := xlsx.SaveAs("./成绩表.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
