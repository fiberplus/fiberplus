package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var createRepositoryCmd = &cobra.Command{
	Use:   "create::repository",
	Short: "Create repository commad",
	Long:  `This subcommand create Repository`,
	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Println("create model" + args[0])
		input := input{
			param:       args[0],
			path:        "pkg",
			errorLine:   "model already exists! ",
			successLine: "File created successfully ",
		}

		if len(args[0]) == 0 {
			fmt.Println("Please provide a valid command")
			return
		}

		errx := os.Mkdir(input.path+"/"+strings.ToLower(input.param), 0755)
		if errx != nil {
			log.Fatal(errx)
		}

		// TODO add basic model boilerplates
		filename := input.path + "/" + input.param + "service.go"

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
	RootCmd.AddCommand(createRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
