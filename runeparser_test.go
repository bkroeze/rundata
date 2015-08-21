package runedata_test

import (
	. "hillsorcerer.com/runedata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runeparser", func() {
	Context("Basic parse tests", func() {
		It("Should parse a simple line", func() {
			record := "fehu,f,feoh,\"cattle, wealth\",\"luck, wealth, starting out, creative spark\""
			rune, err := RecordToRune(record)
			Expect(rune).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
