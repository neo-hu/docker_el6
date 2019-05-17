// +build !windows

package libnetwork

import "github.com/neo-hu/docker_el6/libnetwork/ipamapi"

// Stub implementations for DNS related functions

func (n *network) startResolver() {
}

func defaultIpamForNetworkType(networkType string) string {
	return ipamapi.DefaultIPAM
}
