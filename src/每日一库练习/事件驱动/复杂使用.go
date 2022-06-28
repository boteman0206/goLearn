package main

import (
	"fmt"

	"github.com/gookit/event"
)

type MySubscriber struct {
	// ooo
}

type MyListener struct {
	// userData string
}

func (l *MyListener) Handle(e event.Event) error {
	e.Set("result", "OK")
	return nil
}

func (s *MySubscriber) SubscribedEvents() map[string]interface{} {
	return map[string]interface{}{
		"e1": event.ListenerFunc(s.e1Handler),
		"e2": event.ListenerItem{
			Priority: event.AboveNormal,
			Listener: event.ListenerFunc(func(e event.Event) error {
				return fmt.Errorf("an error")
			}),
		},
		"e3": &MyListener{},
	}
}

func (s *MySubscriber) e1Handler(e event.Event) error {
	e.Set("e1-key", "val1")
	return nil
}

func main() {
	event.AddSubscriber(&MySubscriber{}) //todo 订阅事件

	err, e := event.Fire("e1", event.M{"e1": "e11"})
	if err != nil {
		fmt.Println("e1 error: ", err.Error())
	}
	fmt.Println("e1 data: ", e.Data())

	err, e2 := event.Fire("e2", event.M{"e2": "e12"})
	if err != nil {
		fmt.Println("e2 error: ", err.Error())
	}
	fmt.Println("e2 data: ", e2.Data())

	err, e3 := event.Fire("e3", event.M{"e3": "e13"})
	if err != nil {
		fmt.Println("e3 error: ", err.Error())
	}
	fmt.Println("e3 data: ", e3.Data())

}
