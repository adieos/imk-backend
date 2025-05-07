package utils

import (
	"fmt"
	"strings"
)

func IljFormat(id string, departemen string) string {
	departemenLowerCase := strings.ToLower(departemen)
	departemenSlug := strings.ReplaceAll(departemenLowerCase, " ", "-")
	result := fmt.Sprintln(id + "-departement-" + departemenSlug)
	return result
}
