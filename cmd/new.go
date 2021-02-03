package cmd

import (
	"fmt"
	"regexp"

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
	Long:  `This subcommand create new fiber plus project`,
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

		// // create model IF not exist
		// modelname := input.path + "/" + strcase.ToLowerCamel(input.param) + ".go"
		// modelExists := exists(modelname)

		// if modelExists == false {

		// 	var x string

		// 	x = "package " + input.path + "\n\n" +
		// 		"type " + input.param + " struct {\n " +
		// 		"\n\n}"

		// 	createFile(modelname, x)

		// }

		createComplex("./", input.param)

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
