package helpers

import (
	. "testing"
)

func TestImplode(t *T) {
	src := []string{"hello", "world"}
	result := "hello world"
	if Implode(src, " ") != result {
		t.Error("Implode fail")
	}
}

func TestSliceEqual(t *T) {
	src := []string{"hello", "world"}
	eq := []string{"hello", "world"}
	neq := []string{"hello", "worl"}
	if !SliceEqual(src, eq) {
		t.Error("Slice Equal fail")
	}
	if SliceEqual(src, neq) {
		t.Error("Slice Equal fail")
	}
}

func TestExplode(t *T) {
	src := "hello world"
	result := []string{"hello", "world"}
	if !SliceEqual(Explode(src, " "), result) {
		t.Error("Explode fail")
	}
}

func TestKeys(t *T) {
	src := map[string]interface{}{"hello": "1", "world": "2"}
	result := []string{"hello", "world"}
	if !SliceEqual(Keys(src), result) {
		t.Error("Keys fail")
	}
}

func TestP(t *T) {
	p := NewP()
	p.Set("name", "young")
	if p.Get("name") != "young" {
		t.Error("set fail")
	}
	p.Set("age", 15)
	if p.Get("age") != 15 {
		t.Error("set fail")
	}
	if !p.Exists("name") {
		t.Error("exists fail")
	}
	if p.Exists("not Exists") {
		t.Error("exists fail")
	}
	if !SliceEqual(p.Keys(), []string{"name", "age"}) {
		t.Error("keys fail")
	}
	p.Del("age")
	if p.Exists("age") {
		t.Error("del fail")
	}
}
