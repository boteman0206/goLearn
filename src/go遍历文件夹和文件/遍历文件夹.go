package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

/*
	我们可以利用递归和ioutil.ReadDir(遍历当前所有文件)，IsDir(判断是否是目录)方法，来遍历整个文件夹

*/

var (
	dirName = "./data/img/"
)

func main() {

	getAllDir()

	// 递归判断文件内容
	//IsFiles(dirName)

	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println("获取当前的工作目录：", getwd)

}

// 讀取所有的文件夾
func getAllDir() []fs.FileInfo {

	dir, err := ioutil.ReadDir(dirName)

	if err != nil {
		fmt.Println("read dir err: ", err.Error())
		return nil
	}

	for i := range dir {
		name := dir[i].Name()
		fmt.Println("name: ", name)
		isDir := dir[i].IsDir()
		fmt.Println("idDir: ", isDir)

		if isDir {
			sonPath := path.Join(dirName, name)
			getAllFile(sonPath)
		}

	}

	return dir
}

// 获取子文件下夾下面所有的文件
func getAllFile(sonPath string) {
	dir, err := ioutil.ReadDir(sonPath)
	if err != nil {
		fmt.Println("子文件获取dir： ", err.Error())
		return
	}

	if err != nil {
		return
	}
	for i := range dir {
		sonFile := dir[i]
		fmt.Println(sonFile.Name())
		fmt.Println(sonFile.IsDir())
		fmt.Println(sonFile.Size())
		if !sonFile.IsDir() {
			//打开对应文件架下面的图片
		}
	}

}

// 打开图片
//func OpenImage(file string) {
//	if runtime.GOOS == "windows" {
//		cmd := exec.Command("start", file)
//		err := cmd.Start()
//		fmt.Println(" err:", err)
//	} else {
//		if runtime.GOOS == "linux" {
//			cmd := exec.Command("eog", file)
//			_ = cmd.Start()
//		} else {
//			cmd := exec.Command("open", file)
//			_ = cmd.Start()
//		}
//	}
//}

//递归调用读取文件下面的文件
func IsFiles(dirname string) {
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		return
	}

	for i := range dir {
		join := path.Join(dirname, dir[i].Name())

		if dir[i].IsDir() {
			fmt.Println("文件架名称：", join)
			IsFiles(join)
		} else {
			fmt.Println("文件内容：", join)
		}
	}
}
