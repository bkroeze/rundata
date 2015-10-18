package runedata

import (
	"fmt"
	"strings"

	"github.com/bkroeze/go.utils"
)

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
		data[i] = []string{"&magick-" + rune.Name + ";", "[" + rune.Name + "](" + rune.Name + ".html)", rune.Traditional}
	}

	stringLengths := utils.GetMaxLengthsOfStrings(data)
	if stringLengths[2] < len(header[2]) {
		stringLengths[2] = len(header[2])
	}
	//stringLengths[0] += 10                       // add for %magick-RUNE;
	//stringLengths[1] += (stringLengths[1]*2 + 1) // add for [RUNE](RUNE.html)
	//stringLengths[2] += 1

	stringLengths[0] += 1
	stringLengths[1] += 1

	out := "\n" + DataToLine(header, stringLengths) +
		"\n" + strings.Replace(makeIntFormat(stringLengths[0]-1, "-"), " ", "-", -1) + " | " +
		strings.Replace(makeIntFormat(stringLengths[1]-1, "-"), " ", "-", -1) + " | " +
		strings.Replace(makeIntFormat(stringLengths[2], "-"), " ", "-", -1)

	line := make([]string, 3)

	stringLengths[0] -= 1
	stringLengths[1] -= 1

	for i := 0; i < len(runes); i++ {
		line[0] = fmt.Sprintf("&magick-%s;", runes[i].Name)
		line[1] = fmt.Sprintf("[%s](%s.html)", strings.Title(runes[i].Name), runes[i].Name)
		line[2] = runes[i].Traditional
		out += "\n" + DataToLine(line, stringLengths)
	}

	out += "\n"
	return out
}
