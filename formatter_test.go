package runedata_test

import (
	"strings"

	. "hillsorcerer.com/runedata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formatter", func() {
	fehu := Rune{Name: "fehu", Traditional: "Wealth"}
	uruz := Rune{Name: "uruz", Traditional: "Aurochs"}

	Context("Utility", func() {
		It("Should get the max lengths of strings", func() {
			data := make([][]string, 2)
			data[0] = []string{"one", "two"}
			data[1] = []string{"three", "four"}
			theLengths := GetMaxLengthsOfStrings(data)
			Expect(theLengths).To(BeEquivalentTo([]int{5, 4}))
		})
	})
	Context("Table Format", func() {
		It("Should return an MD formatted table for a single rune", func() {
			singleLineTable := `
Rune          | Name              | Quick Notes
------------- | ----------------- | -----------
&magick-fehu; | [Fehu](fehu.html) | Wealth
`
			runes := make([]Rune, 1)
			runes[0] = fehu
			table := RunesToMDTable(runes)
			Expect(table).To(BeEquivalentTo(singleLineTable))
		})

		It("Should return an MD formatted table for two runes", func() {
			twoLineTable := `
Rune          | Name              | Quick Notes
------------- | ----------------- | -----------
&magick-fehu; | [Fehu](fehu.html) | Wealth
&magick-uruz; | [Uruz](uruz.html) | Aurochs
`
			runes := make([]Rune, 2)
			runes[0] = fehu
			runes[1] = uruz
			table := RunesToMDTable(runes)
			Expect(table).To(BeEquivalentTo(twoLineTable))
		})

		It("Should format a wide line", func() {
			data := strings.Split("Rune Name QuickNotes", " ")
			sizes := []int{10, 10, 10}

			line := DataToLine(data, sizes)
			Expect(line).To(Equal("Rune      | Name      | QuickNotes"))
		})
	})

})
