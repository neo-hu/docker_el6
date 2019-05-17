package daemon 

import (
	"testing"

	"github.com/neo-hu/docker_el6/docker/api/types/network"
	"github.com/neo-hu/docker_el6/docker/errdefs"
	"github.com/gotestyourself/gotestyourself/assert"
)

// Test case for 35752
func TestVerifyNetworkingConfig(t *testing.T) {
	name := "mynet"
	endpoints := make(map[string]*network.EndpointSettings, 1)
	endpoints[name] = nil
	nwConfig := &network.NetworkingConfig{
		EndpointsConfig: endpoints,
	}
	err := verifyNetworkingConfig(nwConfig)
	assert.Check(t, errdefs.IsInvalidParameter(err))
}
