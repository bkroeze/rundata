package runedata_test

import (
	. "github.com/bkroeze/runedata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runeparser", func() {
	var fehu Rune = Rune{Name: "fehu",
		Phonetic:    "f",
		Anglo:       "feoh",
		Traditional: "cattle, wealth",
		Divination:  "luck, wealth, starting out, creative spark"}

	var uruz Rune = Rune{Name: "uruz",
		Phonetic:    "u",
		Anglo:       "ur",
		Traditional: "aurochs",
		Divination:  "Sap, growth, healing, qi, life energy, grounding, rooted center (reversed: blockage)"}

	Context("Basic parse tests", func() {
		It("Should return a Rune struct", func() {
			record := []string{"fehu", "f", "feoh", "cattle, wealth", "luck, wealth, starting out, creative spark"}
			Expect(RecordToRune(record)).Should(BeAssignableToTypeOf(Rune{}))
		})

		It("Should parse a simple line", func() {
			record := []string{"fehu", "f", "feoh", "cattle, wealth", "luck, wealth, starting out, creative spark"}
			Expect(RecordToRune(record)).Should(BeEquivalentTo(fehu))

		})
	})

	Context("Multiline parsing", func() {
		It("Should parse multiple lines", func() {
			lines := "fehu,f,feoh\nuruz,u,ur"
			var shortfehu Rune = Rune{Name: "fehu",
				Phonetic:    "f",
				Anglo:       "feoh",
				Traditional: "",
				Divination:  ""}

			var shorturuz Rune = Rune{Name: "uruz",
				Phonetic:    "u",
				Anglo:       "ur",
				Traditional: "",
				Divination:  ""}

			runes, err := RecordsToRunes(lines, false)
			Expect(err).To(BeNil())
			Expect(runes).To(HaveLen(2))
			Expect(runes[0]).Should(BeEquivalentTo(shortfehu))
			Expect(runes[1]).Should(BeEquivalentTo(shorturuz))

		})
	})

	Context("Parsing a file", func() {
		It("Should parse the test file", func() {
			runes, err := RunesFromFile("runes.csv", true)
			Expect(err).To(BeNil())
			Expect(runes[0]).To(BeEquivalentTo(fehu))
			Expect(runes[1]).To(BeEquivalentTo(uruz))
		})
	})
})
