package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

/**
将一些文案，转换成配置文件需要的类型，比如带上id，和其他字段等信息
*/

func ReadFromFile(filePath string) (string, error) {
	configFile, err := os.Open(filePath)
	defer configFile.Close()
	if err != nil {
		return "", err
	}

	fd, err := ioutil.ReadAll(configFile)
	if err != nil {
		return "", err
	}
	return string(fd), nil
}

func ScanFile(fileName string) []IzakayaChatList {
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(file)

	data := make([]IzakayaChatList, 0)
	// 按行读取文件内容
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("Line %d: %s\n", lineNumber, line)
		data = append(data, IzakayaChatList{
			ID:   int64(lineNumber),
			Desc: line,
		})
		lineNumber++
	}
	// 检查是否发生错误
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return data
}

type RandomQuestionsRes struct {
	QuestionID int64  `form:"question_id" json:"question_id"` // 问题id
	Question   string `form:"question" json:"question"`       // 问题
}

type IzakayaChatList struct {
	ID   int64  `json:"id"`
	Desc string `json:"desc"`
	Sort int    `json:"sort"`
}

func main() {

	//fileStr, err := ReadFromFile("/Users/v_pengwei01/sdk/gitProject/goLearn/baidu_gravity/wenjian.txt")
	//if err != nil {
	//	fmt.Printf("ReadFromFile err %s \n", err.Error())
	//}
	//fmt.Println("fileStr", fileStr)

	data := ScanFile("/Users/v_pengwei01/sdk/gitProject/goLearn/wenjian.txt")
	td, _ := json.Marshal(data)
	//fmt.Println(string(td))

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir = filepath.Join(dir, "")
	fileName := filepath.Join(dir, "myfileTest.txt")
	fmt.Println(fileName)

	// 写入文件，前台打印的时候复制有问题
	WriteToFile(fileName, string(td))

}

func WriteToFile(fileName string, data string) {
	// 打开文件
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入数据到文件
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("WriteToFile err : ", err.Error())
		panic(err)
	}

}
