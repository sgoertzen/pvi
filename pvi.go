package main

import (
    "log"
    "os"
)

// Program to read in poms and determine
func main() {
    // Read path from parameters
	if len(os.Args) < 1 {
		log.Fatal("You must pass in the directory to scan")
	}
    path := os.Args[1]
    /*if len(path) == 0 {
        log.Fatal("You must pass in the directory to scan")
	}*/
    projects := GetProjects(path)
    generateReport(projects)
}


func generateReport(projects []*Project){
	// Sort by parent
    //sort.Sort(projects)
    for i, p := range projects {
        log.Println(i, p.ArtifactId)
    }
}

