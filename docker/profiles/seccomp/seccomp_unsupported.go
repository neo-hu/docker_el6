// +build linux,!seccomp

package seccomp

import (
	"github.com/neo-hu/docker_el6/docker/api/types"
)

// DefaultProfile returns a nil pointer on unsupported systems.
func DefaultProfile() *types.Seccomp {
	return nil
}
