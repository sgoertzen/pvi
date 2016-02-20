// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = "c78c87ec4d8f647d818c"
var testProjectId = "56a00afa2c2fab00290002ae"

func TestGetAllProjectsFromVersionEye(t *testing.T) {
	SetKey(key)
	veProjects := getAllProjectsFromVersionEye()
	assert.True(t, len(veProjects) > 0)
}

func TestGetProjectDetails(t *testing.T) {
	SetKey(key)
	veProject := getProjectDetailsFromVersionEye(testProjectId)
	assert.Equal(t, testProjectId, veProject.Id)
}
