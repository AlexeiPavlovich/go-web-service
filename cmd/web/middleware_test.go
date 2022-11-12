package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteToConsole(t *testing.T) {
	var myH myHandler
	h := WriteToConsole(&myH)
	switch v := h.(type) {
	case http.Handler:
	//do nothing
	default:
		t.Error(fmt.Sprintf("type http.Handler %T", v))

	}
}
