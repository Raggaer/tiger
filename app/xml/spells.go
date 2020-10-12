package xml

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/schollz/closestmatch"
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

// RuneSpellList defines the server rune spell list
type RuneSpellList struct {
	XMLName xml.Name     `xml:"spells"`
	Runes   []*RuneSpell `xml:"rune"`
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

// ConjureSpellList defines the server conjure spell list
type ConjureSpellList struct {
	XMLName  xml.Name        `xml:"spells"`
	Conjures []*ConjureSpell `xml:"conjure"`
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

// LoadConjuretSpells loads the server instant spell list
func LoadConjureSpells(path string) (*ConjureSpellList, error) {
	// Open spells.xml file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create xml decoder
	list := ConjureSpellList{}
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&list); err != nil {
		if s, ok := err.(*xml.SyntaxError); ok {
			return nil, fmt.Errorf("line %d, %s", s.Line, s.Msg)
		}
		return nil, err
	}

	return &list, nil
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
		if s, ok := err.(*xml.SyntaxError); ok {
			return nil, fmt.Errorf("line %d, %s", s.Line, s.Msg)
		}
		return nil, err
	}

	return &list, nil
}

// LoadRuneSpells loads the server rune spell list
func LoadRuneSpells(path string) (*RuneSpellList, error) {
	// Open spells.xml file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create xml decoder
	list := RuneSpellList{}
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&list); err != nil {
		if s, ok := err.(*xml.SyntaxError); ok {
			return nil, fmt.Errorf("line %d, %s", s.Line, s.Msg)
		}
		return nil, err
	}

	return &list, nil
}

// CreateFuzzyClosest creates a conjure spell list fuzzy closest search
func (c *InstantSpellList) CreateFuzzyClosest(s int) *closestmatch.ClosestMatch {
	list := []string{}
	for _, spell := range c.Spells {
		list = append(list, spell.Words)
	}
	cm := closestmatch.New(list, []int{s})
	return cm
}

// CreateFuzzyClosest creates a conjure spell list fuzzy closest search
func (c *RuneSpellList) CreateFuzzyClosest(s int) *closestmatch.ClosestMatch {
	list := []string{}
	for _, r := range c.Runes {
		list = append(list, r.Name)
	}
	cm := closestmatch.New(list, []int{s})
	return cm
}

// CreateFuzzyClosest creates a conjure spell list fuzzy closest search
func (c *ConjureSpellList) CreateFuzzyClosest(s int) *closestmatch.ClosestMatch {
	list := []string{}
	for _, conjure := range c.Conjures {
		list = append(list, conjure.Words)
	}
	cm := closestmatch.New(list, []int{s})
	return cm
}
