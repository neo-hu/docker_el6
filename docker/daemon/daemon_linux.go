package daemon 

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/neo-hu/docker_el6/docker/pkg/fileutils"
	"github.com/neo-hu/docker_el6/docker/pkg/mount"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// On Linux, plugins use a static path for storing execution state,
// instead of deriving path from daemon's exec-root. This is because
// plugin socket files are created here and they cannot exceed max
// path length of 108 bytes.
func getPluginExecRoot(root string) string {
	return "/run/docker/plugins"
}

func (daemon *Daemon) cleanupMountsByID(id string) error {
	logrus.Debugf("Cleaning up old mountid %s: start.", id)
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return err
	}
	defer f.Close()

	return daemon.cleanupMountsFromReaderByID(f, id, mount.Unmount)
}

func (daemon *Daemon) cleanupMountsFromReaderByID(reader io.Reader, id string, unmount func(target string) error) error {
	if daemon.root == "" {
		return nil
	}
	var errors []string

	regexps := getCleanPatterns(id)
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		if fields := strings.Fields(sc.Text()); len(fields) >= 4 {
			if mnt := fields[4]; strings.HasPrefix(mnt, daemon.root) {
				for _, p := range regexps {
					if p.MatchString(mnt) {
						if err := unmount(mnt); err != nil {
							logrus.Error(err)
							errors = append(errors, err.Error())
						}
					}
				}
			}
		}
	}

	if err := sc.Err(); err != nil {
		return err
	}

	if len(errors) > 0 {
		return fmt.Errorf("Error cleaning up mounts:\n%v", strings.Join(errors, "\n"))
	}

	logrus.Debugf("Cleaning up old mountid %v: done.", id)
	return nil
}

// cleanupMounts umounts used by container resources and the daemon root mount
func (daemon *Daemon) cleanupMounts() error {
	if err := daemon.cleanupMountsByID(""); err != nil {
		return err
	}

	infos, err := mount.GetMounts()
	if err != nil {
		return errors.Wrap(err, "error reading mount table for cleanup")
	}

	info := getMountInfo(infos, daemon.root)
	// `info.Root` here is the root mountpoint of the passed in path (`daemon.root`).
	// The ony cases that need to be cleaned up is when the daemon has performed a
	//   `mount --bind /daemon/root /daemon/root && mount --make-shared /daemon/root`
	// This is only done when the daemon is started up and `/daemon/root` is not
	// already on a shared mountpoint.
	if !shouldUnmountRoot(daemon.root, info) {
		return nil
	}

	logrus.WithField("mountpoint", daemon.root).Debug("unmounting daemon root")
	return mount.Unmount(daemon.root)
}

func getCleanPatterns(id string) (regexps []*regexp.Regexp) {
	var patterns []string
	if id == "" {
		id = "[0-9a-f]{64}"
		patterns = append(patterns, "containers/"+id+"/shm")
	}
	patterns = append(patterns, "aufs/mnt/"+id+"$", "overlay/"+id+"/merged$", "zfs/graph/"+id+"$")
	for _, p := range patterns {
		r, err := regexp.Compile(p)
		if err == nil {
			regexps = append(regexps, r)
		}
	}
	return
}

func getRealPath(path string) (string, error) {
	return fileutils.ReadSymlinkedDirectory(path)
}

func shouldUnmountRoot(root string, info *mount.Info) bool {
	if info == nil {
		return false
	}
	if info.Mountpoint != root {
		return false
	}
	if !strings.HasSuffix(root, info.Root) {
		return false
	}
	return hasMountinfoOption(info.Optional, sharedPropagationOption)
}
