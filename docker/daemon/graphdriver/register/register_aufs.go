// +build !exclude_graphdriver_aufs,linux

package register 

import (
	// register the aufs graphdriver
	_ "github.com/neo-hu/docker_el6/docker/daemon/graphdriver/aufs"
)
