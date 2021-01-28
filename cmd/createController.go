package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var createControllerCmd = &cobra.Command{
	Use:   "create::controller",
	Short: "Create model commad",
	Long:  `This sub command create controllers`,
	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Println("create model" + args[0])
		input := input{
			param:       args[0],
			path:        "controllers",
			errorLine:   "Controller already exists! ",
			successLine: "File created successfully ",
		}

		if len(args[0]) == 0 {
			fmt.Println("Please provide a valid command")
			return
		}

		// TODO add basic model boilerplates
		filename := input.path + "/" + input.param + ".go"

		var _, err = os.Stat(filename)

		if os.IsNotExist(err) {
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
		} else {
			fmt.Println(input.errorLine, filename)
			return
		}

		fmt.Println(input.successLine, filename)
		return

	},
}

func init() {
	RootCmd.AddCommand(createControllerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
