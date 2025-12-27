package lib

import (
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInput(filePath string) string {
	data, err := os.ReadFile(filePath)
	Check(err)
	return string(data)
}

func filterRows(rows []string) []string {
	arr := make([]string, 0)
	for _, r := range rows {
		if len(r) > 0 {
			arr = append(arr, r)
		}
	}
	return arr
}

func GetRows(s string) []string {
	rows := strings.Split(s, "\n")
	rows = filterRows(rows)
	return rows
}
