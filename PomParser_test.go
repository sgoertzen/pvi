package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransform(t *testing.T) {

	pomProject := PomProject{ArtifactID: PomArtifactID{Value: "myartifact"}}

	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, pomProject)
	output := transform(pomProjects, false)

	assert.Equal(t, "myartifact", output[0].ArtifactID)
}

func TestTransformParentChild(t *testing.T) {

	parent := PomProject{ArtifactID: PomArtifactID{Value: "parent"}}
	child := PomProject{
		ArtifactID: PomArtifactID{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactID: PomArtifactID{
				Value: "parent",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, parent)
	pomProjects = append(pomProjects, child)
	output := transform(pomProjects, false)

	assert.Equal(t, "parent", output[0].ArtifactID)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "child", output[0].Children[0].ArtifactID)
	}
}

func TestTransformParentChildOutOfOrder(t *testing.T) {

	parent := PomProject{ArtifactID: PomArtifactID{Value: "parent"}}
	child := PomProject{
		ArtifactID: PomArtifactID{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactID: PomArtifactID{
				Value: "parent",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, child)
	pomProjects = append(pomProjects, parent)
	output := transform(pomProjects, false)

	assert.Equal(t, "parent", output[0].ArtifactID)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "child", output[0].Children[0].ArtifactID)
	}
}

func TestTransformParentMatchingVersion(t *testing.T) {

	parent := PomProject{ArtifactID: PomArtifactID{Value: "parent"}, Version: PomVersion{Value: "1.0"}}
	child := PomProject{
		ArtifactID: PomArtifactID{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactID: PomArtifactID{
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
	output := transform(pomProjects, false)

	assert.Equal(t, "parent", output[0].ArtifactID)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "", output[0].Children[0].MismatchParentVersion)
	}
}

func TestTransformParentWrongVersion(t *testing.T) {

	parent := PomProject{ArtifactID: PomArtifactID{Value: "parent"}, Version: PomVersion{Value: "2.0"}}
	child := PomProject{
		ArtifactID: PomArtifactID{
			Value: "child",
		},
		Parent: PomParent{
			ArtifactID: PomArtifactID{
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
	output := transform(pomProjects, false)

	assert.Equal(t, "parent", output[0].ArtifactID)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "1.0", output[0].Children[0].MismatchParentVersion)
	}
}

func TestParentVersionMatchDefault(t *testing.T) {

	parent := PomProject{ArtifactID: PomArtifactID{Value: "parent"}, Version: PomVersion{Value: "2.0"}}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, parent)
	output := transform(pomProjects, false)

	assert.Equal(t, "parent", output[0].ArtifactID)
	assert.Equal(t, "", output[0].MismatchParentVersion)
}

func TestFullPath(t *testing.T) {
	projects := GetProjects("./test-data", false)
	assert.NotEmpty(t, projects[0].FullPath)

}
