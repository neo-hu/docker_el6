// +build !exclude_graphdriver_devicemapper,!static_build,linux

package register 

import (
	// register the devmapper graphdriver
	_ "github.com/neo-hu/docker_el6/docker/daemon/graphdriver/devmapper"
)
