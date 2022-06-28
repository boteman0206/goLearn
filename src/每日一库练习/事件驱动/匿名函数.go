package main

import (
	"fmt"
	"github.com/gookit/event"
)

var fnHandler = func(e event.Event) error {
	fmt.Printf("handle event: %s\n", e.Name())
	return nil
}

func Run() {
	// register
	event.On("evt1", event.ListenerFunc(fnHandler), event.High)

}

func main() {
	Run()

	err, e := event.Fire("evt2", event.M{"name": "jack"})
	if err != nil {
		return
	}

	fmt.Println(e.Data(), e.Name(), e.Get("name"))

}
