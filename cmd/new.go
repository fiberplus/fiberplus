package cmd

import (
	"fmt"
	"regexp"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

var (
	templateType string
	repo         string
)

// NewCmd represents the hello command
var NewCmd = &cobra.Command{
	Use:   "generate::project",
	Short: "Generate new project",
	Long:  `project create succefully`,
	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Println("create model" + args[0])

		var c config
		c.getConfig()

		input := input{
			param:       args[0],
			path:        c.ModelPath,
			errorLine:   "Model already exists! ",
			successLine: "File created successfully ",
		}

		if len(args[0]) == 0 {
			fmt.Println("Please provide a valid command")
			return
		}

		err := createComplex("./"+strcase.ToLowerCamel(input.param), strcase.ToLowerCamel(input.param))

		if err != nil {
			fmt.Println("Error", err)
			return
		}

		fmt.Println(input.successLine, input.param)
		return

	},
}

func init() {
	RootCmd.AddCommand(NewCmd)
}

const githubPrefix = "https://github.com/"
const defaultRepo = "fiberplus/boilerplate"

var fullPathRegex = regexp.MustCompile(`^(http|https|git)`)

func createComplex(projectPath, modName string) (err error) {
	var git string
	if git, err = execLookPath("git"); err != nil {
		return
	}

	toClone := githubPrefix + defaultRepo
	// if isFullPath := fullPathRegex.MatchString(repo); isFullPath {
	// 	toClone = repo
	// }

	if err = runCmd(execCommand(git, "clone", toClone, projectPath)); err != nil {
		return
	}

	// if repo == defaultRepo {
	// 	if err = replace(projectPath, "go.mod", "boilerplate", modName); err != nil {
	// 		return
	// 	}

	// 	if err = replace(projectPath, "*.go", "boilerplate", modName); err != nil {
	// 		return
	// 	}
	// }
	return
}
