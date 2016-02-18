package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func UpdateParentVersion(project Project) {
	if len(project.MismatchParentVersion) == 0 {
		return
	}
	SetParentVersionInPom(project.FullPath, project.MismatchParentVersion, project.Parent.Version)
}

func SetParentVersionInPom(pomPath string, currentVersion string, newVersion string) {
	input, err := ioutil.ReadFile(pomPath)

	output := SetParentVersionInText(string(input), currentVersion, newVersion)
	err = ioutil.WriteFile(pomPath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func SetParentVersionInText(pomText string, currentVersion string, newVersion string) string {

	lines := strings.Split(pomText, "\n")
	var parentFound = false
	for index, line := range lines {
		if parentFound {
			if strings.Contains(line, currentVersion) {
				lines[index] = strings.Replace(line, currentVersion, newVersion, 1)
				break
			}
		}
		if !parentFound || strings.Contains(line, "<parent>") {
			parentFound = true
		}
	}
	output := strings.Join(lines, "\n")
	return output
}
