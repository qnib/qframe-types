package qtypes

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
)

type ContainerStats struct {
	Base
	Stats *docker.Stats
	Container docker.APIContainers
}

func NewContainerStats(src string, stats *docker.Stats, cnt docker.APIContainers) ContainerStats{
	return ContainerStats{
		Base: NewBase(src),
		Stats: stats,
		Container: cnt,
	}
}

type ContainerEvent struct {
	Event events.Message
	Container types.ContainerJSON
}

func (cs *ContainerStats) GetContainer() *types.Container {
	return &types.Container{
		ID: cs.Container.ID,
		Names: cs.Container.Names,
		Command: cs.Container.Command,
		Created: cs.Container.Created,
		Image: cs.Container.Image,
		Labels: cs.Container.Labels,
	}
}

// Flat out copied from https://github.com/elastic/beats
func (cs *ContainerStats) GetCpuStats() CPUStats {
	cnt := cs.GetContainer()
	// TODO: Use NewCPUStats?
	return CPUStats{
		Base: cs.Base,
		Container:                   cnt,
		PerCpuUsage:                 perCpuUsage(cs.Stats),
		TotalUsage:                  totalUsage(cs.Stats),
		UsageInKernelmode:           cs.Stats.CPUStats.CPUUsage.UsageInKernelmode,
		UsageInKernelmodePercentage: usageInKernelmode(cs.Stats),
		UsageInUsermode:             cs.Stats.CPUStats.CPUUsage.UsageInUsermode,
		UsageInUsermodePercentage:   usageInUsermode(cs.Stats),
		SystemUsage:                 cs.Stats.CPUStats.SystemCPUUsage,
		SystemUsagePercentage:       systemUsage(cs.Stats),
	}
}

func (cs *ContainerStats) GetMemStats() MemoryStats {
	cnt := cs.GetContainer()
	// TODO: Use NewMemoryStats?
	return MemoryStats{
		Base: cs.Base,
		Container: cnt,
		Failcnt:   cs.Stats.MemoryStats.Failcnt,
		Limit:     cs.Stats.MemoryStats.Limit,
		MaxUsage:  cs.Stats.MemoryStats.MaxUsage,
		TotalRss:  cs.Stats.MemoryStats.Stats.TotalRss,
		TotalRssP: float64(cs.Stats.MemoryStats.Stats.TotalRss) / float64(cs.Stats.MemoryStats.Limit),
		Usage:     cs.Stats.MemoryStats.Usage,
		UsageP:    float64(cs.Stats.MemoryStats.Usage) / float64(cs.Stats.MemoryStats.Limit),
	}
}