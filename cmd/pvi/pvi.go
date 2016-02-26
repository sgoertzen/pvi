package main

import (
	"flag"
	"sort"
	"strings"
	"github.com/sgoertzen/veye"
	"github.com/sgoertzen/pvi"
	"net/http"
	"io"
)

// Program to read in poms and determine
func main() {
	var path = flag.String("path", ".", "The `directory` that contains subfolders with maven projects.  Example: '/user/code/projects/'")
	var format = flag.String("format", "text", "Specify the output format.  Should be either `'text' or 'json'`")
	var filename = flag.String("filename", "", "The file in which the output should be stored.  If this is left off the output will be printed to the console")
	flag.Parse()

	projects := pvi.GetProjects(*path)
	outputResults(projects, *format, *filename)

	veye.SetKey("something")
	runServer()
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func runServer () {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func outputResults(projects pvi.Projects, format string, filename string) {
	sort.Sort(projects)

	var output string
	if strings.EqualFold(format, "TEXT") {
		output = pvi.AsText(projects)
	} else {
		output = pvi.AsJson(projects)
	}

	if filename != "" {
		pvi.PrintToFile(output, filename)
	} else {
		pvi.PrintToTerminal(output)
	}
}
