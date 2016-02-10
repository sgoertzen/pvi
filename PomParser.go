package main

import (
	"encoding/xml"
	"os"
	"log"
	"io/ioutil"
	"path"
)



/*
func (slice PomProjects) Len() int {
    return len(slice)
}

func (slice PomProjects) Less(i, j int) bool {
    // if neither has parent sort by name
    if slice[i].Parent.ArtifactId.Value == "" && slice[j].Parent.ArtifactId.Value == "" {
        return slice[i].ArtifactId.Value < slice[j].ArtifactId.Value
    }
    // If either parent is null sort that to the top
    if slice[i].Parent.ArtifactId.Value == "" {
        return true
    }
    if  slice[j].Parent.ArtifactId.Value == "" {
        return false
    }
    // If they have the same parent sort by name
    if slice[i].Parent.ArtifactId.Value == slice[j].Parent.ArtifactId.Value {
        return slice[i].ArtifactId.Value < slice[j].ArtifactId.Value
    }
    // Otherwise sort by parent
    return slice[i].Parent.ArtifactId.Value < slice[j].Parent.ArtifactId.Value;
}

func (slice PomProjects) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}*/


type Project struct {
	Parent *Project
	Children []*Project
	ArtifactId string
	GroupId string
	Version string

}

type PomProjects []PomProject

type PomProject struct {
	XMLName xml.Name `xml:"project"`
	Parent PomParent `xml:"parent"`
	GroupId PomGroupId `xml:"groupId"`
	ArtifactId PomArtifactId `xml:"artifactId"`
	Version PomVersion `xml:"version"`
}
type PomParent struct {
	GroupId PomGroupId `xml:"groupId"`
	ArtifactId PomArtifactId `xml:"artifactId"`
	Version PomVersion `xml:"version"`
}
type PomGroupId struct {
	Value  string `xml:",chardata"`
}
type PomArtifactId struct {
	Value  string `xml:",chardata"`
}
type PomVersion struct {
	Value  string `xml:",chardata"`
}


func GetProjects(path2 string) []*Project {
	files := getDirectories(path2)
	pomProjects := PomProjects{}

	// Loop over each one
	for _, directory := range files {
		if !directory.IsDir() {
			continue
		}

		pomFile := path.Join(path2, directory.Name(), "pom.xml")
		//log.Println("Looking at file: " + pomFile)

		// Check for a pom.xml
		if _, err := os.Stat(pomFile); os.IsNotExist(err) {
			//log.Println("Skipping directory " + directory.Name() + "  as it does not have a pom.xml file.")
			continue
		}
		//log.Println("Parsing pom file at " + pomFile)
		pomProject := parseFile(pomFile);

		pomProjects = append(pomProjects, pomProject)

		//log.Println("Found: " + project.ArtifactId.Value + ":" + project.Version.Value + " (Parent: " + project.Parent.ArtifactId.Value + ":" + project.Parent.Version.Value + ")")
	}

	projects := link(pomProjects)
	return projects
}

func getDirectories(path string) []os.FileInfo {
	// Get a list of directories off this
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Error reading the directory")
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


func link(pomProjects PomProjects) []*Project {
	parentProjects := []*Project{}

	var allProjects map[string]*Project
	allProjects = make(map[string]*Project)

	// Loop over each project
	for _, pomProject := range pomProjects {
		// Build up our linked project
		project := Project{}
		project.ArtifactId = pomProject.ArtifactId.Value
		project.GroupId = pomProject.GroupId.Value
		project.Version = pomProject.Version.Value

		// No matter what add it to the all projects map
		allProjects[project.ArtifactId] = &project

		// If it has no parent add it to the parent projects
		if (pomProject.Parent.ArtifactId.Value == "") {
			parentProjects = append(parentProjects, &project)
		} else {
			log.Println("Appending child to parent")
			// If it has a parent look up the parent in the all map
			parent := allProjects[pomProject.Parent.ArtifactId.Value]

			log.Println("Parent was found: " + parent.ArtifactId)

			// GRRRRRRRRRRRRR!!!!!!!!!!!!!!!!!!!!!!!!
			// This isn't going to work as it depends on the ordering of the projects!!
			// Can we sort this ahead of time into something with parents up front?
			// Need to think about this...
			// Could also just loop until all items are included or our the parent isn't found

			// Update the pointer to our parent
			project.Parent = parent
			// Add ourselves to the parents children list
			parent.Children = append(parent.Children, &project)
			log.Printf("Length of children %d",  len(parent.Children))
		}
	}
	return parentProjects
}