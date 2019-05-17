// +build !linux

package vfs 

import "github.com/neo-hu/docker_el6/docker/pkg/chrootarchive"

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(nil).CopyWithTar(srcDir, dstDir)
}
