package xml

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

// ItemList defines the server item list
type ItemList struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

// Item defines a server item
type Item struct {
	XMLName    xml.Name        `xml:"item"`
	ID         int             `xml:"id,attr"`
	Name       string          `xml:"name,attr"`
	Article    string          `xml:"article,attr"`
	FromID     int             `xml:"fromid,attr"`
	ToID       int             `xml:"toid,attr"`
	Attributes []ItemAttribute `xml:"attribute"`
}

// ItemAttribute defines a server item attribute
type ItemAttribute struct {
	XMLName xml.Name `xml:"attribute"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

// LoadItemList loads the server item list
func LoadItemList(path string) (*ItemList, error) {
	// Open items.xml file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create xml decoder
	list := ItemList{}
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&list); err != nil {
		return nil, err
	}

	return &list, nil
}
