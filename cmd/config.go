package cmd

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// exists returns whether the given file or directory exists
// func exists(path string) (bool, error) {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return true, nil
// 	}
// 	if os.IsNotExist(err) {
// 		return false, nil
// 	}
// 	return false, err
// }

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

// type blueprint struct {
// 	service    string
// 	repository string
// }

// var BlueprintDefault = blueprint{
// 	service: "package " + input.path + "\n\n" +
// 		"type " + input.param + " struct {\n " +
// 		"\n\n}",

// 	repository: "package " + input.path + "\n\n" +
// 		"type " + input.param + " struct {\n " +
// 		"\n\n}",
// }
