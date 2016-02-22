package main

import (
	"flag"
	"sort"
	"strings"
	"github.com/sgoertzen/veye"
)

// Program to read in poms and determine
func main() {
	var path = flag.String("path", ".", "The `directory` that contains subfolders with maven projects.  Example: '/user/code/projects/'")
	var format = flag.String("format", "text", "Specify the output format.  Should be either `'text' or 'json'`")
	var filename = flag.String("filename", "", "The file in which the output should be stored.  If this is left off the output will be printed to the console")
	flag.Parse()

	projects := GetProjects(*path)
	outputResults(projects, *format, *filename)

	veye.SetKey("something")
}

func outputResults(projects Projects, format string, filename string) {
	sort.Sort(projects)

	var output string
	if strings.EqualFold(format, "TEXT") {
		output = toText(projects)
	} else {
		output = toJson(projects)
	}

	if filename != "" {
		printToFile(output, filename)
	} else {
		printToTerminal(output)
	}
}
