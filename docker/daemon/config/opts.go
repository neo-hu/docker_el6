package config

import (
	"github.com/neo-hu/docker_el6/docker/api/types/swarm"
	"github.com/neo-hu/docker_el6/docker/daemon/cluster/convert"
	"github.com/neo-hu/docker_el6/docker/swarmkit/api/genericresource"
)

// ParseGenericResources parses and validates the specified string as a list of GenericResource
func ParseGenericResources(value []string) ([]swarm.GenericResource, error) {
	if len(value) == 0 {
		return nil, nil
	}

	resources, err := genericresource.Parse(value)
	if err != nil {
		return nil, err
	}

	obj := convert.GenericResourcesFromGRPC(resources)
	return obj, nil
}
