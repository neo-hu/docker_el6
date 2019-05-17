// +build linux solaris darwin freebsd

package local

import (
	"os"
	"syscall"
	"time"

	"github.com/neo-hu/docker_el6/containerd/sys"
)

func getATime(fi os.FileInfo) time.Time {
	if st, ok := fi.Sys().(*syscall.Stat_t); ok {
		return time.Unix(int64(sys.StatAtime(st).Sec),
			int64(sys.StatAtime(st).Nsec))
	}

	return fi.ModTime()
}
