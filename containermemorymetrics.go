package qtypes


import (
	"strings"
	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
	"fmt"
	"math"
)

// Inspired by https://github.com/elastic/beats/blob/master/metricbeat/module/docker/cpu/helper.go
type MemoryStats struct {
	Base
	Container   *types.Container
	Failcnt   	uint64
	Limit     	uint64
	MaxUsage  	uint64
	TotalRss  	uint64
	TotalRssP 	float64
	Usage     	uint64
	UsageP 		float64
}

func NewMemoryStats(src Base, stats *dc.Stats) MemoryStats {
	return MemoryStats{
		Base:      src,
		Failcnt:   stats.MemoryStats.Failcnt,
		Limit:     stats.MemoryStats.Limit,
		MaxUsage:  stats.MemoryStats.MaxUsage,
		TotalRss:  stats.MemoryStats.Stats.TotalRss,
		TotalRssP: calcUsage(float64(stats.MemoryStats.Stats.TotalRss), float64(stats.MemoryStats.Limit)),
		Usage:     stats.MemoryStats.Usage,
		UsageP:    calcUsage(float64(stats.MemoryStats.Usage), float64(stats.MemoryStats.Limit)),
	}
}

func (ms *MemoryStats) ToMetrics(src string) []Metric {
	dim := map[string]string{
		"container_id": ms.Container.ID,
		"container_name": strings.Trim(ms.Container.Names[0], "/"),
		"image_name": ms.Container.Image,
		"command": strings.Replace(ms.Container.Command, " ", "#", -1),
		"created": fmt.Sprintf("%d", ms.Container.Created),
	}
	for k, v := range ms.Container.Labels {
		dim[k] = strings.Replace(v, " ", "#", -1)
	}
	return []Metric{
		ms.NewExtMetric(src, "memory.usage.percent", Gauge, ms.UsageP, dim, ms.Time, true),
		ms.NewExtMetric(src, "memory.total_rss.percent", Gauge, ms.TotalRssP, dim, ms.Time, true),
	}
}

func calcUsage(frac, all float64) float64 {
	v := float64(frac / all)
	if math.IsNaN(v) {
		v = 0.0
	}
	return v
}
