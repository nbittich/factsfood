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
			f := fmt.Sprintf("\"%s,omitempty\"", field)
			t := "string"
			if strings.HasSuffix(field, "_t") ||
				strings.HasSuffix(field, "_100g") ||
				strings.HasSuffix(field, "_serving") ||
				strings.HasSuffix(field, "_n") {
				t = "int"
			}

			if strings.HasSuffix(field, "_datetime") {
				t = "types.TimeISO8601"
			}
			if field == "creator" {
				f = "\"-\""
			}
			s += fmt.Sprintf("%s %s `json:%s bson:%s`\n", strings.Join(temp, ""), t, f, f)
		}
	}
	s += "}\n"
	fmt.Println(s)
}
