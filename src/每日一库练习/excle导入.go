package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/golang/glog"

	"github.com/spf13/cast"
	"log"
)

func main() {
	//f := excelize.NewFile()
	//// Create a new sheet.
	//index := f.NewSheet("Sheet2")
	//// Set value of a cell.
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	//f.SetCellValue("Sheet1", "B2", 100)
	//// Set active sheet of the workbook.
	//f.SetActiveSheet(index)
	//// Save spreadsheet by the given path.
	//if err := f.SaveAs("Book1.xlsx"); err != nil {
	//	fmt.Println(err)
	//}

	//i := make([]map[string]string, 0)
	//
	//m1 := map[string]string{"name": "jack", "age": "78", "addr": "beijing"}
	//m2 := map[string]string{"name": "jack1", "age": "781", "addr": "shanghai"}
	//m3 := map[string]string{"name": "jack2", "age": "781", "addr": "nanjing"}
	//i2 := append(i, m1, m2, m3)
	//var header = []string{"name", "age", "addr"}
	//
	//ExportExcle(i2, header)
	ExportDemo1()
}

//导出数据
func ExportExcle(data []map[string]string, heard []string) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("导入商品信息")

	lenData := len(heard)
	count := 0
	for i := 65; i < lenData+65; i++ {
		// 填充头部信息
		s := string(i) + cast.ToString(1)
		fmt.Println("s :", s)
		f.SetCellValue("导入商品信息", s, heard[count])
		count++
	}

	//for i := 65; i < lenData+65; i++ {
	//	for d := 1; d < lenData+1; d++ {
	//		toString := cast.ToString(d)
	//		cell := string(i) + toString
	//		fmt.Println("cell: ", cell)
	//	}
	//}
	//f.DeleteSheet("Sheet1") // 如果需要将自己的表放置在第一，就删除sheet1表即可
	f.SetActiveSheet(index)

	err := f.SaveAs("./成绩表1112.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}

func ExportDemo1() {
	f := excelize.NewFile()
	header := []string{"门店ID", "门店名称", "所属大区", "省", "城市", "商品名称", "商品类别", "商品类目", "店内商品分类", "渠道", "状态"}
	for i := 0; i < len(header); i++ {
		glog.Info("cell: ", cast.ToString(65+i)+"1")
		f.SetCellValue("Sheet1", string(65+i)+"1", header[i])
	}

	err := f.SaveAs("./demo1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	//创建excel文件
//	xlsx := excelize.NewFile()
//
//	//创建新表单
//	index := xlsx.NewSheet("成绩表")
//
//	//写入数据
//	data := map[string]string{
//		//学科
//		"B1": "语文",
//		"C1": "数学",
//		"D1": "英语",
//		"E1": "理综",
//
//		//姓名
//		"A2": "啊俊",
//		"A3": "小杰",
//		"A4": "老王",
//
//		//啊俊成绩
//		"B2": "112",
//		"C2": "115",
//		"D2": "128",
//		"E2": "255",
//
//		//小杰成绩
//		"B3": "100",
//		"C3": "90",
//		"D3": "110",
//		"E3": "200",
//
//		//老王成绩
//		"B4": "70",
//		"C4": "140",
//		"D4": "60",
//		"E4": "265",
//	}
//	for k, v := range data {
//		//设置单元格的值
//		xlsx.SetCellValue("成绩表", k, v)
//	}
//
//	//设置默认打开的表单
//	xlsx.SetActiveSheet(index)
//
//	//保存文件到指定路径
//	err := xlsx.SaveAs("./成绩表.xlsx")
//	if err != nil {
//		log.Fatal(err)
//	}
//}
