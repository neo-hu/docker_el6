package controlapi

import (
	"errors"

	"github.com/neo-hu/docker_el6/docker/pkg/plugingetter"
	"github.com/neo-hu/docker_el6/docker/swarmkit/ca"
	"github.com/neo-hu/docker_el6/docker/swarmkit/manager/drivers"
	"github.com/neo-hu/docker_el6/docker/swarmkit/manager/state/raft"
	"github.com/neo-hu/docker_el6/docker/swarmkit/manager/state/store"
)

var (
	errInvalidArgument = errors.New("invalid argument")
)

// Server is the Cluster API gRPC server.
type Server struct {
	store          *store.MemoryStore
	raft           *raft.Node
	securityConfig *ca.SecurityConfig
	pg             plugingetter.PluginGetter
	dr             *drivers.DriverProvider
}

// NewServer creates a Cluster API server.
func NewServer(store *store.MemoryStore, raft *raft.Node, securityConfig *ca.SecurityConfig, pg plugingetter.PluginGetter, dr *drivers.DriverProvider) *Server {
	return &Server{
		store:          store,
		dr:             dr,
		raft:           raft,
		securityConfig: securityConfig,
		pg:             pg,
	}
}
