// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var createModelCmd = &cobra.Command{
	Use:   "create-model",
	Short: "Create model commad",
	Long:  `This subcommand create model`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create model::" + args[0])
		if len(args[0]) == 0 {
			fmt.Println("Please provide a filename")
			return
		}

		filename := "models/" + args[0] + ".go"

		var _, err = os.Stat(filename)

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
	},
}

func init() {
	RootCmd.AddCommand(createModelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
