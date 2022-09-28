package cache

import (
	"fmt"
	"testing"
	"time"
)

func getMethod1(args ...any) (any, error) {
	return true, nil
}

func getMethod2(args ...any) (any, error) {
	return true, nil
}

func TestNewNilParameters(t *testing.T) {
	c := New(nil)
	if !c.t.IsZero() {
		t.Errorf("nil value expected for time.Time parameter, got: %v", c.t)
	}
	if c.d != 0 {
		t.Errorf("nil value expected for time.Duration parameter, got: %v, want: 0", c.d)
	}
}

func TestNewWithParameters(t *testing.T) {
	tt := time.Date(2022, time.September, 28, 0, 0, 0, 0, time.UTC)
	parameters := Parameters{
		100,
		tt,
	}
	c := New(&parameters)
	if c.t != tt {
		t.Errorf("wrong value for time parameter `t`, got: %v, want: %v", c.t, tt)
	}
	if c.d != 100 {
		t.Errorf("wrong value for duration parameter `d`, got: %v, want: %d", c.d, 100)
	}
}

func TestCache_CanBeUpdateNilParameters(t *testing.T) {
	c := New(nil)
	if c.CanBeUpdate() {
		t.Error("can't update cache with nil *Parameters, got: true, want: false")
	}
}

func TestCache_CanBeUpdateWithParameters(t *testing.T) {
	tt := time.Date(2022, time.September, 28, 0, 0, 0, 0, time.UTC)
	parameters := Parameters{
		100,
		tt,
	}
	c := New(&parameters)

	c.Cache(getMethod1, "okok", false, 12)

	nowFunc = func() time.Time {
		return time.Date(2022, time.September, 28, 0, 0, 0, 0, time.UTC)
	}
	result := c.CanBeUpdate()
	if result {
		t.Errorf("cache can't be updated, got: %v, want: %v", result, false)
	}

	nowFunc = func() time.Time {
		return time.Date(2022, time.September, 28, 0, 0, 101, 0, time.UTC)
	}
	result = c.CanBeUpdate()
	if !result {
		t.Errorf("cache can be updated, got: %v, want: %v", result, true)
	}

}

func TestCache_CacheNoParameters(t *testing.T) {
	c := New(nil)

	if _, ok := c.Get("getMethod1", "okok", false, 12); ok {
		t.Error("getMethod1 is not cached yet.")
	}
	c.Cache(getMethod1, "okok", false, 12)
	if _, ok := c.Get("getMethod1", "okok", false, 12); !ok {
		t.Error("getMethod1 have been cached, it wasn't found.")
	}
	if _, ok := c.Get("getMethod1", "okok", false, 13); ok {
		t.Error("getMethod1 with such arguments is not cached yet.")
	}
}

func TestCache_CacheWithParameters(t *testing.T) {
	tt := time.Date(2022, time.September, 28, 0, 0, 0, 0, time.UTC)
	parameters := Parameters{
		100,
		tt,
	}
	c := New(&parameters)

	fmt.Println(c)
}

func TestCache_Get(t *testing.T) {
	c := New(nil)
	c.Cache(getMethod1, "okok", false, 12)
	c.Cache(getMethod1, "okok", false, 13)
	c.Cache(getMethod2, "okok", false, 12, "XX")

	testCases := []struct {
		method string
		args   []any
		isIn   bool
	}{
		{"getMethod1", []any{"okok", false, 12}, true},
		{"getMethod1", []any{"okok", false, 13}, true},
		{"getMethod1", []any{"okok", false, 14}, false},
		{"getMethod2", []any{"okok", false, 12, "XX"}, true},
		{"getMethod2", []any{"okok", false, 12}, false},
		{"getMethod3", []any{"okok", false, 12, "XX"}, false},
	}
	for _, x := range testCases {
		res, ok := c.Get(x.method, x.args...)
		if ok != x.isIn {
			t.Errorf("unexpected get return, got: %v, want: %v", ok, x.isIn)
		}
		if x.isIn {
			if res != true {
				t.Errorf("unexpected cached value, got: %v, want: %v", res, true)
			}
		}
	}
}
