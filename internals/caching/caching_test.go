package caching

import (
	"fmt"
	"testing"
	"time"
)

type complexData struct {
	message string
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

func TestCachingExpiration(t *testing.T) {
	service1 := NewCustomExpiringCachingService(time.Second * 3)
	service2 := NewCustomExpiringCachingService(time.Minute * 1)
	service3 := NewCustomExpiringCachingService(time.Hour * 1)
	k1 := "hello"
	k2 := "hello2"
	k3 := "hello3"
	d1 := complexData{message: "world1"}
	d2 := complexData{message: "world2"}
	d3 := complexData{message: "world3"}

	service1.Set(k1, d1)
	service2.Set(k2, d2)
	service3.Set(k3, d3)

	v1 := retrieveAndPanic[complexData](t, service1, k1)
	v2 := retrieveAndPanic[complexData](t, service2, k2)
	v3 := retrieveAndPanic[complexData](t, service3, k3)
	if v1.message != d1.message || v2.message != d2.message || v3.message != d3.message {
		t.Error("retrieved value did not match the original value", v1, v2, v3)
	}

	fmt.Println("PASS: immediate access")

	time.Sleep(time.Second * 1)
	v1 = retrieveAndPanic[complexData](t, service1, k1)
	v2 = retrieveAndPanic[complexData](t, service2, k2)
	v3 = retrieveAndPanic[complexData](t, service3, k3)
	if v1.message != d1.message || v2.message != d2.message || v3.message != d3.message {
		t.Error("retrieved value did not match the original value", v1, v2, v3)
		t.FailNow()
	}
	fmt.Println("PASS: delayed access after 1 second")

	time.Sleep(time.Second * 5)
	v2 = retrieveAndPanic[complexData](t, service2, k2)
	v3 = retrieveAndPanic[complexData](t, service3, k3)
	if v2.message != d2.message || v3.message != d3.message {
		t.Error("retrieved value did not match the original value", v1, v2, v3)
	}
	v1, err := RetrieveCachedData[complexData](service1, k1)
	if err == nil {
		t.Error("Failed to expire cache with key:", k1)
		t.FailNow()
	} else {
		t.Log("v1 has expired:", v1, err)
	}
	fmt.Println("PASS: delayed access after 5 seconds")

	time.Sleep(time.Second * 55)
	v3 = retrieveAndPanic[complexData](t, service3, k3)
	if v3.message != d3.message {
		t.Error("retrieved value did not match the original value", v3)
	}
	v2, err = RetrieveCachedData[complexData](service2, k2)
	if err == nil {
		t.Error("Failed to expire cache with key:", k2)
		t.FailNow()
	} else {
		t.Log("v2 has expired:", v2, err)
	}
	fmt.Println("PASS: delayed access after 55 seconds")
}

func retrieveAndPanic[T any](t *testing.T, service ICachingService, key string) T {
	v, err := RetrieveCachedData[T](service, key)
	if err != nil {
		t.Errorf("corrupted cached data for %s\n%v", key, err)
		return v
	}
	return v
}
