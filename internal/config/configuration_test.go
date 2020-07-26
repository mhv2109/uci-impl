package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	c "github.com/mhv2109/uci-impl/internal/config"
)

var _ = Describe("Config", func() {
	var _ = Describe("NewConfiguration", func() {
		var config c.Configuration

		BeforeEach(func() {
			config = c.NewConfiguration()
		})

		var _ = Describe("Setting and Getting key", func() {
			testKey, testVal := "testKey", "testVal"

			BeforeEach(func() {
				config.Set(testKey, testVal)
			})

			It("Get key", func() {
				ret := config.Get(testKey)
				Expect(*ret).
					To(Equal(testVal))
			})
		})
	})

	var _ = Describe("Global Configuration", func() {
		var _ = Describe("Setting and Getting key", func() {
			testKey, testVal := "testKey", "testVal"

			BeforeEach(func() {
				c.Config.Set(testKey, testVal)
			})

			It("Get key", func() {
				ret := c.Config.Get(testKey)
				Expect(*ret).
					To(Equal(testVal))
			})
		})
	})
})
