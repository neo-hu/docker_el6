package libnetwork

import (
	"github.com/neo-hu/docker_el6/libnetwork/drivers/null"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/remote"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/windows"
	"github.com/neo-hu/docker_el6/libnetwork/drivers/windows/overlay"
)

func getInitializers(experimental bool) []initializer {
	return []initializer{
		{null.Init, "null"},
		{overlay.Init, "overlay"},
		{remote.Init, "remote"},
		{windows.GetInit("transparent"), "transparent"},
		{windows.GetInit("l2bridge"), "l2bridge"},
		{windows.GetInit("l2tunnel"), "l2tunnel"},
		{windows.GetInit("nat"), "nat"},
		{windows.GetInit("ics"), "ics"},
	}
}
