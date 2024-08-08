package pkg

import "encoding/json"

type JsonReader struct {
}

func (j JsonReader) ReadDB(data []byte) (Recipes, error) {

	var recipes Recipes

	err := json.Unmarshal(data, &recipes)

	if err != nil {
		return recipes, err
	}
	return recipes, nil
}
