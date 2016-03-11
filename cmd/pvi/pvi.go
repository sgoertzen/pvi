package main

import (
	"github.com/sgoertzen/pvi"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type config struct {
	path     *string
	format   *string
	filename *string
	nocolor  *bool
	version  *bool
}

// Program to read in poms and determine
func main() {
	config := getConfiguration()
	projects := pvi.GetProjects(*config.path)
	validate(projects, *config.path)
	outputResults(projects, *config.format, *config.filename, *config.nocolor)
}

func getConfiguration() config {
	config := config{}
	config.path = kingpin.Arg("path", "The `directory` that contains subfolders with maven projects.  Defaults to current directory.  Example: '/user/code/projects/'").Default(".").String()
	config.format = kingpin.Flag("format", "Specify the output format.  Should be either 'text' or 'json'").Default("text").Short('o').String()
	config.filename = kingpin.Flag("filename", "The file in which the output should be stored.  If this is left off the output will be printed to the console").Short('f').String()
	config.nocolor = kingpin.Flag("nocolor", "Do not color the output.  Ignored if filename is specified.").Default("false").Short('n').Bool()
	kingpin.Version("1.0.0")
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.CommandLine.HelpFlag.Short('?')

	kingpin.Parse()

	*config.path, _ = filepath.Abs(*config.path)
	return config
}

func validate(projects pvi.Projects, path string) {
	if projects.Len() > 0 {
		log.Printf("No project directories found under %s", path)
		os.Exit(0)
	}
}

func outputResults(projects pvi.Projects, format string, filename string, noColor bool) {
	sort.Sort(projects)

	// Turn color off if we are printing to a file
	noColor = filename != "" || noColor

	var output string
	if strings.EqualFold(format, "TEXT") {
		output = pvi.AsText(projects, noColor)
	} else {
		output = pvi.AsJson(projects)
	}

	if filename != "" {
		pvi.PrintToFile(output, filename)
	} else {
		pvi.PrintToTerminal(output)
	}
}
