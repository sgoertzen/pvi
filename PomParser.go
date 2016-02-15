package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type Projects []*Project

type Project struct {
	Parent                *Project `json:"-"`
	Children              Projects
	ArtifactId            string
	GroupId               string
	Version               string
	MismatchParentVersion string
	FullPath              string
}

type PomProjects []PomProject

type PomProject struct {
	XMLName    xml.Name      `xml:"project"`
	Parent     PomParent     `xml:"parent"`
	GroupId    PomGroupId    `xml:"groupId"`
	ArtifactId PomArtifactId `xml:"artifactId"`
	Version    PomVersion    `xml:"version"`
	FullPath   string
}
type PomParent struct {
	GroupId    PomGroupId    `xml:"groupId"`
	ArtifactId PomArtifactId `xml:"artifactId"`
	Version    PomVersion    `xml:"version"`
}
type PomGroupId struct {
	Value string `xml:",chardata"`
}
type PomArtifactId struct {
	Value string `xml:",chardata"`
}
type PomVersion struct {
	Value string `xml:",chardata"`
}

func GetProjects(path2 string) Projects {
	files := getDirectories(path2)
	pomProjects := PomProjects{}

	// Loop over each one
	for _, directory := range files {
		if !directory.IsDir() {
			continue
		}
		pomFile := path.Join(path2, directory.Name(), "pom.xml")

		// Check for a pom.xml
		if _, err := os.Stat(pomFile); os.IsNotExist(err) {
			continue
		}
		pomProject := parseFile(pomFile)
		pomProject.FullPath = pomFile

		pomProjects = append(pomProjects, pomProject)
	}

	projects := transform(pomProjects)
	return projects
}

func getDirectories(path string) []os.FileInfo {
	// Get a list of directories off this
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Error reading the directory: " + path)
	}
	return files
}

func parseFile(pomFile string) PomProject {
	v := new(PomProject)
	xmlFile, err := os.Open(pomFile)
	if err != nil {
		log.Println("Error opening file:", err)
		// TODO: Return an error here
		return *v
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(b, v)
	return *v
}

func transform(pomProjects PomProjects) Projects {
	//parentProjects := []*Project{}
	parentProjects := Projects{}

	var allProjects map[string]*Project
	allProjects = make(map[string]*Project)

	var remaining int
	remaining = 0
	for remaining != len(pomProjects)-len(allProjects) {
		remaining = len(pomProjects) - len(allProjects)
		// Loop over each project
		for _, pomProject := range pomProjects {
			if allProjects[pomProject.ArtifactId.Value] != nil {
				continue
			}

			if pomProject.Parent.ArtifactId.Value != "" && allProjects[pomProject.Parent.ArtifactId.Value] == nil {
				continue
			}

			// Build up our linked project
			project := Project{}
			project.ArtifactId = pomProject.ArtifactId.Value
			project.GroupId = pomProject.GroupId.Value
			project.Version = pomProject.Version.Value
			project.FullPath = pomProject.FullPath

			// No matter what add it to the all projects map
			allProjects[project.ArtifactId] = &project

			// If it has no parent add it to the parent projects
			if pomProject.Parent.ArtifactId.Value == "" {
				parentProjects = append(parentProjects, &project)
			} else {
				// If it has a parent look up the parent in the all map
				parent := allProjects[pomProject.Parent.ArtifactId.Value]

				// Update the pointer to our parent
				project.Parent = parent
				// Add ourselves to the parents children list
				parent.Children = append(parent.Children, &project)
				// Does parent version match what we need
				if pomProject.Parent.Version.Value != parent.Version {
					project.MismatchParentVersion = pomProject.Parent.Version.Value
				}
			}
		}

	}
	return parentProjects
}

func (slice Projects) Len() int {
	return len(slice)
}

func (slice Projects) Less(i, j int) bool {
	return strings.ToLower(slice[i].ArtifactId) < strings.ToLower(slice[j].ArtifactId)
}

func (slice Projects) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
