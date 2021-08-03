package 常用log日志包的集合

import (
	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func init() {
	hookConfig := logredis.HookConfig{
		Host:     "localhost",
		Key:      "mykey",
		Format:   "v0",
		App:      "syslog",
		Hostname: "testapp",
		Password: "demo",
		Port:     6379,
		DB:       0,
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

func main() {
	logrus.Info("just some info logging...")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"foo":    "bar",
		"this":   "that",
	}).Info("additional fields are being logged as well")

	logrus.SetOutput(ioutil.Discard)
	logrus.Error("This will only be sent to Redis")
}
