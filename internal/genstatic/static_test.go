package genstatic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pathInfo(t *testing.T) {
	snake, relativePath := pathInfo("/a/b/c/", "/a/b/c/d/e.go")
	assert.Equal(t, "DE", snake, "")
	assert.Equal(t, "d/e.go", relativePath, "")

	snake, relativePath = pathInfo("/a/", "/a/e.go")
	assert.Equal(t, "E", snake, "")
	assert.Equal(t, "e.go", relativePath, "")
}
