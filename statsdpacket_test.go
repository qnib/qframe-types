package qtypes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewStatsdPacket(t *testing.T) {
	sp := NewStatsdPacket("testCounter", float64(1.2), "c")
	assert.Equal(t,"testCounter", sp.Bucket)
}

func TestStatsdPacket_String(t *testing.T) {

	sp := NewStatsdPacket("testCounter", float64(1.2), "c")
	assert.Equal(t,"", sp.DimensionString())
	sp.AddDimension("key1", "val1")
	assert.Equal(t,"key1=val1", sp.DimensionString())
	sp.AddDimension("key2", "val2")
	assert.Equal(t,"key1=val1,key2=val2", sp.DimensionString())
}