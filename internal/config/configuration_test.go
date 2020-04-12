package config

import "testing"

func TestSetAndGetKey(t *testing.T) {
	config := NewConfiguration()

	testKey, testVal := "testKey", "testVal"
	config.Set(testKey, testVal)

	if ret := config.Get(testKey); *ret != testVal {
		t.Errorf("config.Get Expected: %s, Actual: %s", testVal, *ret)
	}
}

func TestGlobalSetAndGet(t *testing.T) {
	testKey, testVal := "testKey", "testVal"
	Config.Set(testKey, testVal)

	if ret := Config.Get(testKey); *ret != testVal {
		t.Errorf("config.Get Expected: %s, Actual: %s", testVal, *ret)
	}
}
