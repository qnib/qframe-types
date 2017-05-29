package qtypes

import (
	"github.com/docker/docker/api/types"
	"fmt"
	"strings"
)
// AssembleServiceSlot create {{.Service.Name}}.{{.Task.Slot}}
func AssembleServiceSlot(cnt *types.Container) string {
	if tn, tnok := cnt.Labels["com.docker.swarm.task.name"]; tnok {
		arr := strings.Split(tn, ".")
		if len(arr) != 3 {
			return "<nil>"
		}
		return fmt.Sprintf("%s.%s", arr[0], arr[1])
	}
	return "<nil>"
}

