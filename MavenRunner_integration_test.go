// +build integration

package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessfulPom(t *testing.T) {
	projects := GetProjects("./test-data/", false)
	parent := projects.find("parent-test")
	assert.NotEmpty(t, parent.FullPath)

	value, err := parent.build(false, false)
	assert.Equal(t, 0, value)
	assert.Nil(t, err)
}

func TestFailingPom(t *testing.T) {
	projects := GetProjects("./test-data/", false)
	failProject := projects.find("failing-test")
	assert.NotEmpty(t, failProject.FullPath)

	value, err := failProject.build(false, false)
	assert.Equal(t, 1, value)
	assert.NotNil(t, err)
}
