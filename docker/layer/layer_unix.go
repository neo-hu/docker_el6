// +build linux freebsd darwin openbsd

package layer 

import "github.com/neo-hu/docker_el6/docker/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
