package xml

import (
	"encoding/xml"
	"os"
)

// VocationList defines a server list of vocations
type VocationList struct {
	XMLName   xml.Name   `xml:"vocations"`
	Vocations []Vocation `xml:"vocation"`
}

// Vocation defines a server voation
type Vocation struct {
	XMLName     xml.Name `xml:"vocation"`
	ID          int      `xml:"id,attr"`
	Name        string   `xml:"name,attr"`
	Description string   `xml:"description,attr"`
	GainCap     int      `xml:"gaincap,attr"`
	GainHealth  int      `xml:"gainhp,attr"`
	GainMana    int      `xml:"gainmana,attr"`
	BaseSpeed   int      `xml:"basespeed,attr"`
}

// VocationSkill defines a vocation skill value
type VocationSkill struct {
	ID         int     `xml:"id,attr"`
	Multiplier float64 `xml:"multiplier,attr"`
}

// LoadVocationList loads the server vocations.xml file
func LoadVocationList(path string) (*VocationList, error) {
	// Open monsters.xml file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create xml decoder
	list := VocationList{}
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&list); err != nil {
		return nil, err
	}

	return &list, nil
}
