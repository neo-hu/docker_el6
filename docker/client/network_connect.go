package client 

import (
	"context"

	"github.com/neo-hu/docker_el6/docker/api/types"
	"github.com/neo-hu/docker_el6/docker/api/types/network"
)

// NetworkConnect connects a container to an existent network in the docker host.
func (cli *Client) NetworkConnect(ctx context.Context, networkID, containerID string, config *network.EndpointSettings) error {
	nc := types.NetworkConnect{
		Container:      containerID,
		EndpointConfig: config,
	}
	resp, err := cli.post(ctx, "/networks/"+networkID+"/connect", nil, nc, nil)
	ensureReaderClosed(resp)
	return err
}
