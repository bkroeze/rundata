package runedata

import (
	"fmt"
	"strings"
)

// move me to stringutils

func GetMaxLengthsOfStrings(data [][]string) []int {
	fieldLens := make([]int, len(data[0]))
	for i := 0; i < len(fieldLens); i++ {
		fieldLens[i] = 0
	}

	for recordIx := 0; recordIx < len(data); recordIx++ {
		record := data[recordIx]
		for fieldIx := 0; fieldIx < len(record); fieldIx++ {
			thisLen := len(record[fieldIx])
			if fieldLens[fieldIx] < thisLen {
				fieldLens[fieldIx] = thisLen
			}
		}
	}
	return fieldLens

}

func makeIntFormat(needed int, fill string) string {
	template := "%" + (fmt.Sprintf("%v", needed)) + "s"
	return fmt.Sprintf(template, fill)
}

func DataToLine(data []string, stringLengths []int) string {

	var out string

	maxLen := len(data)

	spaces := make([]string, maxLen)

	for i := 0; i < maxLen; i++ {
		spaces[i] = makeIntFormat(stringLengths[i]-len(data[i]), " ")
	}

	for i := 0; i < (maxLen - 1); i++ {
		out += data[i] + spaces[i] + "| "
	}
	out += data[maxLen-1]
	return out
}

func RunesToMDTable(runes []Rune) string {
	data := make([][]string, len(runes))
	header := make([]string, 3)
	header[0] = "Rune"
	header[1] = "Name"
	header[2] = "Quick Notes"

	for i := 0; i < len(runes); i++ {
		rune := runes[i]
		data[i] = []string{rune.Name, rune.Name, rune.Traditional}
	}

	stringLengths := GetMaxLengthsOfStrings(data)
	if stringLengths[2] < len(header[2]) {
		stringLengths[2] = len(header[2])
	}
	stringLengths[0] += 10 // add for %magick-RUNE;
	stringLengths[1] += 14 // add for [RUNE](RUNE.html)
	stringLengths[2] += 1  // add for [RUNE](RUNE.html)

	out := "\n" + DataToLine(header, stringLengths) +
		"\n" + strings.Replace(makeIntFormat(stringLengths[0]-1, "-"), " ", "-", -1) + " | " +
		strings.Replace(makeIntFormat(stringLengths[1]-1, "-"), " ", "-", -1) + " | " +
		strings.Replace(makeIntFormat(stringLengths[2]-1, "-"), " ", "-", -1)

	line := make([]string, 3)

	for i := 0; i < len(runes); i++ {
		line[0] = fmt.Sprintf("&magick-%s;", runes[i].Name)
		line[1] = fmt.Sprintf("[%s](%s.html)", strings.Title(runes[i].Name), runes[i].Name)
		line[2] = runes[i].Traditional
		out += "\n" + DataToLine(line, stringLengths)
	}

	out += "\n"
	return out
}
