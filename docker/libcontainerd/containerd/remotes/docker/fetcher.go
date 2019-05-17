package docker

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/neo-hu/docker_el6/containerd/errdefs"
	"github.com/neo-hu/docker_el6/containerd/images"
	"github.com/neo-hu/docker_el6/containerd/log"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type dockerFetcher struct {
	*dockerBase
}

func (r dockerFetcher) Fetch(ctx context.Context, desc ocispec.Descriptor) (io.ReadCloser, error) {
	ctx = log.WithLogger(ctx, log.G(ctx).WithFields(
		logrus.Fields{
			"base":   r.base.String(),
			"digest": desc.Digest,
		},
	))

	urls, err := r.getV2URLPaths(ctx, desc)
	if err != nil {
		return nil, err
	}

	ctx, err = contextWithRepositoryScope(ctx, r.refspec, false)
	if err != nil {
		return nil, err
	}

	return newHTTPReadSeeker(desc.Size, func(offset int64) (io.ReadCloser, error) {
		for _, u := range urls {
			rc, err := r.open(ctx, u, desc.MediaType, offset)
			if err != nil {
				if errdefs.IsNotFound(err) {
					continue // try one of the other urls.
				}

				return nil, err
			}

			return rc, nil
		}

		return nil, errors.Wrapf(errdefs.ErrNotFound,
			"could not fetch content descriptor %v (%v) from remote",
			desc.Digest, desc.MediaType)

	})
}

func (r dockerFetcher) open(ctx context.Context, u, mediatype string, offset int64) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", strings.Join([]string{mediatype, `*`}, ", "))

	if offset > 0 {
		// TODO(stevvooe): Only set this header in response to the
		// "Accept-Ranges: bytes" header.
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", offset))
	}

	resp, err := r.doRequestWithRetries(ctx, req, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		// TODO(stevvooe): When doing a offset specific request, we should
		// really distinguish between a 206 and a 200. In the case of 200, we
		// can discard the bytes, hiding the seek behavior from the
		// implementation.

		resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			return nil, errors.Wrapf(errdefs.ErrNotFound, "content at %v not found", u)
		}
		return nil, errors.Errorf("unexpected status code %v: %v", u, resp.Status)
	}

	return resp.Body, nil
}

// getV2URLPaths generates the candidate urls paths for the object based on the
// set of hints and the provided object id. URLs are returned in the order of
// most to least likely succeed.
func (r *dockerFetcher) getV2URLPaths(ctx context.Context, desc ocispec.Descriptor) ([]string, error) {
	var urls []string

	if len(desc.URLs) > 0 {
		// handle fetch via external urls.
		for _, u := range desc.URLs {
			log.G(ctx).WithField("url", u).Debug("adding alternative url")
			urls = append(urls, u)
		}
	}

	switch desc.MediaType {
	case images.MediaTypeDockerSchema2Manifest, images.MediaTypeDockerSchema2ManifestList,
		images.MediaTypeDockerSchema1Manifest,
		ocispec.MediaTypeImageManifest, ocispec.MediaTypeImageIndex:
		urls = append(urls, r.url(path.Join("manifests", desc.Digest.String())))
	}

	// always fallback to attempting to get the object out of the blobs store.
	urls = append(urls, r.url(path.Join("blobs", desc.Digest.String())))

	return urls, nil
}
