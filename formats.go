package main

import (
	"encoding/json"
	"bytes"
	"strings"
	"sort"
	"log"
	"os"
	"bufio"
)


type Formatter interface {
	format(projects Projects) string
}
type Printer interface {
	print(output string)
}

func toJson(projects Projects) string {
	b, err := json.Marshal(projects)
	check(err)
	return string(b)
}

func toText(projects Projects) string {
	var buffer bytes.Buffer
	for _, p := range projects {
		printProject(p, 0, &buffer)
	}
	return buffer.String()
}

func printProject(project *Project, depth int, buffer *bytes.Buffer) {
	buffer.WriteString(strings.Repeat("--", depth))
	buffer.WriteString(project.ArtifactId)
	buffer.WriteString("(")
	buffer.WriteString(project.Version)
	buffer.WriteString(")")

	if project.MismatchParentVersion != "" {
		buffer.WriteString(" ** Warning: looking for parent version: ")
		buffer.WriteString(project.MismatchParentVersion)
	}
	buffer.WriteString("\n")
	sort.Sort(project.Children)
	for _, child := range project.Children {
		printProject(child, depth+1, buffer)
	}
}

func printToTerminal(output string){
	log.Println(output)
}

func printToFile(output string) {
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(output)
	check(err)

	w.Flush()

	f.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}