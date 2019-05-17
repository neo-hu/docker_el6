package macvlan

import (
	"fmt"

	"github.com/neo-hu/docker_el6/libnetwork/driverapi"
	"github.com/neo-hu/docker_el6/libnetwork/netlabel"
	"github.com/neo-hu/docker_el6/libnetwork/netutils"
	"github.com/neo-hu/docker_el6/libnetwork/ns"
	"github.com/neo-hu/docker_el6/libnetwork/osl"
	"github.com/neo-hu/docker_el6/libnetwork/types"
	"github.com/sirupsen/logrus"
)

// CreateEndpoint assigns the mac, ip and endpoint id for the new container
func (d *driver) CreateEndpoint(nid, eid string, ifInfo driverapi.InterfaceInfo,
	epOptions map[string]interface{}) error {
	defer osl.InitOSContext()()

	if err := validateID(nid, eid); err != nil {
		return err
	}
	n, err := d.getNetwork(nid)
	if err != nil {
		return fmt.Errorf("network id %q not found", nid)
	}
	ep := &endpoint{
		id:     eid,
		nid:    nid,
		addr:   ifInfo.Address(),
		addrv6: ifInfo.AddressIPv6(),
		mac:    ifInfo.MacAddress(),
	}
	if ep.addr == nil {
		return fmt.Errorf("create endpoint was not passed an IP address")
	}
	if ep.mac == nil {
		ep.mac = netutils.GenerateMACFromIP(ep.addr.IP)
		if err := ifInfo.SetMacAddress(ep.mac); err != nil {
			return err
		}
	}
	// disallow portmapping -p
	if opt, ok := epOptions[netlabel.PortMap]; ok {
		if _, ok := opt.([]types.PortBinding); ok {
			if len(opt.([]types.PortBinding)) > 0 {
				logrus.Warnf("%s driver does not support port mappings", macvlanType)
			}
		}
	}
	// disallow port exposure --expose
	if opt, ok := epOptions[netlabel.ExposedPorts]; ok {
		if _, ok := opt.([]types.TransportPort); ok {
			if len(opt.([]types.TransportPort)) > 0 {
				logrus.Warnf("%s driver does not support port exposures", macvlanType)
			}
		}
	}

	if err := d.storeUpdate(ep); err != nil {
		return fmt.Errorf("failed to save macvlan endpoint %s to store: %v", ep.id[0:7], err)
	}

	n.addEndpoint(ep)

	return nil
}

// DeleteEndpoint removes the endpoint and associated netlink interface
func (d *driver) DeleteEndpoint(nid, eid string) error {
	defer osl.InitOSContext()()
	if err := validateID(nid, eid); err != nil {
		return err
	}
	n := d.network(nid)
	if n == nil {
		return fmt.Errorf("network id %q not found", nid)
	}
	ep := n.endpoint(eid)
	if ep == nil {
		return fmt.Errorf("endpoint id %q not found", eid)
	}
	if link, err := ns.NlHandle().LinkByName(ep.srcName); err == nil {
		if err := ns.NlHandle().LinkDel(link); err != nil {
			logrus.WithError(err).Warnf("Failed to delete interface (%s)'s link on endpoint (%s) delete", ep.srcName, ep.id)
		}
	}

	if err := d.storeDelete(ep); err != nil {
		logrus.Warnf("Failed to remove macvlan endpoint %s from store: %v", ep.id[0:7], err)
	}

	n.deleteEndpoint(ep.id)

	return nil
}
