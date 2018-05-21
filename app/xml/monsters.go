package xml

import (
	"encoding/xml"
	"os"
)

// MonsterList defines the monsters.xml file
type MonsterList struct {
	XMLName  xml.Name             `xml:"monsters"`
	Monsters []monsterListElement `xml:"monster"`
}

type monsterListElement struct {
	XMLName xml.Name `xml:"monster"`
	Name    string   `xml:"name,attr"`
	File    string   `xml:"file,attr"`
}

// LoadMonsterList loads the monsters.xml file
func LoadMonsterList(path string) (*MonsterList, error) {
	// Open monsters.xml file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create xml decoder
	list := MonsterList{}
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&list); err != nil {
		return nil, err
	}

	return &list, nil
}
