package qtypes


import (
	"time"
	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
)

// Inspired by https://github.com/elastic/beats/blob/master/metricbeat/module/docker/cpu/helper.go
type MemoryStats struct {
	Time        time.Time
	Container   *types.Container
	Failcnt   	uint64
	Limit     	uint64
	MaxUsage  	uint64
	TotalRss  	uint64
	TotalRssP 	float64
	Usage     	uint64
	UsageP 		float64
}

func NewMemoryStats(stats *dc.Stats) MemoryStats {
	return MemoryStats{
		Time:      stats.Read,
		Failcnt:   stats.MemoryStats.Failcnt,
		Limit:     stats.MemoryStats.Limit,
		MaxUsage:  stats.MemoryStats.MaxUsage,
		TotalRss:  stats.MemoryStats.Stats.TotalRss,
		TotalRssP: float64(stats.MemoryStats.Stats.TotalRss) / float64(stats.MemoryStats.Limit),
		Usage:     stats.MemoryStats.Usage,
		UsageP:    float64(stats.MemoryStats.Usage) / float64(stats.MemoryStats.Limit),
	}
}

