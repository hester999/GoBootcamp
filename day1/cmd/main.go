package main

import (
	"day1/pkg"
	"fmt"
)

type DBReader interface {
	Read(path string)
}

//func ReadDateBase(DbReader DBReader, path string) {
//	DBReader.Read()
//}

type Recipes struct {
	Cake []struct {
		Name        string `json:"name"`
		Time        string `json:"time"`
		Ingredients []struct {
			IngredientName  string  `json:"ingredient_name"`
			IngredientCount string  `json:"ingredient_count"`
			IngredientUnit  *string `json:"ingredient_unit,omitempty"`
		} `json:"ingredients"`
	} `json:"cake"`
}

func (j Recipes) Read(path string) {

}

func GetExpansion(path string) string {
	var res string
	for i, k := range path {
		if k == '.' {
			res = path[i+1:]
		}
	}

	return res
}

func main() {
	flag := pkg.Flag{}
	var JsonReader DBReader = RecipesJson{}
	path := flag.ParseCommandLine()
	JsonReader.Read(path)
	test := GetExpansion(path)
	fmt.Println(test)

}
