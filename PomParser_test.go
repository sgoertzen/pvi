package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"log"
)


func TestGetProjects(t *testing.T) {
	projects := GetProjects("./test-data/")

	log.Printf("Length of projects is: %d", len(projects))
	parent := projects.find("parent-test")
	assert.NotEmpty(t, parent)

	assert.Equal(t, "parent-test", parent.ArtifactId)
	assert.Equal(t, "3.1.4", parent.Version)
	assert.Equal(t, 1, len(parent.Children))
	if len(parent.Children) > 0 {
		assert.Equal(t, "child-test", parent.Children[0].ArtifactId)
		assert.Equal(t, "2.2", parent.Children[0].Version)
		assert.Equal(t, "3.1.1", parent.Children[0].MismatchParentVersion)
	}
}

func TestTransform(t *testing.T) {

	pomProject := PomProject{ArtifactId: PomArtifactId{Value: "myartifact"}}

	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, pomProject)
	output := transform(pomProjects)

	assert.Equal(t, "myartifact", output[0].ArtifactId)
}

func TestTransformParentChild(t *testing.T) {

	parent := PomProject{ArtifactId: PomArtifactId{Value: "parent"}}
	child := PomProject{
		ArtifactId: PomArtifactId{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactId: PomArtifactId{
				Value: "parent",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, parent)
	pomProjects = append(pomProjects, child)
	output := transform(pomProjects)

	assert.Equal(t, "parent", output[0].ArtifactId)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "child", output[0].Children[0].ArtifactId)
	}
}

func TestTransformParentChildOutOfOrder(t *testing.T) {

	parent := PomProject{ArtifactId: PomArtifactId{Value: "parent"}}
	child := PomProject{
		ArtifactId: PomArtifactId{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactId: PomArtifactId{
				Value: "parent",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, child)
	pomProjects = append(pomProjects, parent)
	output := transform(pomProjects)

	assert.Equal(t, "parent", output[0].ArtifactId)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "child", output[0].Children[0].ArtifactId)
	}
}

func TestTransformParentMatchingVersion(t *testing.T) {

	parent := PomProject{ArtifactId: PomArtifactId{Value: "parent"}, Version: PomVersion{Value: "1.0"}}
	child := PomProject{
		ArtifactId: PomArtifactId{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactId: PomArtifactId{
				Value: "parent",
			},
			Version: PomVersion{
				Value: "1.0",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, child)
	pomProjects = append(pomProjects, parent)
	output := transform(pomProjects)

	assert.Equal(t, "parent", output[0].ArtifactId)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "", output[0].Children[0].MismatchParentVersion)
	}
}

func TestTransformParentWrongVersion(t *testing.T) {

	parent := PomProject{ArtifactId: PomArtifactId{Value: "parent"}, Version: PomVersion{Value: "2.0"}}
	child := PomProject{
		ArtifactId: PomArtifactId{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactId: PomArtifactId{
				Value: "parent",
			},
			Version: PomVersion{
				Value: "1.0",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, child)
	pomProjects = append(pomProjects, parent)
	output := transform(pomProjects)

	assert.Equal(t, "parent", output[0].ArtifactId)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "1.0", output[0].Children[0].MismatchParentVersion)
	}
}

func TestParentVersionMatchDefault(t *testing.T) {

	parent := PomProject{ArtifactId: PomArtifactId{Value: "parent"}, Version: PomVersion{Value: "2.0"}}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, parent)
	output := transform(pomProjects)

	assert.Equal(t, "parent", output[0].ArtifactId)
	assert.Equal(t, "", output[0].MismatchParentVersion)
}

func TestFullPath(t *testing.T) {
	projects := GetProjects("./test-data")
	assert.NotEmpty(t, projects[0].FullPath)

}
