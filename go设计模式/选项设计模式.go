package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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

	New(SetTimeout(3*time.Second), SetCheckReadyTimeout(5*time.Microsecond))
}

//================grpc连接池的案例=====================

const (
	defaultTimeout    = 100 * time.Second
	checkReadyTimeout = 5 * time.Second
	heartbeatInterval = 20 * time.Second
)

type ConnectionTracker struct {
	sync.RWMutex
	timeout           time.Duration
	checkReadyTimeout time.Duration
	heartbeatInterval time.Duration

	ctx    context.Context
	cannel context.CancelFunc
}

type TrackerOption func(*ConnectionTracker)

func SetTimeout(timeout time.Duration) TrackerOption {
	return func(o *ConnectionTracker) {
		o.timeout = timeout
	}
}

// SetCheckReadyTimeout custom checkReadyTimeout
func SetCheckReadyTimeout(timeout time.Duration) TrackerOption {
	return func(o *ConnectionTracker) {
		o.checkReadyTimeout = timeout
	}
}

func New(opts ...TrackerOption) *ConnectionTracker {
	ctx, cannel := context.WithCancel(context.Background())
	ct := &ConnectionTracker{
		ctx:               ctx,
		cannel:            cannel,
		timeout:           defaultTimeout,
		checkReadyTimeout: checkReadyTimeout,
		heartbeatInterval: heartbeatInterval,
	}
	for _, opt := range opts {
		opt(ct) // 选项模式
	}

	fmt.Println("ct data: ", ct)
	return ct
}
