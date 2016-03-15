package pvi

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
	"sort"
	"strings"
)

// AsJSON returns the projects in a minimized JSON format.
func (projects Projects) AsJSON() string {
	b, err := json.Marshal(projects)
	check(err)
	return string(b)
}

// AsText will return the projects in a readable test format.
func (projects Projects) AsText(noColor bool) string {
	var buffer bytes.Buffer
	for _, p := range projects {
		printProject(p, 0, &buffer, noColor)
	}
	return buffer.String()
}

func printProject(project *Project, depth int, buffer *bytes.Buffer, noColor bool) {
	color.NoColor = noColor

	buffer.WriteString(strings.Repeat("    ", depth))
	buffer.WriteString(color.GreenString(project.ArtifactID))
	buffer.WriteString(" (")
	buffer.WriteString(project.Version)
	buffer.WriteString(")")

	if project.MismatchParentVersion != "" {
		buffer.WriteString(color.YellowString(" Warning: looking for parent version: "))
		buffer.WriteString(project.MismatchParentVersion)
	}
	buffer.WriteString("\n")
	sort.Sort(project.Children)
	for _, child := range project.Children {
		printProject(child, depth+1, buffer, noColor)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
