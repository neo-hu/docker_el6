package client 

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/neo-hu/docker_el6/docker/api/types"
	"github.com/neo-hu/docker_el6/docker/api/types/swarm"
)

// ServiceInspectWithRaw returns the service information and the raw data.
func (cli *Client) ServiceInspectWithRaw(ctx context.Context, serviceID string, opts types.ServiceInspectOptions) (swarm.Service, []byte, error) {
	if serviceID == "" {
		return swarm.Service{}, nil, objectNotFoundError{object: "service", id: serviceID}
	}
	query := url.Values{}
	query.Set("insertDefaults", fmt.Sprintf("%v", opts.InsertDefaults))
	serverResp, err := cli.get(ctx, "/services/"+serviceID, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return swarm.Service{}, nil, wrapResponseError(err, serverResp, "service", serviceID)
	}

	body, err := ioutil.ReadAll(serverResp.body)
	if err != nil {
		return swarm.Service{}, nil, err
	}

	var response swarm.Service
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&response)
	return response, body, err
}