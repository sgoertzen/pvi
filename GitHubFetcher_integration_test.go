// +build integration

package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
	err := CloneAllRepos("AKQASF")
	assert.Nil(t, err)
}
