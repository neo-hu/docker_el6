package daemon

import (
	"github.com/neo-hu/docker_el6/docker/container"
	"github.com/neo-hu/docker_el6/docker/libcontainerd"
)

// postRunProcessing perfoms any processing needed on the container after it has stopped.
func (daemon *Daemon) postRunProcessing(_ *container.Container, _ libcontainerd.EventInfo) error {
	return nil
}
