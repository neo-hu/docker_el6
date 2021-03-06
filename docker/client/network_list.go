package client 

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/neo-hu/docker_el6/docker/api/types"
	"github.com/neo-hu/docker_el6/docker/api/types/filters"
)

// NetworkList returns the list of networks configured in the docker host.
func (cli *Client) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error) {
	query := url.Values{}
	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToParamWithVersion(cli.version, options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}
	var networkResources []types.NetworkResource
	resp, err := cli.get(ctx, "/networks", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return networkResources, err
	}
	err = json.NewDecoder(resp.body).Decode(&networkResources)
	return networkResources, err
}
