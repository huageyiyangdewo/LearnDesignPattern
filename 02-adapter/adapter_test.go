package adapter

import "testing"

func TestMotorAdapter(t *testing.T) {
	name := "electric"
	adapter := MotorAdapter(name)
	if adapter.Drive() != "electricDrive" {
		t.Fatal("test failed")
	}
}
