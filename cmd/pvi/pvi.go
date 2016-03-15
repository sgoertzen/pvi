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
	verbose  *bool
}

// Program to read in poms and determine
func main() {
	config := getConfiguration()
	projects := pvi.GetProjects(*config.path, *config.verbose)
	validate(projects, config)
	outputResults(projects, config)
}

func getConfiguration() config {
	config := config{}
	config.path = kingpin.Arg("path", "The `directory` that contains subfolders with maven projects.  Defaults to current directory.  Example: '/user/code/projects/'").Default(".").String()
	config.format = kingpin.Flag("format", "Specify the output format.  Should be either 'text' or 'json'").Default("text").Short('o').String()
	config.filename = kingpin.Flag("filename", "The file in which the output should be stored.  If this is left off the output will be printed to the console").Short('f').String()
	config.nocolor = kingpin.Flag("nocolor", "Do not color the output.  Ignored if filename is specified.").Default("false").Short('n').Bool()
	config.verbose = kingpin.Flag("verbose", "Output debug information during the run.").Default("false").Short('d').Bool()
	kingpin.Version("1.0.0")
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.CommandLine.HelpFlag.Short('?')

	kingpin.Parse()
    
    if *config.verbose {
        config.print()
    }

	*config.path, _ = filepath.Abs(*config.path)
	return config
}

func validate(projects pvi.Projects, c config) {
	if projects.Len() == 0 {
		log.Printf("No project directories found under %s", *c.path)
		os.Exit(0)
	}
}

func outputResults(projects pvi.Projects, c config) {    
	sort.Sort(projects)

	// Turn color off if we are printing to a file
	*c.nocolor = *c.filename != "" || *c.nocolor

	var output string
	if strings.EqualFold(*c.format, "TEXT") {
		output = projects.AsText(*c.nocolor)
	} else {
		output = projects.AsJSON()
	}

	if *c.filename != "" {
		pvi.PrintToFile(output, *c.filename)
	} else {
		pvi.PrintToTerminal(output)
	}
}

func (c config) print() {
    log.Printf("Running with filename: %s, format: %s, Nocolor: %t, Path: %s", *c.filename, *c.format, *c.nocolor, *c.path)
}