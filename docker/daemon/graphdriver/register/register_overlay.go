// +build !exclude_graphdriver_overlay,linux

package register 

import (
	// register the overlay graphdriver
	_ "github.com/neo-hu/docker_el6/docker/daemon/graphdriver/overlay"
)
