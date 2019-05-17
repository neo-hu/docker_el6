package fstest

import (
	"time"

	"github.com/neo-hu/docker_el6/containerd/errdefs"
)

// Lchtimes changes access and mod time of file without following symlink
func Lchtimes(name string, atime, mtime time.Time) Applier {
	return applyFn(func(root string) error {
		return errdefs.ErrNotImplemented
	})
}
