package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	homeDir string

	execLookPath = exec.LookPath
	execCommand  = exec.Command
	osExit       = os.Exit

	skipSpinner bool
)

func init() {
	homeDir, _ = os.UserHomeDir()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var exists = func(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func createFile(filePath, content string) (err error) {
	var f *os.File
	if f, err = os.Create(filePath); err != nil {
		return
	}

	defer func() { _ = f.Close() }()

	_, err = f.WriteString(content)

	return
}

func createDirectory(path string) {
	_ = os.Mkdir(path, 0700)
	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	os.Mkdir(path, 755)
	// }
}

type input struct {
	param       string
	path        string
	errorLine   string
	successLine string
}

type config struct {
	ModelPath      string `yaml:"modelpath"`
	PkgPath        string `yaml:"pkgpath"`
	ControllerPath string `yaml:"controllerpath"`
}

func (c *config) getConfig() *config {

	yamlFile, err := ioutil.ReadFile(".fiberplus.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// if yaml not exist use below values

	if len(c.ModelPath) < 1 {
		c.ModelPath = "models"
	}
	if len(c.PkgPath) < 1 {
		c.PkgPath = "pkg"
	}
	if len(c.ControllerPath) < 1 {
		c.ControllerPath = "controllers"
	}

	createDirectory(c.ModelPath)
	createDirectory(c.ControllerPath)
	createDirectory(c.PkgPath)

	return c
}

func runCmd(cmd *exec.Cmd) (err error) {

	var (
		stderr io.ReadCloser
		stdout io.ReadCloser
	)

	if stderr, err = cmd.StderrPipe(); err != nil {
		return
	}
	defer func() {
		_ = stderr.Close()
	}()
	go func() { _, _ = io.Copy(os.Stderr, stderr) }()

	if stdout, err = cmd.StdoutPipe(); err != nil {
		return
	}
	defer func() {
		_ = stdout.Close()
	}()
	go func() { _, _ = io.Copy(os.Stdout, stdout) }()

	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("failed to run %s", cmd.String())
	}

	return
}

// replaces matching file patterns in a path, including subdirectories
func replace(path, pattern, old, new string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return replaceWalkFn(path, info, pattern, []byte(old), []byte(new))
	})
}

func replaceWalkFn(path string, info os.FileInfo, pattern string, old, new []byte) (err error) {
	var matched bool
	if matched, err = filepath.Match(pattern, info.Name()); err != nil {
		return
	}

	if matched {
		var oldContent []byte
		if oldContent, err = ioutil.ReadFile(filepath.Clean(path)); err != nil {
			return
		}

		if err = ioutil.WriteFile(path, bytes.Replace(oldContent, old, new, -1), 0); err != nil {
			return
		}
	}

	return
}
