package services

import (
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"testing"
)

func TestContextFields(t *testing.T) {
	t.Run("defines expected context fields", func(t *testing.T) {
		fields := ContextFields()
		assert.Assert(t, is.Len(fields, 0))
	})
}
