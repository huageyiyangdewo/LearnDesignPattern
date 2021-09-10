package simple_factory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(n int) API {
	if n == 1 {
		return &hiAPI{}
	} else if n == 2 {
		return &helloAPI{}
	}

	return nil
}

type hiAPI struct {}

func (h *hiAPI)Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct {}

func (h *helloAPI)Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}