package main

import (
    "log"
    "os"
	"strings"
	"sort"
)

// Program to read in poms and determine
func main() {
    // Read path from parameters
	if len(os.Args) < 1 {
		log.Fatal("You must pass in the directory to scan")
	}
    path := os.Args[1]
    projects := GetProjects(path)
    generateReport(projects)
}


func generateReport(projects Projects){
	sort.Sort(projects)
    for _, p := range projects {
		printProject(p, 0)
    }
}

func printProject(project *Project, depth int) {
	var misMatchError string
	if project.MismatchParentVersion != "" {
		misMatchError = " ** Warning: looking for parent version: " + project.MismatchParentVersion
	}
	log.Printf("%s%s (%s)%s", strings.Repeat("--", depth), project.ArtifactId, project.Version, misMatchError)
	sort.Sort(project.Children)
	for _, child := range project.Children {
		printProject(child, depth+1)
	}
}
