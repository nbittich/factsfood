package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data-fields.txt
var fields string

func main() {
	s := "type OpenFoodFactCSVEntry struct {\n"
	for _, field := range strings.Split(fields, "\n") {
		field = strings.ReplaceAll(strings.TrimSpace(field), "-", "")
		if len(field) == 0 {
			continue
		}
		temp := strings.Split(field, "_")
		for i := range temp {
			temp[i] = strings.Title(temp[i])
		}
		s += fmt.Sprintf("%s string `json:\"%s\" bson:\"%s\"`\n", strings.Join(temp, ""), field, field)
	}

	s += "//keep stuff that are not in the csv at the end\n"
	s += "ID string `bson:\"_id\" json:\"_id\"`\n"
	s += "}\n"
	fmt.Println(s)
}
