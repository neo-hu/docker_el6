package client 

import (
	"context"
	"encoding/json"

	"github.com/neo-hu/docker_el6/docker/api/types"
)

// ServerVersion returns information of the docker client and server host.
func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error) {
	resp, err := cli.get(ctx, "/version", nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return types.Version{}, err
	}

	var server types.Version
	err = json.NewDecoder(resp.body).Decode(&server)
	return server, err
}
