package main

import "os/exec"

// TODO: turn PomParser into module
// TODO: import into here and main
// TODO: Pass project into this method
func BuildProject() {
	app := "mvn"

	args := "clean install verify -Pintegration-tests"

	cmd := exec.Command(app, args)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}
