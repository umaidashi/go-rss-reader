package utils

import (
	"bytes"
	"fmt"
	"grr/static"
	"path/filepath"
)

const (
	AaDir = "aa"
)

func GetAAFromFp(fileName string) string {
	text := readFile(fileName)
	return text
}

func readFile(fileName string) string {
	filePath := filepath.Join(AaDir, fileName)
	file, err := static.Aa.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.String()
}
