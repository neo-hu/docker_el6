package daemon

import (
	apitypes "github.com/neo-hu/docker_el6/docker/api/types"
	lncluster "github.com/neo-hu/docker_el6/libnetwork/cluster"
)

// Cluster is the interface for github.com/neo-hu/docker_el6/docker/daemon/cluster.(*Cluster).
type Cluster interface {
	ClusterStatus
	NetworkManager
	SendClusterEvent(event lncluster.ConfigEventType)
}

// ClusterStatus interface provides information about the Swarm status of the Cluster
type ClusterStatus interface {
	IsAgent() bool
	IsManager() bool
}

// NetworkManager provides methods to manage networks
type NetworkManager interface {
	GetNetwork(input string) (apitypes.NetworkResource, error)
	GetNetworks() ([]apitypes.NetworkResource, error)
	RemoveNetwork(input string) error
}
