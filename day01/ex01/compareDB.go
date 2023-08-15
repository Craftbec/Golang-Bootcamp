package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Ingredient struct {
	Ingredient_name  string `xml:"itemname" json:"ingredient_name" diff: "itemname, identifier"`
	Ingredient_count string `xml:"itemcount" json:"ingredient_count" diff: "ingredient_count"`
	Ingredient_unit  string `xml:"itemunit" json:"ingredient_unit" diff: "ingredient_unit"`
}

type Cake struct {
	Name        string       `xml:"name" json:"name" diff: "name, identifier"`
	Time        string       `xml:"stovetime" json:"time" diff: "time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients" diff: "ingredients, identifier"`
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Recipes []Cake   `xml:"cake" json:"cake" diff: "name, identifier"`
}

type XML Recipes
type JSON Recipes

type DBReader interface {
	read(file []byte) (Recipes, error)
}

func (x *XML) read(file []byte) (Recipes, error) {

	err := xml.Unmarshal(file, x)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Recipes(*x), err
	}
	return Recipes(*x), err

}

func (j *JSON) read(file []byte) (Recipes, error) {
	err := json.Unmarshal(file, j)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return Recipes(*j), err
	}
	return Recipes(*j), err
}

func seal(a DBReader, file []byte) []byte {
	var result []byte
	fil, err := a.read(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	switch a.(type) {
	case *JSON:
		result, err = xml.MarshalIndent(fil, "", "    ")
		if err != nil {
			fmt.Printf("error: %v", err)
			return nil
		}
	case *XML:
		result, err = json.MarshalIndent(fil, "", "    ")
		if err != nil {
			fmt.Printf("error: %v", err)
			return nil
		}
	default:
		break

	}
	return result
}

func check_format(a string) int8 {
	if path.Ext(a) == ".xml" {
		return 1
	} else if path.Ext(a) == ".json" {
		return 2
	} else {
		return 0
	}
}

func chek_ing(arr1, arr2 []Ingredient, s1, s2 string) {
	for i := 0; i < len(arr1); i++ {
		check := true
		for j := 0; j < len(arr1); j++ {
			if arr1[i].Ingredient_name == arr2[j].Ingredient_name {
				check = false
				if arr1[i].Ingredient_count != arr2[j].Ingredient_count {
					if arr2[j].Ingredient_count == "" {
						fmt.Printf("REMOVED unit count \"%s\" for ingredient \"%s\" for cake  \"%s\"\n",
							arr1[i].Ingredient_count, arr1[i].Ingredient_name, s1)
					} else {
						fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
							arr1[i].Ingredient_name, s1, arr2[j].Ingredient_count, arr1[i].Ingredient_count)
					}

				}
				if arr1[i].Ingredient_unit != arr2[j].Ingredient_unit {
					if arr2[j].Ingredient_unit == "" {
						fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n",
							arr1[i].Ingredient_unit, arr1[i].Ingredient_name, s1)
					} else {
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
							arr1[i].Ingredient_name, s1, arr2[j].Ingredient_unit, arr1[i].Ingredient_unit)
					}
				}
			}
		}
		if check == true {
			fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\" \n", arr1[i].Ingredient_name, s1)

		}
	}

	for i := 0; i < len(arr2); i++ {
		check := true
		for j := 0; j < len(arr1); j++ {
			if arr2[i].Ingredient_name == arr1[j].Ingredient_name {
				check = false
			}
		}
		if check == true {
			fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\" \n", arr2[i].Ingredient_name, s2)
		}
	}
}

func compare(a XML, b JSON) {
	for i := 0; i < len(a.Recipes); i++ {
		check := true
		for j := 0; j < len(b.Recipes); j++ {
			if a.Recipes[i].Name == b.Recipes[j].Name {
				check = false
				if a.Recipes[i].Time != b.Recipes[j].Time {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
						a.Recipes[i].Name, b.Recipes[j].Time, a.Recipes[i].Time)
				}
				chek_ing(a.Recipes[i].Ingredients, b.Recipes[j].Ingredients, a.Recipes[i].Name, b.Recipes[j].Name)
			}
		}
		if check == true {
			fmt.Printf("REMOVED cake \"%s\"\n", a.Recipes[i].Name)
		}
	}
	for i := 0; i < len(b.Recipes); i++ {
		check := true
		for j := 0; j < len(a.Recipes); j++ {
			if b.Recipes[i].Name == a.Recipes[j].Name {
				check = false
			}
		}
		if check == true {
			fmt.Printf("ADDED cake \"%s\"\n", b.Recipes[i].Name)
		}
	}
}

func main() {
	xmlbd := new(XML)
	jsonbd := new(JSON)
	if len(os.Args) == 5 && os.Args[1] == "--old" && check_format(os.Args[2]) == 1 && os.Args[3] == "--new" && check_format(os.Args[4]) == 2 {
		filex, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println("No such file")
			os.Exit(5)

		}
		filej, err := ioutil.ReadFile(os.Args[4])
		if err != nil {
			fmt.Println("No such file")
			os.Exit(5)

		}
		_, err = DBReader.read(xmlbd, filex)
		if err != nil {
			os.Exit(4)
		}
		_, err = DBReader.read(jsonbd, filej)
		if err != nil {
			os.Exit(4)
		}
		compare(*xmlbd, *jsonbd)
	} else {
		fmt.Println("Incorrect application launch")
		os.Exit(2)
	}
}
