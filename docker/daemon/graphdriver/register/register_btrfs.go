// +build !exclude_graphdriver_btrfs,linux

package register

import (
	// register the btrfs graphdriver
	_ "github.com/neo-hu/docker_el6/docker/daemon/graphdriver/btrfs"
)
