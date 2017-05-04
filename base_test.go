package qtypes

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestNewBase(t *testing.T) {
	before := time.Now()
	b := NewBase("src1")
	after := time.Now()
	assert.Equal(t, version, b.BaseVersion)
	assert.Equal(t, "src1", b.SourcePath[0])
	assert.True(t, before.UnixNano() < b.Time.UnixNano())
	assert.True(t, after.UnixNano() > b.Time.UnixNano())
}


func TestNewTimedBase(t *testing.T) {
	now := time.Now()
	b := NewTimedBase("src1", now)
	assert.Equal(t, now, b.Time)
}

func TestBase_GetTimeUnix(t *testing.T) {
	now := time.Now()
	b := NewTimedBase("src1", now)
	assert.Equal(t, now.Unix(), b.GetTimeUnix())
}

func TestBase_GetTimeUnixNano(t *testing.T) {
	now := time.Now()
	b := NewTimedBase( "src1", now)
	assert.Equal(t, now.UnixNano(), b.GetTimeUnixNano())
}


func TestBase_AppendSrc(t *testing.T) {
	b := NewBase("src1")
	b.AppendSource("src2")
	assert.Equal(t, "src1", b.SourcePath[0])
	assert.Equal(t, "src2", b.SourcePath[1])
}
