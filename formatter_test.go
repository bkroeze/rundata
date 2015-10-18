package runedata_test

import (
	"strings"

	. "github.com/bkroeze/runedata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formatter", func() {
	fehu := Rune{Name: "fehu", Traditional: "Wealth"}
	uruz := Rune{Name: "uruz", Traditional: "Aurochs"}

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

	Context("Rune details parsing", func() {
		It("Should insert details for a named rune", func() {
			runes := make([]Rune, 2)
			runes[0] = fehu
			runes[1] = uruz

			expected := "# &magick-fehu; Fehu"
			formatted := RuneToMD("fehu", runes)
			Expect(formatted).To(BeEquivalentTo(expected))
		})
	})

	/*
	   for future test
	   	template := `<!-- rune:details:ac -->
	   test
	   <!-- /rune:details:ac -->`
	*/

})
