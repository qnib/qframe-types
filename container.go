package qtypes

import (
	"github.com/fsouza/go-dockerclient"
)

type ContainerStats struct {
	Stats *docker.Stats
	Container docker.APIContainers
}
