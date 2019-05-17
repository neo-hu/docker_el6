package cnmallocator

import (
	"github.com/neo-hu/docker_el6/libnetwork/drivers/overlay/ovmanager"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/remote"
	"github.com/neo-hu/docker_el6/docker/swarmkit/manager/allocator/networkallocator"
)

var initializers = []initializer{
	{remote.Init, "remote"},
	{ovmanager.Init, "overlay"},
}

// PredefinedNetworks returns the list of predefined network structures
func PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return nil
}
