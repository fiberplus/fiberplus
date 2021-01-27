package folder

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please give one argument.")
		return
	}
	path := arguments[1]

	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("Path does not exist!", err)
	}
}
