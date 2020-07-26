package random_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("random", func() {

		RegisterFailHandler(Fail)
		RunSpecs(GinkgoT(), "Random Suite")
	})
})
