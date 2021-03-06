package daemon

import (
	"sort"

	"github.com/neo-hu/docker_el6/docker/api/types/mount"
	"github.com/neo-hu/docker_el6/docker/container"
	"github.com/neo-hu/docker_el6/docker/pkg/idtools"
	"github.com/neo-hu/docker_el6/docker/volume"
)

// setupMounts configures the mount points for a container by appending each
// of the configured mounts on the container to the OCI mount structure
// which will ultimately be passed into the oci runtime during container creation.
// It also ensures each of the mounts are lexicographically sorted.

// BUGBUG TODO Windows containerd. This would be much better if it returned
// an array of runtime spec mounts, not container mounts. Then no need to
// do multiple transitions.

func (daemon *Daemon) setupMounts(c *container.Container) ([]container.Mount, error) {
	var mnts []container.Mount
	for _, mount := range c.MountPoints { // type is volume.MountPoint
		if err := daemon.lazyInitializeVolume(c.ID, mount); err != nil {
			return nil, err
		}
		s, err := mount.Setup(c.MountLabel, idtools.IDPair{0, 0}, nil)
		if err != nil {
			return nil, err
		}

		mnts = append(mnts, container.Mount{
			Source:      s,
			Destination: mount.Destination,
			Writable:    mount.RW,
		})
	}

	sort.Sort(mounts(mnts))
	return mnts, nil
}

// setBindModeIfNull is platform specific processing which is a no-op on
// Windows.
func setBindModeIfNull(bind *volume.MountPoint) {
	return
}

func (daemon *Daemon) validateBindDaemonRoot(m mount.Mount) (bool, error) {
	return false, nil
}
