package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestText(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactId:"testproj",Version:"1.0"})

	json := toText(projects);

	assert.Equal(t, "testproj(1.0)\n", json)
}

func TestJson(t *testing.T) {
	projects := Projects{}
	projects = append(projects, &Project{ArtifactId:"testproj"})

	json := toJson(projects);

	assert.Equal(t, "[{\"Children\":null,\"ArtifactId\":\"testproj\",\"GroupId\":\"\",\"Version\":\"\",\"MismatchParentVersion\":\"\",\"FullPath\":\"\"}]", json)
}

func TestJsonWithChild(t *testing.T) {
	parent := Project{ArtifactId:"parent"}
	child := Project{ArtifactId:"child",Parent:&parent}
	parent.Children = append(parent.Children, &child)

	projects := Projects{}
	projects = append(projects, &parent)

	json := toJson(projects)

	assert.NotNil(t, json)
}

