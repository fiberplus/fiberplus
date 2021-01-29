package cmd

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type input struct {
	param       string
	path        string
	errorLine   string
	successLine string
}
