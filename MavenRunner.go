package main

import (
	"os/exec"
	"bufio"
	"log"
)

// TODO: turn PomParser into module
// TODO: import into here and main
// TODO: Pass project into this method
func BuildProject(project Project) (int, error) {
	app := "mvn"
	args := "clean install verify -Pintegration-tests"



	cmd := exec.Command(app, args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		return 1, err
	}

	// read command's stdout line by line
	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Printf(in.Text()) // write each line to your log, or anything you need
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}
	return 0, nil
}
