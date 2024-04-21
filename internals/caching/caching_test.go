package caching

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCachingExpiration(t *testing.T) {
	c := NewCustomExpiringCachingService(time.Second * 3)
	c.Set("hello", "world")
	v, err := c.Get("hello")
	if err != nil || string(v) == "world" {
		t.Error("failed to retrieve cache:", v, err)
		return
	}

	time.Sleep(time.Second * 1)
	v2, err := c.Get("hello")
	if err != nil || string(v2) == "world" {
		t.Error("failed to retrieve cache:", v, err)
	}

	time.Sleep(time.Second * 2)
	v3, err := c.Get("hello")
	if err == nil || v3 != nil {
		t.Error("failed to expire cache:", v3, err)
	}
}

func TestCachingMaps(t *testing.T) {
	c := NewCustomExpiringCachingService(time.Second * 3)
	key := "hello-3"
	d := map[string]string{"message": "Hello World"}
	if err := c.Set(key, d); err != nil {
		t.Error("failed to save data:", err)
		return
	}

	rd, err := c.Get(key)
	if err != nil {
		t.Error("failed to retrieve data:", err)
	}
	data := map[string]string{}
	if err = json.Unmarshal([]byte(rd), &data); err != nil {
		t.Error("filed to parse cached bytes:", err)
	}
	fmt.Println(data)
	if data["message"] != "Hello World" {
		t.Error("filed to parse cached bytes:", err)
	}
}

func TestCachingComplexData(t *testing.T) {
	type complexData struct {
		message string
	}

	c := NewCustomExpiringCachingService(time.Second * 3)
	key := "hello-2"
	b, _ := json.Marshal(complexData{message: "Hello World"})
	if err := c.Set(key, b); err != nil {
		t.Error("failed to save data:", err)
		return
	}

	rd, err := c.Get(key)
	if err != nil {
		t.Error("failed to retrieve data:", err)
	}
	data := complexData{}
	if err = json.Unmarshal([]byte(rd), &data); err != nil {
		t.Error("filed to parse cached bytes:", err)
	}
	fmt.Println(data)
	if data.message != "Hello World" {
		t.Error("filed to parse cached bytes:", err)
	}
}
