package runedata

import (
	"io/ioutil"

	"hillsorcerer.com/utils"
)

type Rune struct {
	Name        string
	Phonetic    string
	Anglo       string
	Traditional string
	Divination  string
}

func RecordToRune(record []string) (Rune, error) {
	var rune Rune

	switch len(record) {
	case 5:
		rune.Divination = record[4]
		fallthrough
	case 4:
		rune.Traditional = record[3]
		fallthrough
	case 3:
		rune.Anglo = record[2]
		fallthrough
	case 2:
		rune.Phonetic = record[1]
		fallthrough
	case 1:
		rune.Name = record[0]
	}
	return rune, nil
}

func RecordsToRunes(raw string, skipFirst bool) ([]Rune, error) {

	records, err := utils.SplitMultilineCSV(raw, skipFirst)
	if err != nil {
		return nil, err
	}
	runes := make([]Rune, len(records))

	for ix := 0; ix < len(records); ix++ {
		//fmt.Printf("%02d %s\n", ix, lines[ix])
		rune, err := RecordToRune(records[ix])
		if err != nil {
			return runes, err
		}
		runes[ix] = rune
	}
	return runes, nil
}

func RunesFromFile(filename string) ([]Rune, error) {
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	converted := string(raw[:])

	return RecordsToRunes(converted, true)
}
