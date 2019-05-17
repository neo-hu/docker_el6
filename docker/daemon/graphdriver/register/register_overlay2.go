// +build !exclude_graphdriver_overlay2,linux

package register

import (
	// register the overlay2 graphdriver
	_ "github.com/neo-hu/docker_el6/docker/daemon/graphdriver/overlay2"
)
