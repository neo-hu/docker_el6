package register

import (
	// register the windows graph drivers
	_ "github.com/neo-hu/docker_el6/docker/daemon/graphdriver/lcow"
	_ "github.com/neo-hu/docker_el6/docker/daemon/graphdriver/windows"
)
