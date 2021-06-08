package main

import "fmt"

type API interface {
	Say(name string) string
}

type hiAPI struct {
}

func (h *hiAPI) Say(name string) string {
	fmt.Println(" i am hi API")
	return ""
}

type HelloAPI struct {
}

func (hello *HelloAPI) Say(name string) string {
	fmt.Println("i am hello api")
	return ""
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &HelloAPI{}
	} else {
		return nil
	}

}

func main() {
	api := NewAPI(1)
	api.Say("")

	newAPI := NewAPI(2)
	newAPI.Say("")

}
