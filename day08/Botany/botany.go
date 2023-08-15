package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(inter interface{}) {
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s", field.Name)
		if len(field.Tag) != 0 {
			tag := string(field.Tag)
			idx := strings.Index(string(field.Tag), ":")
			fmt.Printf("(%s=%s)", tag[:idx], field.Tag.Get(tag[:idx]))
		}
		fmt.Printf(":%v\n", v.FieldByName(field.Name))
	}
}

func main() {
	var plant = UnknownPlant{"rose", "jagged", 30}
	var plant2 = AnotherUnknownPlant{10, "lanceolate", 15}
	describePlant(plant)
	fmt.Println()
	describePlant(plant2)
}
