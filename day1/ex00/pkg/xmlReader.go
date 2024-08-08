package pkg

import "encoding/xml"

type XmlReader struct{}

func (x XmlReader) ReadDB(data []byte) (Recipes, error) {
	var recipes Recipes

	err := xml.Unmarshal(data, &recipes)

	if err != nil {
		return recipes, err
	}
	return recipes, nil
}
