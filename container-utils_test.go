package qtypes

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/docker/docker/api/types"
)

func TestAssembleServiceSlot(t *testing.T) {
	cnt := &types.Container{
		Labels: map[string]string{},
	}
	got := AssembleServiceSlot(cnt)
	assert.Equal(t, "<nil>", got)
	cnt.Labels = map[string]string{
		"com.docker.swarm.service.name": "service1",
		"com.docker.swarm.task.slot": "1",
	}
	got = AssembleServiceSlot(cnt)
	assert.Equal(t, "service1.1", got)

}
