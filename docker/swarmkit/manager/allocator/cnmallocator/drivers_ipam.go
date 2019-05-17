package cnmallocator

import (
	"github.com/neo-hu/docker_el6/libnetwork/drvregistry"
	"github.com/neo-hu/docker_el6/libnetwork/ipamapi"
	builtinIpam "github.com/neo-hu/docker_el6/libnetwork/ipams/builtin"
	nullIpam "github.com/neo-hu/docker_el6/libnetwork/ipams/null"
	remoteIpam "github.com/neo-hu/docker_el6/libnetwork/ipams/remote"
)

func initIPAMDrivers(r *drvregistry.DrvRegistry) error {
	for _, fn := range [](func(ipamapi.Callback, interface{}, interface{}) error){
		builtinIpam.Init,
		remoteIpam.Init,
		nullIpam.Init,
	} {
		if err := fn(r, nil, nil); err != nil {
			return err
		}
	}

	return nil
}
