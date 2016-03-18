package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestText(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactID: "testproj", Version: "1.0"})

	text := projects.AsText(true)

	assert.Equal(t, "testproj (1.0)\n", text)
}
func TestTextColor(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactID: "testproj", Version: "1.0"})

	text := projects.AsText(false)

	assert.Equal(t, "\x1b[32mtestproj\x1b[0m (1.0)\n", text)
}

func TestJson(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactID: "testproj"})

	json := projects.AsJSON()

	assert.Equal(t, "[{\"Children\":null,\"ArtifactID\":\"testproj\",\"GroupID\":\"\",\"Version\":\"\",\"MismatchParentVersion\":\"\",\"FullPath\":\"\",\"MissingParent\":\"\"}]", json)
}

func TestJsonWithChild(t *testing.T) {
	parent := Project{ArtifactID: "parent"}
	child := Project{ArtifactID: "child", Parent: &parent}
	parent.Children = append(parent.Children, &child)

	projects := Projects{}
	projects = append(projects, &parent)

	json := projects.AsJSON()

	assert.NotNil(t, json)
}
