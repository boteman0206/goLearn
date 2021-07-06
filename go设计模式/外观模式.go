package main

import "fmt"

type API interface {
	Test() string
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type aModuleImpl struct{}

//AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct{}

//BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func main() {

	api := NewAPI()

	test := api.Test()
	fmt.Print("test: ", test)
}
