package config_test

import (
	"testing"

	. "github.com/onsi/gomega"

	c "github.com/mhv2109/uci-impl/internal/config"
)

func TestSetAndGetKey(t *testing.T) {
	g := NewGomegaWithT(t)

	config := c.NewConfiguration()

	testKey, testVal := "testKey", "testVal"
	config.Set(testKey, testVal)

	ret := config.Get(testKey)
	g.Expect(*ret).
		To(Equal(testVal))
}

func TestGlobalSetAndGet(t *testing.T) {
	g := NewGomegaWithT(t)

	testKey, testVal := "testKey", "testVal"
	c.Config.Set(testKey, testVal)

	ret := c.Config.Get(testKey)
	g.Expect(*ret).
		To(Equal(testVal))
}
