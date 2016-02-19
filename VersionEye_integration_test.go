// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllProjectsFromVersionEye(t *testing.T) {
	veProjects := getAllProjectsFromVersionEye()
	assert.True(t, len(veProjects) > 0)
}
