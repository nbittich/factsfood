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
		field = strings.ReplaceAll(field, "url", "URL")
		if len(field) == 0 {
			continue
		}
		temp := strings.Split(field, "_")
		for i := range temp {
			temp[i] = strings.Title(temp[i])
		}
		if field == "code" {
			s += fmt.Sprintf("%s string `json:\"_id\" bson:\"_id\"`\n", strings.Join(temp, ""))
		} else {
			t := "string"
			if strings.HasSuffix(field, "_t") || strings.HasSuffix(field, "_100g") || strings.HasSuffix(field, "_serving") {
				t = "int"
			}
			s += fmt.Sprintf("%s %s `json:\"%s,omitempty\" bson:\"%s,omitempty\"`\n", strings.Join(temp, ""), t, field, field)
		}
	}
	s += "}\n"
	fmt.Println(s)
}
