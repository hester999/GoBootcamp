package main

import (
	"day1/pkg"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type DBReader interface {
	Read(path string, recipes Recipes) (*Recipes, error)
}

type JsonReader struct {
}

type XmlReader struct{}

type Recipes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"time"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>ingredient"`
}

type Ingredient struct {
	IngredientName  string  `json:"ingredient_name" xml:"name"`
	IngredientCount string  `json:"ingredient_count" xml:"count"`
	IngredientUnit  *string `json:"ingredient_unit,omitempty" xml:"unit,omitempty"`
}

func (j *JsonReader) Read(path string, recipes Recipes) (*Recipes, error) {
	extension := filepath.Ext(path)

	if extension != ".json" {
		fmt.Println("Not a JSON file")
		os.Exit(1)
	}

	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data, err := io.ReadAll(file)

	if err != nil {
		return nil, fmt.Errorf("Error reading the file: %v", err)
	}

	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}
	return &recipes, nil
}

func (x *XmlReader) Read(path string, recipes Recipes) (*Recipes, error) {
	extension := filepath.Ext(path)

	if extension != ".xml" {
		fmt.Println("Not a JSON file")
		os.Exit(1)
	}

	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading the file: %v", err)
	}
	err = xml.Unmarshal(data, &recipes)

	if err != nil {
		return nil, err
	}
	return &recipes, nil
}

func main() {
	flag := pkg.Flag{}
	str := flag.ParseCommandLine()
	reader := &XmlReader{}
	recipes, _ := reader.Read(str, Recipes{})
	fmt.Println(recipes)

}
