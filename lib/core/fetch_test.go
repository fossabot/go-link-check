package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchLinksFromPage(t *testing.T) {
	t.Skip("TestFetchLinksFromPage")
}

func TestMakeURLAbsolute(t *testing.T) {

	baseURL := "https://example.com"
	firstPath := "/something/other"
	secondPath := "other/something"
	fullPath := "https://example.com/abcdef"

	assert.Equal(t, baseURL+firstPath, MakeURLAbsolute(firstPath, baseURL))
	assert.Equal(t, baseURL+"/"+secondPath, MakeURLAbsolute(secondPath, baseURL))
	assert.Equal(t, fullPath, MakeURLAbsolute(fullPath, baseURL))
}
