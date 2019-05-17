package libnetwork

import "github.com/neo-hu/docker_el6/libnetwork/drivers/ipvlan"

func additionalDrivers() []initializer {
	return []initializer{
		{ipvlan.Init, "ipvlan"},
	}
}
