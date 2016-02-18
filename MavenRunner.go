package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func (project *Project) build() (int, error) {
	// TODO: support -Pintegration-tests
	log.Println("Path: " + project.FullPath)
	app, err := exec.LookPath("mvn")
	cmd := exec.Command(app, "clean", "install")
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

//
//func (project *Project) build() (int, error) {
//	// TODO: support -Pintegration-tests
//	app, err := exec.LookPath("mvn")
//	//app := "/usr/local/bin"
//	args := "clean install"
//	//args := []string{"clean", "install"}
//	cmd := exec.Command(app, args)
//	cmd.Dir = project.FullPath
//	//cmd.Start()
//
//	//stdout, err := cmd.StdoutPipe()
//	_, err = cmd.StdoutPipe()
//	if err != nil {
//		return 0, err
//	}
//
//	// start the command after having set up the pipe
//	if err := cmd.Start(); err != nil {
//		log.Printf("error2: %s", err)
//		return 1, err
//	}
///*
//	// read command's stdout line by line
//	in := bufio.NewScanner(stdout)
//
//	for in.Scan() {
//		log.Printf(in.Text()) // write each line to your log, or anything you need
//	}
//	if err := in.Err(); err != nil {
//		log.Printf("error: %s", err)
//	}*/
//	return 0, nil
//}
