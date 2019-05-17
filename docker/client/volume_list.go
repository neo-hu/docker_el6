package client 

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/neo-hu/docker_el6/docker/api/types/filters"
	volumetypes "github.com/neo-hu/docker_el6/docker/api/types/volume"
)

// VolumeList returns the volumes configured in the docker host.
func (cli *Client) VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error) {
	var volumes volumetypes.VolumesListOKBody
	query := url.Values{}

	if filter.Len() > 0 {
		filterJSON, err := filters.ToParamWithVersion(cli.version, filter)
		if err != nil {
			return volumes, err
		}
		query.Set("filters", filterJSON)
	}
	resp, err := cli.get(ctx, "/volumes", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return volumes, err
	}

	err = json.NewDecoder(resp.body).Decode(&volumes)
	return volumes, err
}
