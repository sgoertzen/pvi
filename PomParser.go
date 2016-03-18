package pvi

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// Projects is a list of type Project
type Projects []*Project

// Project represent a single project with links to parent and child projects
type Project struct {
	Parent                *Project `json:"-"`
	Children              Projects
	ArtifactID            string
	GroupID               string
	Version               string
	MismatchParentVersion string
	FullPath              string
	MissingParent         string
}

// PomProjects a list of type PomProject
type PomProjects []*PomProject

// PomProject represent a pom file
type PomProject struct {
	XMLName    xml.Name      `xml:"project"`
	Parent     PomParent     `xml:"parent"`
	GroupID    PomGroupID    `xml:"groupId"`
	ArtifactID PomArtifactID `xml:"artifactId"`
	Version    PomVersion    `xml:"version"`
	FullPath   string
}

// PomParent contains information on this projects parent
type PomParent struct {
	GroupID    PomGroupID    `xml:"groupId"`
	ArtifactID PomArtifactID `xml:"artifactId"`
	Version    PomVersion    `xml:"version"`
}

// PomGroupID is the group to which this project belongs
type PomGroupID struct {
	Value string `xml:",chardata"`
}

// PomArtifactID the id of the given pom file
type PomArtifactID struct {
	Value string `xml:",chardata"`
}

// PomVersion is the version of this project
type PomVersion struct {
	Value string `xml:",chardata"`
}

// GetProjects get all projects by reading the given directory
func GetProjects(projectPath string, debug bool) Projects {
	files := getDirectories(projectPath)
	pomProjects := PomProjects{}

	if debug {
		log.Printf("Found %d files/directories under %s", len(files), projectPath)
	}

	// Loop over each one
	for _, directory := range files {

		if !directory.IsDir() {
			if debug {
				log.Printf("Skipping %s as it is not a directory", directory.Name())
			}
			continue
		}
		pomFile := path.Join(projectPath, directory.Name(), "pom.xml")

		// Check for a pom.xml
		if _, err := os.Stat(pomFile); os.IsNotExist(err) {
			if debug {
				log.Printf("Unable to find pom file at %s", pomFile)
			}
			continue
		}

		pomProject, err := parseFile(pomFile)
		if err != nil || len(pomProject.ArtifactID.Value) == 0 {
			log.Println("Invalid pom file at: " + pomFile)
		}
		pomProject.FullPath = pomFile

		pomProjects = append(pomProjects, pomProject)

		if debug {
			log.Printf("Successfully read in project %s from %s", pomProject.ArtifactID, pomFile)
		}
	}

	projects := transform(pomProjects, debug)
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

func parseFile(pomFile string) (*PomProject, error) {
	v := new(PomProject)
	xmlFile, err := os.Open(pomFile)
	if err != nil {
		return v, err
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)

	err = xml.Unmarshal(b, v)
	return v, err
}

type dualProject struct {
	Pom  *PomProject
	Proj *Project
}

func linkParentChild(parentDual *dualProject, childDual *dualProject) {
	// Add to parent and point parent to us
	childDual.Proj.Parent = parentDual.Proj
	parentDual.Proj.Children = append(parentDual.Proj.Children, childDual.Proj)
	// Does parent version match what we need
	if childDual.Pom.Parent.Version.Value != parentDual.Pom.Version.Value {
		childDual.Proj.MismatchParentVersion = childDual.Pom.Parent.Version.Value
	}
}

func transform(pomProjects PomProjects, debug bool) Projects {
	if debug {
		log.Printf("Transforming %d projects", len(pomProjects))
	}
	parentProjects := Projects{}
	childrenMissingParents := PomProjects{}
	allProjects := make(map[string]*dualProject)

	// Loop over all projects
	for _, pomProject := range pomProjects {
		project := processProject(pomProject, debug)
		dual := dualProject{Pom: pomProject, Proj: project}
		allProjects[project.ArtifactID] = &dual

		if pomProject.Parent.ArtifactID.Value == "" {
			parentProjects = append(parentProjects, project)
		} else {
			childrenMissingParents = append(childrenMissingParents, pomProject)
		}
	}

	// Loop over children missing parents
	for _, child := range childrenMissingParents {
		childDual := allProjects[child.ArtifactID.Value]
		parentDual := allProjects[child.Parent.ArtifactID.Value]
		if parentDual == nil {
			// Parent not found, just add to parentProjects
			childDual.Proj.MissingParent = child.Parent.ArtifactID.Value
			parentProjects = append(parentProjects, childDual.Proj)

		} else {
			linkParentChild(parentDual, childDual)
		}
	}

	// for _, project := range allProjects {
	//     // If it has no parent add it to the parent projects
	//     if pomProject.Parent.ArtifactID.Value == "" {
	//         parentProjects = append(parentProjects, &project)
	//     } else {
	//         // If it has a parent look up the parent in the all map
	//         parent := allProjects[pomProject.Parent.ArtifactID.Value]

	//         // Update the pointer to our parent
	//         project.Parent = parent
	//         // Add ourselves to the parents children list
	//         parent.Children = append(parent.Children, &project)
	//         // Does parent version match what we need
	//         if pomProject.Parent.Version.Value != parent.Version {
	//             project.MismatchParentVersion = pomProject.Parent.Version.Value
	//         }
	//     }
	//     if debug {
	//         if project.Parent == nil {
	//             log.Printf("%s added with no parent", project.ArtifactID)
	//         } else {
	//             log.Printf("%s added with parent %s and mismatch version of %s", project.ArtifactID, project.Parent.ArtifactID, project.MismatchParentVersion)
	//         }
	//     }
	// }
	// loop again, checking if the parent is in the hashmap
	// for all items whos parent is not in hashmap or whos parent is null, process immediately
	//

	// Loop until we don't process any projects in a single run
	// for remaining != len(pomProjects)-len(allProjects) {
	// 	remaining = len(pomProjects) - len(allProjects)
	// 	// Loop over each project
	// 	for _, pomProject := range pomProjects {
	// 		if allProjects[pomProject.ArtifactID.Value] != nil {
	//             if debug {
	//                 log.Printf("Skipping %s as it has already been processed", pomProject.ArtifactID.Value)
	//             }
	// 			continue
	// 		}
	// 		if pomProject.Parent.ArtifactID.Value != "" && allProjects[pomProject.Parent.ArtifactID.Value] == nil {
	//             if debug {
	//                 log.Printf("Skipping %s as the parent project has not been processed yet", pomProject.ArtifactID.Value)
	//             }
	// 			continue
	// 		}

	//         // No matter what add it to the all projects map

	// 	}

	// }
	return parentProjects
}

func processProject(pomProject *PomProject, debug bool) *Project {
	// Build up our linked project
	project := Project{}
	project.ArtifactID = pomProject.ArtifactID.Value
	project.GroupID = pomProject.GroupID.Value
	project.Version = pomProject.Version.Value
	project.FullPath = pomProject.FullPath

	return &project
}

func (slice Projects) find(artifactID string) Project {
	for _, project := range slice {
		if project.ArtifactID == artifactID {
			return *project
		}
	}
	return Project{}
}

func (slice Projects) Len() int {
	return len(slice)
}

func (slice Projects) Less(i, j int) bool {
	return strings.ToLower(slice[i].ArtifactID) < strings.ToLower(slice[j].ArtifactID)
}

func (slice Projects) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
