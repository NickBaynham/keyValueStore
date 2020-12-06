package keyValueStore

import (
	"errors"
	"testing"
)

func TestDelete(t *testing.T) {
	err:= Put("key", "value")
	if err != nil {
		t.Error("Put key/value failed.")
	}
	err= Delete("key")
	if errors.Is(err, ErrorNoSuchKey) {
		t.Error("No Such Key Error")
	}
}

func TestGet(t *testing.T) {
	var _ = Put("key", "value")
	value, err := Get("key")
	if err != nil {
		t.Error("Get Failed.")
	}
	if value != "value" {
		t.Error("Expected \"value\" but got \"", value + "\"")
	}
}

func TestPut(t *testing.T) {
	err := Put("key", "value")
	if err != nil {
		t.Error("Put failed!")
	}
}
