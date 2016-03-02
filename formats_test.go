package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestText(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactId: "testproj", Version: "1.0"})

	json := AsText(projects, true)

	assert.Equal(t, "testproj (1.0)\n", json)
}
func TestTextColor(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactId: "testproj", Version: "1.0"})

	json := AsText(projects, false)

	assert.Equal(t, "\x1b[32mtestproj\x1b[0m (1.0)\n", json)
}

func TestJson(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactId: "testproj"})

	json := AsJson(projects)

	assert.Equal(t, "[{\"Children\":null,\"ArtifactId\":\"testproj\",\"GroupId\":\"\",\"Version\":\"\",\"MismatchParentVersion\":\"\",\"FullPath\":\"\"}]", json)
}

func TestJsonWithChild(t *testing.T) {
	parent := Project{ArtifactId: "parent"}
	child := Project{ArtifactId: "child", Parent: &parent}
	parent.Children = append(parent.Children, &child)

	projects := Projects{}
	projects = append(projects, &parent)

	json := AsJson(projects)

	assert.NotNil(t, json)
}
