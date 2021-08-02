package main

import "fmt"

type Options struct {
	str1 string
	str2 string
	str3 string
	int1 int32
	int2 int32
}

// 申明一个函数类型的变量用于传递参数
type Option func(opts *Options)

func InitOptions(opts ...Option) {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}

	fmt.Println("options: ", options)
}

func WithStringOption1(str string) Option {
	return func(opts *Options) {
		opts.str1 = str
	}
}

func WithStringOption2(str string) Option {
	return func(opts *Options) {
		opts.str2 = str
	}
}

func WithStringOption3(str string) Option {
	return func(opts *Options) {
		opts.str3 = str
	}
}
func WithStringOption4(int1 int32) Option {
	return func(opts *Options) {
		opts.int1 = int1
	}
}
func WithStringOption5(int1 int32) Option {
	return func(opts *Options) {
		opts.int2 = int1
	}
}

func main() {

	InitOptions(WithStringOption1("str1test"), WithStringOption2("str2test"), WithStringOption1("changetest"))
}
