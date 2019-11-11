package core

import (
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

func TestContextFields(t *testing.T) {
	fields := ContextFields()
	assert.Assert(t, is.Len(fields, 0))
}
