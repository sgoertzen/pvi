// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessfulPom(t *testing.T) {
	projects := GetProjects("./test-data/")
	parent := projects.find("parent-test")
	assert.NotEmpty(t, parent.FullPath)

	value, err := parent.build(false)
	assert.Equal(t, 0, value)
	assert.Nil(t, err)
}

func TestFailingPom(t *testing.T) {
	projects := GetProjects("./test-data/")
	failProject := projects.find("failing-test")
	assert.NotEmpty(t, failProject.FullPath)

	value, err := failProject.build(false)
	assert.Equal(t, 1, value)
	assert.NotNil(t, err)
}
