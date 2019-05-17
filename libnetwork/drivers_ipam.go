package libnetwork

import (
	"github.com/neo-hu/docker_el6/libnetwork/drvregistry"
	"github.com/neo-hu/docker_el6/libnetwork/ipamapi"
	builtinIpam "github.com/neo-hu/docker_el6/libnetwork/ipams/builtin"
	nullIpam "github.com/neo-hu/docker_el6/libnetwork/ipams/null"
	remoteIpam "github.com/neo-hu/docker_el6/libnetwork/ipams/remote"
	"github.com/neo-hu/docker_el6/libnetwork/ipamutils"
)

func initIPAMDrivers(r *drvregistry.DrvRegistry, lDs, gDs interface{}, addressPool []*ipamutils.NetworkToSplit) error {
	builtinIpam.SetDefaultIPAddressPool(addressPool)
	for _, fn := range [](func(ipamapi.Callback, interface{}, interface{}) error){
		builtinIpam.Init,
		remoteIpam.Init,
		nullIpam.Init,
	} {
		if err := fn(r, lDs, gDs); err != nil {
			return err
		}
	}

	return nil
}
