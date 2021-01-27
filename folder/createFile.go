package folder

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		return
	}
	filename := os.Args[1]

	var _, err = os.Stat(filename)

	err := os.Mkdir("./models", 0755)
	if err != nil {
		log.Fatal(err)
	}

	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	} else {
		fmt.Println("File already exists!", filename)
		return
	}

	fmt.Println("File created successfully", filename)
}
