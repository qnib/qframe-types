package qtypes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDimensions_String(t *testing.T) {
	d := NewDimensions()
	assert.Equal(t,"", d.String())
	d.Add("key1", "val1")
	assert.Equal(t,"key1=val1", d.String())
	d.Add("key2", "val2")
	assert.Equal(t,"key1=val1,key2=val2", d.String())
}
