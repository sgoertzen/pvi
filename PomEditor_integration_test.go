// +build integration

package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParentEdit(t *testing.T) {
	// Cleanup
    defer SetParentVersionInPom("./test-data/child/pom.xml", "3.1.4", "3.1.1")

	projects := GetProjects("./test-data/", false)
	parent := projects.find("parent-test")
	assert.Equal(t, "parent-test", parent.ArtifactID)
	assert.NotEmpty(t, parent.Children[0].MismatchParentVersion)
	assert.Equal(t, "3.1.4", parent.Version)
	assert.Equal(t, "child-test", parent.Children[0].ArtifactID)
	assert.Equal(t, "2.2", parent.Children[0].Version)

	UpdateParentVersion(*parent.Children[0])

	projectsUpdated := GetProjects("./test-data/", false)
	assert.Empty(t, projectsUpdated.find("parent-test").Children[0].MismatchParentVersion)
}
