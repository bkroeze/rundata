package runedata

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Rune struct {
	Name        string
	Phonetic    string
	Anglo       string
	Traditional string
	Divination  string
}

func RecordToRune(record string) (Rune, error) {
	var rune Rune
	fields := strings.Split(record, ",")
	rune.Name = fields[0]
	rune.Phonetic = fields[1]
	rune.Anglo = fields[2]
	rune.Traditional = fields[3]
	rune.Divination = fields[4]
	return rune, nil
}

func main() {

	raw, err := ioutil.ReadFile("runes.csv")

	if err != nil {
		panic(err)
	}

	converted := string(raw[:])
	var lines []string = strings.Split(converted, "\n")

	var runes []Rune

	for ix := 0; ix < len(lines); ix++ {
		//fmt.Printf("%02d %s\n", ix, lines[ix])
		rune, err := RecordToRune(lines[ix])
		if err != nil {
			panic(err)
		}
		runes = append(runes, rune)
	}

	fmt.Printf("%+v\n", runes)
}
