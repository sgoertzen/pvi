package pvi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetParentInText(t *testing.T) {
	pomText := "<parent>\n<version>1.0</version>\n</parent>"
	output := SetParentVersionInText(pomText, "1.0", "2.0")
	assert.Contains(t, output, "2.0")
}
