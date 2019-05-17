package daemon

import (
	"github.com/neo-hu/docker_el6/docker/api/types/container"
	"github.com/neo-hu/docker_el6/docker/libcontainerd"
)

func toContainerdResources(resources container.Resources) *libcontainerd.Resources {
	// We don't support update, so do nothing
	return nil
}
