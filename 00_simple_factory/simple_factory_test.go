package simple_factory

import "testing"

func TestHelloAPI_Say(t *testing.T) {
	h := NewAPI(2)
	if h.Say("test") != "Hello, test" {
		t.Fatal("TestHelloAPI_Say failed")
	}
}

func TestHiAPI_Say(t *testing.T) {
	h := NewAPI(1)
	if h.Say("test") != "Hi, test" {
		t.Fatal("TestHelloAPI_Say failed")
	}
}