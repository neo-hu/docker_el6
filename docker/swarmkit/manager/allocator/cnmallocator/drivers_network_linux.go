package cnmallocator

import (
	"github.com/neo-hu/docker_el6/libnetwork/drivers/bridge/brmanager"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/host"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/ipvlan/ivmanager"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/macvlan/mvmanager"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/overlay/ovmanager"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/remote"
	"github.com/neo-hu/docker_el6/docker/swarmkit/manager/allocator/networkallocator"
)

var initializers = []initializer{
	{remote.Init, "remote"},
	{ovmanager.Init, "overlay"},
	{mvmanager.Init, "macvlan"},
	{brmanager.Init, "bridge"},
	{ivmanager.Init, "ipvlan"},
	{host.Init, "host"},
}

// PredefinedNetworks returns the list of predefined network structures
func PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return []networkallocator.PredefinedNetworkData{
		{Name: "bridge", Driver: "bridge"},
		{Name: "host", Driver: "host"},
	}
}
