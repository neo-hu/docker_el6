package checkpoint

import "github.com/neo-hu/docker_el6/docker/api/types"

// Backend for Checkpoint
type Backend interface {
	CheckpointCreate(container string, config types.CheckpointCreateOptions) error
	CheckpointDelete(container string, config types.CheckpointDeleteOptions) error
	CheckpointList(container string, config types.CheckpointListOptions) ([]types.Checkpoint, error)
}
