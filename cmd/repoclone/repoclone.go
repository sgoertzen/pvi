package main

import (
	"github.com/sgoertzen/pvi"
	"gopkg.in/alecthomas/kingpin.v2"
)

type config struct {
	organization  *string
}

// Clone all the repos of an orgnaization
func main() {
	config := getConfiguration()
	pvi.CloneAllRepos(*config.organization)
}

func getConfiguration() config {
	config := config{}
	config.organization = kingpin.Arg("organization", "GitHub organization that should be cloned").Required().String()
	kingpin.Version("1.0.0")
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.CommandLine.HelpFlag.Short('?')

	kingpin.Parse()
	return config
}
