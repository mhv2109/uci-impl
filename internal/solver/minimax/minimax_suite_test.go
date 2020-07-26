package minimax_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinimax(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Minimax Suite")
}
