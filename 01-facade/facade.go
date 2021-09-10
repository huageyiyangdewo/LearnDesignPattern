package facade

import "fmt"

func NewAPI() Api {
	return &apiImpl{
		a: NewAModuleApi(),
		b: NewBModuleApi(),
	}
}

// Api 定义外观角色
type Api interface {
	Test() string
}

type apiImpl struct {
	a AModuleApi
	b BModuleApi
}

func (a *apiImpl)Test() string {
	aRet := a.a.Test()
	bRet := a.b.Test()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}


// AModuleApi 具体的子系统
type AModuleApi interface {
	Test() string
}

type aModuleImpl struct {}

func (a *aModuleImpl)Test() string {
	return "A module running"
}

func NewAModuleApi() AModuleApi {
	return &aModuleImpl{}
}

// BModuleApi 具体的子系统
type BModuleApi interface {
	Test() string
}

type bModuleImpl struct {}

func (a *bModuleImpl)Test() string {
	return "B module running"
}

func NewBModuleApi() BModuleApi {
	return &bModuleImpl{}
}


