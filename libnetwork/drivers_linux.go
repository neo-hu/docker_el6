package libnetwork

import (
	"github.com/neo-hu/docker_el6/libnetwork/drivers/bridge"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/host"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/macvlan"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/null"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/overlay"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/remote"
)

func getInitializers(experimental bool) []initializer {
	in := []initializer{
		{bridge.Init, "bridge"},
		{host.Init, "host"},
		{macvlan.Init, "macvlan"},
		{null.Init, "null"},
		{remote.Init, "remote"},
		{overlay.Init, "overlay"},
	}

	if experimental {
		in = append(in, additionalDrivers()...)
	}
	return in
}
