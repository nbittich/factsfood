package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed data-fields.txt
var fields string

func firstToLower(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}

func main() {
	s := "type OpenFoodFactCSVEntry struct {\n"
	for _, field := range strings.Split(fields, "\n") {
		field = strings.ReplaceAll(strings.TrimSpace(field), "-", "_")
		field = strings.ReplaceAll(field, "url", "URL")
		if len(field) == 0 {
			continue
		}
		temp := strings.Split(field, "_")
		for i := range temp {
			temp[i] = cases.Title(language.Und).String(temp[i])
		}
		if field == "code" {
			s += fmt.Sprintf("%s string `json:\"_id\" bson:\"_id\"`\n", strings.Join(temp, ""))
		} else {
			goFieldName := strings.Join(temp, "")
			jsonFieldName := firstToLower(goFieldName)
			f := fmt.Sprintf("\"%s,omitempty\"", jsonFieldName)
			t := "string"

			if strings.HasSuffix(field, "_t") ||
				strings.HasSuffix(field, "_100g") ||
				strings.HasSuffix(field, "_serving") ||
				strings.HasSuffix(field, "_quantity") ||
				strings.HasSuffix(field, "_n") {
				t = "*int"
			}

			if strings.HasSuffix(field, "_datetime") {
				t = "types.TimeISO8601"
			}

			if field == "completeness" {
				t = "*float64"
			}
			if field == "creator" ||
				field == "states" ||
				field == "no_nutrition_data" ||
				field == "last_modified_by" ||
				strings.HasPrefix(field, "packaging") ||
				strings.HasPrefix(field, "popularity") {
				f = "\"-\""
			}

			s += fmt.Sprintf("%s %s `json:%s bson:%s`\n", goFieldName, t, f, f)
		}
	}
	s += "}\n"
	fmt.Println(s)
}
