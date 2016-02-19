package main

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

func (project *Project) build(runIT bool) (int, error) {
	app, err := exec.LookPath("mvn")
	var cmd *exec.Cmd
	if runIT {
		cmd = exec.Command(app, "clean", "install", "-Pintegration-tests")
	} else {
		cmd = exec.Command(app, "clean", "install")
	}

	path := strings.Replace(project.FullPath, "pom.xml", "", 1)
	cmd.Dir = path
	stdout, err := cmd.StdoutPipe()
	check(err)

	err = cmd.Start()
	check(err)

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		//log.Printf(in.Text())
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		return 1, err
	}
	return 0, nil
}
