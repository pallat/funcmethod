package funcmethod

import (
	"testing"
)

func TestTransformFuncToImplementInterface(t *testing.T) {
	r, err := Handler(work(dosomething))
	if err != nil {
		t.Error("it should not error like this", err)
		return
	}

	if _, ok := r.(map[string]interface{}); !ok {
		t.Error("the return type should be map[string]interface{}")
		return
	}

	if v, ok := r.(map[string]interface{})["message"]; ok {
		if v != "cool" {
			t.Error("it should be ok but got", v)
			return
		}
		return
	}

	t.Error("the result should contains message key")
}

func TestTransformStringToImplementInterface(t *testing.T) {
	w := transform("ok")
	r, err := Handler(w)

	if err != nil {
		t.Error("it should not error like this", err)
		return
	}

	if _, ok := r.(map[string]interface{}); !ok {
		t.Error("the return type should be map[string]interface{}")
		return
	}

	if v, ok := r.(map[string]interface{})["message"]; ok {
		if v != "ok" {
			t.Error("it should be ok but got", v)
			return
		}
		return
	}

	t.Error("the result should contains message key")
}
