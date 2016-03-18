// +build integration

package pvi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPrintToFile(t *testing.T) {
	var filename = "fileToPrint.txt"
	defer os.Remove(filename)

	PrintToFile("This is written!", filename)
	_, err := os.Stat(filename)
	assert.False(t, os.IsNotExist(err))
}
