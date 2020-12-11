package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoteTitle(t *testing.T) {
	n := &note{content: "Note Name"}
	assert.Equal(t, "Note Name", n.title())

	n = &note{content: "line1\nline2"}
	assert.Equal(t, "line1", n.title())
}
