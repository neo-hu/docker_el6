package daemon

import (
	"github.com/neo-hu/docker_el6/docker/pkg/archive"
)

// defaultTarCopyOptions is the setting that is used when unpacking an archive
// for a copy API event.
func (daemon *Daemon) defaultTarCopyOptions(noOverwriteDirNonDir bool) *archive.TarOptions {
	return &archive.TarOptions{
		NoOverwriteDirNonDir: noOverwriteDirNonDir,
		UIDMaps:              daemon.idMappings.UIDs(),
		GIDMaps:              daemon.idMappings.GIDs(),
	}
}
