package qtypes

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
)

type ContainerStats struct {
	Stats *docker.Stats
	Container docker.APIContainers
}

type ContainerEvent struct {
	Event events.Message
	Container types.Container
}

// Flat out copied from https://github.com/elastic/beats
func (cs *ContainerStats) GetCpuStats() CPUStats {
	cnt := &types.Container{
		ID: cs.Container.ID,
		Names: cs.Container.Names,
		Command: cs.Container.Command,
		Created: cs.Container.Created,
		Image: cs.Container.Image,
	}
	return CPUStats{
		Time:                        cs.Stats.Read,
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
