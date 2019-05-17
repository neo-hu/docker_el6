package libnetwork

import (
	"github.com/neo-hu/docker_el6/libnetwork/drivers/null"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/remote"
)

func getInitializers(experimental bool) []initializer {
	return []initializer{
		{null.Init, "null"},
		{remote.Init, "remote"},
	}
}
