package distribution

import (
	"github.com/neo-hu/docker_el6/distribution"
	"github.com/neo-hu/docker_el6/distribution/reference"
	"github.com/neo-hu/docker_el6/docker/api/types"
	"golang.org/x/net/context"
)

// Backend is all the methods that need to be implemented
// to provide image specific functionality.
type Backend interface {
	GetRepository(context.Context, reference.Named, *types.AuthConfig) (distribution.Repository, bool, error)
}
