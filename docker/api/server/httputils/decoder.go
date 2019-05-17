package httputils 

import (
	"io"

	"github.com/neo-hu/docker_el6/docker/api/types/container"
	"github.com/neo-hu/docker_el6/docker/api/types/network"
)

// ContainerDecoder specifies how
// to translate an io.Reader into
// container configuration.
type ContainerDecoder interface {
	DecodeConfig(src io.Reader) (*container.Config, *container.HostConfig, *network.NetworkingConfig, error)
	DecodeHostConfig(src io.Reader) (*container.HostConfig, error)
}
