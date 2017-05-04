package qtypes

import (
	"time"
)

const (
	version = "0.3.1"
)

type Base struct {
	BaseVersion string
	Time			time.Time
	SourceID		int
	SourcePath		[]string
	SourceSuccess 	bool
}

func NewBase(src string) Base {
	return NewTimedBase(src, time.Now())
}

func NewTimedBase(src string, t time.Time) Base {
	return Base {
		BaseVersion: version,
		Time: t,
		SourceID: 0,
		SourcePath: []string{src},
		SourceSuccess: true,
	}
}

func (b *Base) GetTimeRFC() string {
	return b.Time.Format("2006-01-02T15:04:05.999999-07:00")
}

func (b *Base) GetTimeUnix() int64 {
	return b.Time.Unix()
}

func (b *Base) GetTimeUnixNano() int64 {
	return b.Time.UnixNano()
}

func (b *Base) AppendSource(src string) {
	b.SourcePath = append(b.SourcePath, src)
}

func (b *Base) IsLastSource(src string) bool {
	return b.SourcePath[len(b.SourcePath)-1] == src
}

func (b *Base) InputsMatch(inputs []string) bool {
	for _, inp := range inputs {
		if b.IsLastSource(inp) {
			return true
		}

	}
	return false
}