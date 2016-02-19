package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)


func (project *Project) build(runIT bool) (int, error) {
	log.Println("Path: " + project.FullPath)
	app, err := exec.LookPath("mvn")
	var cmd *exec.Cmd
	if runIT {
		cmd = exec.Command(app, "clean", "install", "-Pintegration-tests")
	} else {
		cmd = exec.Command(app, "clean", "install")
	}

	path := strings.Replace(project.FullPath, "pom.xml", "", 1)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		return 1, err
	}
	return 0, nil
}