package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLink(t *testing.T) {

	pomProject := PomProject{ArtifactId:PomArtifactId{Value:"myartifact"}}

	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, pomProject)
	output := link(pomProjects)

	assert.Equal(t, "myartifact", output[0].ArtifactId)
}

func TestLinkParentChild(t *testing.T) {

	parent := PomProject{ArtifactId:PomArtifactId{Value:"parent"}}
	child := PomProject{
		ArtifactId:PomArtifactId{
			Value:"child",
		},
		Parent:PomParent{
			ArtifactId:PomArtifactId{
				Value:"parent",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, parent)
	pomProjects = append(pomProjects, child)
	output := link(pomProjects)

	assert.Equal(t, "parent", output[0].ArtifactId)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "child", output[0].Children[0].ArtifactId)
	}
}

func TestLinkParentChildOutOfOrder(t *testing.T) {

	parent := PomProject{ArtifactId:PomArtifactId{Value:"parent"}}
	child := PomProject{
		ArtifactId:PomArtifactId{
			Value:"child",
		},
		Parent:PomParent{
			ArtifactId:PomArtifactId{
				Value:"parent",
			},
		},
	}
	pomProjects := PomProjects{}
	pomProjects = append(pomProjects, child)
	pomProjects = append(pomProjects, parent)
	output := link(pomProjects)

	assert.Equal(t, "parent", output[0].ArtifactId)
	assert.Equal(t, 1, len(output[0].Children))
	if len(output[0].Children) > 0 {
		assert.Equal(t, "child", output[0].Children[0].ArtifactId)
	}
}
