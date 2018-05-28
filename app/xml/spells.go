package xml

import (
	"encoding/xml"
	"os"
)

// InstantSpellList defines the server instant spell list
type InstantSpellList struct {
	XMLName xml.Name        `xml:"spells"`
	Spells  []*InstantSpell `xml:"instant"`
}

// InstantSpell defines a server instant spell
type InstantSpell struct {
	XMLName    xml.Name   `xml:"instant"`
	Group      string     `xml:"group,attr"`
	SpellID    int        `xml:"spellid,attr"`
	Name       string     `xml:"name,attr"`
	Words      string     `xml:"words,attr"`
	Level      int        `xml:"lvl,attr"`
	Mana       int        `xml:"mana,attr"`
	Premium    bool       `xml:"prem,attr"`
	Range      int        `xml:"range,attr"`
	NeedTarget bool       `xml:"needtarget,attr"`
	NeedWeapon bool       `xml:"needweapon,attr"`
	CoolDown   int        `xml:"cooldown,attr"`
	Script     string     `xml:"script,attr"`
	Vocations  []Vocation `xml:"vocation"`
}

// RuneSpell defines a server rune spell
type RuneSpell struct {
	XMLName       xml.Name   `xml:"rune"`
	Group         string     `xml:"group,attr"`
	SpellID       int        `xml:"spellid,attr"`
	Name          string     `xml:"name,attr"`
	ID            int        `xml:"id,attr"`
	Level         int        `xml:"lvl,attr"`
	MagicLevel    int        `xml:"maglv,attr"`
	Aggresive     bool       `xml:"aggresive,attr"`
	GroupCooldown int        `xml:"groupcooldown,attr"`
	Vocations     []Vocation `xml:"vocation"`
}

// ConjureSpell defines a server conjure spell
type ConjureSpell struct {
	XMLName   xml.Name   `xml:"conjure"`
	Group     string     `xml:"group,attr"`
	SpellID   int        `xml:"spellid,attr"`
	Name      string     `xml:"name,attr"`
	Level     int        `xml:"lvl,attr"`
	Mana      int        `xml:"mana,attr"`
	Soul      int        `xml:"soul,attr"`
	Words     string     `xml:"words,attr"`
	CoolDown  int        `xml:"cooldown,attr"`
	Vocations []Vocation `xml:"vocation"`
}

// LoadInstantSpells loads the server instant spell list
func LoadInstantSpells(path string) (*InstantSpellList, error) {
	// Open spells.xml file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create xml decoder
	list := InstantSpellList{}
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&list); err != nil {
		return nil, err
	}

	return &list, nil
}
