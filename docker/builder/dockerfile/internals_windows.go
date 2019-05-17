package dockerfile 

import "github.com/neo-hu/docker_el6/docker/pkg/idtools"

func parseChownFlag(chown, ctrRootPath string, idMappings *idtools.IDMappings) (idtools.IDPair, error) {
	return idMappings.RootPair(), nil
}
