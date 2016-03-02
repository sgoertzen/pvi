package main

import (
	"flag"
	"sort"
	"strings"
	//"github.com/sgoertzen/veye"
	"github.com/sgoertzen/pvi"
	"net/http"
	"io"
)

// Program to read in poms and determine
func main() {
	var path = flag.String("path", ".", "The `directory` that contains subfolders with maven projects.  Example: '/user/code/projects/'")
	var format = flag.String("format", "text", "Specify the output format.  Should be either `'text' or 'json'`")
	var filename = flag.String("filename", "", "The file in which the output should be stored.  If this is left off the output will be printed to the console")
	var noColor = flag.Bool("nocolor", false, "Do not color the output.  Ignored if filename is specified.")

	flag.Parse()

	projects := pvi.GetProjects(*path)
	outputResults(projects, *format, *filename, *noColor)

	//veye.SetKey("something")
	//runServer()
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func runServer () {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
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
