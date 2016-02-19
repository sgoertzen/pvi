// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllProjectsFromVersionEye(t *testing.T) {
	SetKey("c78c87ec4d8f647d818c")
	veProjects := getAllProjectsFromVersionEye()
	assert.True(t, len(veProjects) > 0)
}
