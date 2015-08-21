package runedata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRunedata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runedata Suite")
}
