package libcontainer

import "github.com/neo-hu/docker_el6/cmd/runc/libcontainer/cgroups"
import "github.com/neo-hu/docker_el6/cmd/runc/libcontainer/intelrdt"

type Stats struct {
	Interfaces    []*NetworkInterface
	CgroupStats   *cgroups.Stats
	IntelRdtStats *intelrdt.Stats
}
