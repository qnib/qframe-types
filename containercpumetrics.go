package qtypes

import (
	"time"
	"strconv"
	"strings"
	"github.com/elastic/beats/libbeat/common"
	"github.com/docker/docker/api/types"
	dc "github.com/fsouza/go-dockerclient"
)

// Inspired by https://github.com/elastic/beats/blob/master/metricbeat/module/docker/cpu/helper.go
type CPUStats struct {
	Time                        time.Time
	Container                   *types.Container
	PerCpuUsage                 common.MapStr
	TotalUsage                  float64
	UsageInKernelmode           uint64
	UsageInKernelmodePercentage float64
	UsageInUsermode             uint64
	UsageInUsermodePercentage   float64
	SystemUsage                 uint64
	SystemUsagePercentage       float64
}

func (cs *CPUStats) ToMetrics() []Metric {
	dim := map[string]string{
		"container_id": cs.Container.ID,
		"container_name": strings.Trim(cs.Container.Names[0], "/"),
		"image_name": cs.Container.Image,
		"command": cs.Container.Command,
		"created": string(cs.Container.Created),
	}
	return []Metric{
		NewExt("filter", "filter-dstat", "usage_kernel_percent", Gauge, cs.UsageInKernelmodePercentage, dim, cs.Time, false),
		NewExt("filter", "filter-dstat", "usage_user_percent", Gauge, cs.UsageInUsermodePercentage, dim, cs.Time, false),
		NewExt("filter", "filter-dstat", "system_usage_percent", Gauge, cs.SystemUsagePercentage, dim, cs.Time, false),
	}
}


func perCpuUsage(stats *dc.Stats) common.MapStr {
	var output common.MapStr
	if len(stats.CPUStats.CPUUsage.PercpuUsage) == len(stats.PreCPUStats.CPUUsage.PercpuUsage) {
		output = common.MapStr{}
		for index := range stats.CPUStats.CPUUsage.PercpuUsage {
			cpu := common.MapStr{}
			cpu["pct"] = calculateLoad(stats.CPUStats.CPUUsage.PercpuUsage[index] - stats.PreCPUStats.CPUUsage.PercpuUsage[index])
			cpu["ticks"] = stats.CPUStats.CPUUsage.PercpuUsage[index]
			output[strconv.Itoa(index)] = cpu
		}
	}
	return output
}

func totalUsage(stats *dc.Stats) float64 {
	return calculateLoad(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
}

func usageInKernelmode(stats *dc.Stats) float64 {
	return calculateLoad(stats.CPUStats.CPUUsage.UsageInKernelmode - stats.PreCPUStats.CPUUsage.UsageInKernelmode)
}

func usageInUsermode(stats *dc.Stats) float64 {
	return calculateLoad(stats.CPUStats.CPUUsage.UsageInUsermode - stats.PreCPUStats.CPUUsage.UsageInUsermode)
}

func systemUsage(stats *dc.Stats) float64 {
	return calculateLoad(stats.CPUStats.SystemCPUUsage - stats.PreCPUStats.SystemCPUUsage)
}

func calculateLoad(value uint64) float64 {
	return float64(value) / float64(1000000000)
}

// \beats

