// +build !linux,!darwin,!windows

package cnmallocator

import (
	"github.com/neo-hu/docker_el6/docker/swarmkit/manager/allocator/networkallocator"
)

const initializers = nil

// PredefinedNetworks returns the list of predefined network structures
func PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return nil
}
