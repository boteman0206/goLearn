package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}
func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.WithMessage(err, "could not read config")
}
func main() {
	_, err := ReadConfig()
	if err != nil {
		err := errors.WithStack(err) // todo 附加堆栈信息

		log.Errorf("%+v\n", err) // 打印堆栈信息
		fmt.Println("简单的错误信息：", err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}

	err = errors.New("test异常")
	err = errors.WithStack(err)

	err = errors.Cause(err)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v \n", errors.Cause(err)))
		return
	}
}
