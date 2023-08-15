package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"

	"io/ioutil"
	"os"
	"path"
)

type Cake struct {
	Name        string       `xml:"name" json:"name"`
	Time        string       `xml:"stovetime" json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

type Ingredient struct {
	Ingredient_name  string `xml:"itemname" json:"ingredient_name"`
	Ingredient_count string `xml:"itemcount" json:"ingredient_count"`
	Ingredient_unit  string `xml:"itemunit" json:"ingredient_unit"`
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Recipes []Cake   `xml:"cake" json:"cake"`
}

type XML Recipes
type JSON Recipes

type DBReader interface {
	read(file []byte) (Recipes, error)
}

func (x *XML) read(file []byte) (Recipes, error) {

	err := xml.Unmarshal(file, x)
	if err != nil {
		fmt.Printf("error: %v", err)
		return Recipes(*x), err
	}
	return Recipes(*x), err

}

func (j *JSON) read(file []byte) (Recipes, error) {
	err := json.Unmarshal(file, j)
	if err != nil {
		fmt.Printf("error: %v", err)
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

func main() {
	var f bool
	var path_file string
	flag.BoolVar(&f, "f", false, "Point the way")
	flag.Parse()
	if f {
		if !(len(os.Args) == 3) {
			fmt.Println("No file")
			os.Exit(3)
		}
		path_file = os.Args[2]
		if len(path_file) < 5 {
			fmt.Println("Incorrect file")
			os.Exit(2)
		}
		file, err := ioutil.ReadFile(path_file)
		if err != nil {
			fmt.Println("No such file")
			os.Exit(5)

		}
		if path.Ext(path_file) == ".xml" {
			mystr := new(XML)
			fmt.Printf("%s\n", seal(mystr, file))
		} else if path.Ext(path_file) == ".json" {
			mystr := new(JSON)
			fmt.Printf("%s\n", seal(mystr, file))
		} else {
			fmt.Println("Incorrect file")
			os.Exit(2)
		}
	} else {
		fmt.Println("Flag -f must be used")
	}
}
