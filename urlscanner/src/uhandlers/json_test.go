package uhandlers

import (
	"testing"
)

type Test struct {
	Status string `json:"status"`
}

func TestToJSON(t *testing.T) {
	test := Test{
		Status: "test",
	}
	if ToJSON(test) != `{"status":"test"}` {
		t.Error(`Expected {"status":"test"}`)
	}
}
