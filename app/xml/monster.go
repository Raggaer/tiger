package xml

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

// Monster defines a server cerature
type Monster struct {
	XMLName      xml.Name            `xml:"monster"`
	Name         string              `xml:"name,attr"`
	Description  string              `xml:"nameDescription,attr"`
	Race         string              `xml:"race,attr"`
	Experience   int                 `xml:"experience,attr"`
	Speed        int                 `xml:"speed,attr"`
	Health       MonsterHealth       `xml:"health"`
	Look         MonsterLook         `xml:"look"`
	TargetChange MonsterTargetChange `xml:"targetchange"`
	Attacks      MonsterAttackList   `xml:"attacks"`
	Defenses     MonsterDefenseList  `xml:"defenses"`
	Voices       MonsterVoiceList    `xml:"voices"`
	Loot         MonsterLootList     `xml:"loot"`
}

// MonsterHealth defines the monster health values
type MonsterHealth struct {
	XMLName xml.Name `xml:"health"`
	Now     int      `xml:"now,attr"`
	Max     int      `xml:"max,attr"`
}

// MonsterLook defines the monster looktype values
type MonsterLook struct {
	XMLName xml.Name `xml:"look"`
	Type    int      `xml:"type,attr"`
	Corpse  int      `xml:"corpse,attr"`
}

// MonsterTargetChange defines the monster targetting change values
type MonsterTargetChange struct {
	XMLName  xml.Name `xml:"targetchange"`
	Interval int      `xml:"interval,attr"`
	Chance   int      `xml:"chance,attr"`
}

// MonsterAttackList defines a list of monster attacks
type MonsterAttackList struct {
	Attacks []MonsterAttack `xml:"attack"`
}

// MonsterAttack defines a monster attack
type MonsterAttack struct {
	XMLName    xml.Name           `xml:"attack"`
	Name       string             `xml:"name,attr"`
	Interval   int                `xml:"interval,attr"`
	Range      int                `xml:"range,attr"`
	Min        int                `xml:"min,attr"`
	Max        int                `xml:"max,attr"`
	Target     int                `xml:"target,attr"`
	Attributes []MonsterAttribute `xml:"attribute"`
}

// MonsterAttribute defines a monster attribute
type MonsterAttribute struct {
	XMLName xml.Name `xml:"attribute"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

// MonsterDefenseList defines a monster defense list
type MonsterDefenseList struct {
	Armor    int              `xml:"armor,attr"`
	Defense  int              `xml:"defense,attr"`
	Defenses []MonsterDefense `xml:"defense"`
}

// MonsterDefense defines a monster defense value
type MonsterDefense struct {
	Name       string             `xml:"name,attr"`
	Interval   int                `xml:"interval,attr"`
	Chance     int                `xml:"chance,attr"`
	Min        int                `xml:"min,attr"`
	Max        int                `xml:"max,attr"`
	Attributes []MonsterAttribute `xml:"attribute"`
}

// MonsterVoiceList defines a list of monster voices
type MonsterVoiceList struct {
	Interval int            `xml:"interval,attr"`
	Chance   int            `xml:"chance,attr"`
	Voices   []MonsterVoice `xml:"voice"`
}

// MonsterVoice defines a monster sentence
type MonsterVoice struct {
	XMLName  xml.Name `xml:"voice"`
	Sentence string   `xml:"sentence,attr"`
}

// MonsterLootList defines a list of monster lootable items
type MonsterLootList struct {
	Loot []MonsterItem `xml:"item"`
}

// MonsterItem defines a monster lootable item
type MonsterItem struct {
	XMLName  xml.Name `xml:"item"`
	ID       int      `xml:"id,attr"`
	Name     string   `xml:"name,attr"`
	CountMax int      `xml:"countmax,attr"`
	Chance   int      `xml:"chance,attr"`
}

// LoadMonster loads the given monster xml file
func LoadMonster(path string) (*Monster, error) {
	// Open monster .xml file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create xml decoder
	monster := Monster{}
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&monster); err != nil {
		return nil, err
	}

	return &monster, nil
}
