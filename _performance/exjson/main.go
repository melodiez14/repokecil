package main

import "github.com/melodiez14/repokecil/_performance/exjson/templates"

func main() {
	a := &templates.SmallStruct{
		Age:       10,
		FirstName: "Risal",
		LastName:  "Falah",
	}
	println(a.JSON())
}
