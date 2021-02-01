package cmd

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var exists = func(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func createFile(filePath, content string) (err error) {
	var f *os.File
	if f, err = os.Create(filePath); err != nil {
		return
	}

	defer func() { _ = f.Close() }()

	_, err = f.WriteString(content)

	return
}

type input struct {
	param       string
	path        string
	errorLine   string
	successLine string
}
