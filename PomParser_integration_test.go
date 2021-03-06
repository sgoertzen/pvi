// +build integration

package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProjects(t *testing.T) {
	projects := GetProjects("./test-data/", false)
	parent := projects.find("parent-test")

	assert.Equal(t, "parent-test", parent.ArtifactID)
	assert.Equal(t, "3.1.4", parent.Version)
	assert.Equal(t, 1, len(parent.Children))
	if len(parent.Children) > 0 {
		assert.Equal(t, "child-test", parent.Children[0].ArtifactID)
		assert.Equal(t, "2.2", parent.Children[0].Version)
		assert.Equal(t, "3.1.1", parent.Children[0].MismatchParentVersion)
	}
}

func TestOrphan(t *testing.T) {
	projects := GetProjects("./test-data/", false)
	orphan := projects.find("orphan-test")
	assert.NotEmpty(t, orphan.ArtifactID)
}
