package cmd

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isFilePathExist() {

}

type input struct {
	param       string
	path        string
	errorLine   string
	successLine string
}
