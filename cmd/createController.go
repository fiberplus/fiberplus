package cmd

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var createControllerCmd = &cobra.Command{
	Use:   "create::controller",
	Short: "Create model commad",
	Long:  `This sub command create controllers`,
	Run: func(cmd *cobra.Command, args []string) {

		var c config
		c.getConfig()

		input := input{
			param:       args[0],
			path:        c.ControllerPath,
			errorLine:   "Controller already exists! ",
			successLine: "Controller created successfully ",
		}

		if len(args[0]) == 0 {
			fmt.Println("Please provide a valid command")
			return
		}

		// create model IF not exist
		controllerName := input.path + "/" + strcase.ToLowerCamel(input.param) + ".go"
		controllerExists := exists(controllerName)

		if controllerExists == false {

			var x string

			x = "package " + input.path + "\n\n" +
				"type " + input.param + " struct {\n " +
				"\n\n}"

			createFile(controllerName, x)

		}

		fmt.Println(input.successLine, input.param)
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
