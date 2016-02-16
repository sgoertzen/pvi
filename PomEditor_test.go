package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetParentInText(t *testing.T) {
	pomText := "<parent>\n<version>1.0</version>\n</parent>"
	output := SetParentVersionInText(pomText, "1.0", "2.0")
	assert.Contains(t, output, "2.0")
}

func TestParentEdit(t *testing.T) {
	// Cleanup in case of failed previous test
	SetParentVersionInPom("./test-data/project2/pom.xml", "3.1.4", "3.1.1")

	projects := GetProjects("./test-data/")
	assert.Equal(t, "parent-test", projects[0].ArtifactId)
	assert.NotEmpty(t, projects[0].Children[0].MismatchParentVersion)
	assert.Equal(t, "3.1.4", projects[0].Version)
	assert.Equal(t, "child-test", projects[0].Children[0].ArtifactId)
	assert.Equal(t, "2.2", projects[0].Children[0].Version)

	UpdateParentVersion(*projects[0].Children[0])

	projectsUpdated := GetProjects("./test-data/")
	assert.Empty(t, projectsUpdated[0].Children[0].MismatchParentVersion)

	// Cleanup
	SetParentVersionInPom("./test-data/project2/pom.xml", "3.1.4", "3.1.1")
}