package caching

import (
	"testing"
	"time"
)

func TestCachingExpiration(t *testing.T) {
	c := NewCustomExpiringCachingService(time.Second * 3)
	c.Set("hello", "world")

	time.Sleep(time.Second * 1)
	v, err := c.Get("hello")
	if ogValue, ok := v.(string); !ok {
		t.Error("invalid returned data type")
		return
	} else {
		if err != nil || ogValue != "world" {
			t.Error("failed to retrieve cache:", v, err)
		}
	}

	time.Sleep(time.Second * 2)
	_, err = c.Get("hello")
	if err == nil {
		t.Error("failed to expire cache:", err)
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
		return
	}
	if ogValue, ok := rd.(map[string]string); !ok {
		t.Error("invalid data type", ogValue)
	} else {
		if ogValue["message"] != "Hello World" {
			t.Error("filed to parse cached bytes:", err)
		}
	}
}

func TestCachingComplexData(t *testing.T) {
	type complexData struct {
		message string
	}

	c := NewCustomExpiringCachingService(time.Second * 3)
	key := "hello-2"
	if err := c.Set(key, complexData{message: "Hello World"}); err != nil {
		t.Error("unable to cache data:", err)
	}

	rd, err := c.Get(key)
	if err != nil {
		t.Error("unable to retrieve data:", err)
	}
	if d, ok := rd.(complexData); !ok {
		t.Error("invalid data type:", d)
	} else if d.message != "Hello World" {
		t.Error("data got corrupted:", d)
	}
}
