package registry

import (
	"net/url"

	"github.com/neo-hu/docker_el6/distribution/registry/api/errcode"
	"github.com/neo-hu/docker_el6/docker/errdefs"
)

type notFoundError string

func (e notFoundError) Error() string {
	return string(e)
}

func (notFoundError) NotFound() {}

func translateV2AuthError(err error) error {
	switch e := err.(type) {
	case *url.Error:
		switch e2 := e.Err.(type) {
		case errcode.Error:
			switch e2.Code {
			case errcode.ErrorCodeUnauthorized:
				return errdefs.Unauthorized(err)
			}
		}
	}

	return err
}
