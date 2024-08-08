package pkg

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type DBReader interface {
	ReadDB(data []byte) (Recipes, error)
}

type Recipes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string  `json:"ingredient_name" xml:"itemname"`
	Count string  `json:"ingredient_count" xml:"itemcount"`
	Unit  *string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	data, _ := io.ReadAll(file)

	return data
}

func GetDB(path string) (Recipes, error) {
	if path == "" {
		return Recipes{}, errors.New("ERROR: need to use -f [filename]")
	}
	data := readFile(path)

	var dbReader DBReader

	extension := filepath.Ext(path)

	switch extension {
	case ".json":

		dbReader = &JsonReader{}
	case ".xml":
		dbReader = &XmlReader{}
	}

	recipes, err := dbReader.ReadDB(data)

	if err != nil {
		fmt.Println("Error parsing data:", err)
		return Recipes{}, err
	}

	return recipes, nil

}

func PrintDB(db Recipes, path string) {
	var content []byte
	var err error
	extension := filepath.Ext(path)
	if extension == ".json" {
		content, err = xml.MarshalIndent(db, "", "    ")
	} else if extension == ".xml" {
		content, err = json.MarshalIndent(db, "", "    ")
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n", content)
}
