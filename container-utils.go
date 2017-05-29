package qtypes

import (
	"github.com/docker/docker/api/types"
	"fmt"
)
// AssembleServiceSlot create {{.Service.Name}}.{{.Task.Slot}}
func AssembleServiceSlot(cnt *types.Container) string {
	sn, snok := cnt.Labels["com.docker.swarm.service.name"]
	ts, tsok := cnt.Labels["com.docker.swarm.task.slot"]
	if snok && tsok {
		return fmt.Sprintf("%s.%s", sn,ts)
	}
	return "<nil>"
}

