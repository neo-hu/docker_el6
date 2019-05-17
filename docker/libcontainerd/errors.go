package libcontainerd

import (
	"errors"

	"github.com/neo-hu/docker_el6/docker/errdefs"
)

func newNotFoundError(err string) error { return errdefs.NotFound(errors.New(err)) }

func newInvalidParameterError(err string) error { return errdefs.InvalidParameter(errors.New(err)) }

func newConflictError(err string) error { return errdefs.Conflict(errors.New(err)) }
