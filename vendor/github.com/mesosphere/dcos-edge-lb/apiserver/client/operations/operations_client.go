// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetConfigContainer Get the entire configuration object including v1 and v2 pools.
*/
func (a *Client) GetConfigContainer(params *GetConfigContainerParams) (*GetConfigContainerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigContainerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConfigContainer",
		Method:             "GET",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigContainerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigContainerOK), nil

}

/*
GetPoolContainer Returns a v1 or v2 load balancer pool based on a single name.
*/
func (a *Client) GetPoolContainer(params *GetPoolContainerParams) (*GetPoolContainerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolContainerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPoolContainer",
		Method:             "GET",
		PathPattern:        "/pools/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPoolContainerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPoolContainerOK), nil

}

/*
Ping Healthcheck endpoint.
*/
func (a *Client) Ping(params *PingParams) (*PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Ping",
		Method:             "GET",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PingOK), nil

}

/*
V1CreateLoadBalancerPool Creates a new load balancer pool.
*/
func (a *Client) V1CreateLoadBalancerPool(params *V1CreateLoadBalancerPoolParams) (*V1CreateLoadBalancerPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1CreateLoadBalancerPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1CreateLoadBalancerPool",
		Method:             "POST",
		PathPattern:        "/v1/loadbalancers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1CreateLoadBalancerPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1CreateLoadBalancerPoolOK), nil

}

/*
V1DeleteLoadBalancerArtifact Clears an override for a artifact. Not all artifacts can be deleted.
*/
func (a *Client) V1DeleteLoadBalancerArtifact(params *V1DeleteLoadBalancerArtifactParams) (*V1DeleteLoadBalancerArtifactOK, *V1DeleteLoadBalancerArtifactNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1DeleteLoadBalancerArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1DeleteLoadBalancerArtifact",
		Method:             "DELETE",
		PathPattern:        "/v1/loadbalancers/{name}/artifacts/{artifactName}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1DeleteLoadBalancerArtifactReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *V1DeleteLoadBalancerArtifactOK:
		return value, nil, nil
	case *V1DeleteLoadBalancerArtifactNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
V1DeleteLoadBalancerPool Deletes a single load balancer pool based on the name supplied.
*/
func (a *Client) V1DeleteLoadBalancerPool(params *V1DeleteLoadBalancerPoolParams) (*V1DeleteLoadBalancerPoolNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1DeleteLoadBalancerPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1DeleteLoadBalancerPool",
		Method:             "DELETE",
		PathPattern:        "/v1/loadbalancers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1DeleteLoadBalancerPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1DeleteLoadBalancerPoolNoContent), nil

}

/*
V1GetConfig Get the entire configuration.
*/
func (a *Client) V1GetConfig(params *V1GetConfigParams) (*V1GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1GetConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1GetConfig",
		Method:             "GET",
		PathPattern:        "/v1/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1GetConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1GetConfigOK), nil

}

/*
V1GetLoadBalancerArtifact Returns an configuration artifact for a load balancer pool.
*/
func (a *Client) V1GetLoadBalancerArtifact(params *V1GetLoadBalancerArtifactParams) (*V1GetLoadBalancerArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1GetLoadBalancerArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1GetLoadBalancerArtifact",
		Method:             "GET",
		PathPattern:        "/v1/loadbalancers/{name}/artifacts/{artifactName}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1GetLoadBalancerArtifactReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1GetLoadBalancerArtifactOK), nil

}

/*
V1GetLoadBalancerArtifacts Returns artifacts available for a load balancer pool.
*/
func (a *Client) V1GetLoadBalancerArtifacts(params *V1GetLoadBalancerArtifactsParams) (*V1GetLoadBalancerArtifactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1GetLoadBalancerArtifactsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1GetLoadBalancerArtifacts",
		Method:             "GET",
		PathPattern:        "/v1/loadbalancers/{name}/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1GetLoadBalancerArtifactsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1GetLoadBalancerArtifactsOK), nil

}

/*
V1GetLoadBalancerPool Returns a load balancer pool based on a single name.
*/
func (a *Client) V1GetLoadBalancerPool(params *V1GetLoadBalancerPoolParams) (*V1GetLoadBalancerPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1GetLoadBalancerPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1GetLoadBalancerPool",
		Method:             "GET",
		PathPattern:        "/v1/loadbalancers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1GetLoadBalancerPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1GetLoadBalancerPoolOK), nil

}

/*
V1GetLoadBalancerPools Get all load balancer pools.
*/
func (a *Client) V1GetLoadBalancerPools(params *V1GetLoadBalancerPoolsParams) (*V1GetLoadBalancerPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1GetLoadBalancerPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1GetLoadBalancerPools",
		Method:             "GET",
		PathPattern:        "/v1/loadbalancers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1GetLoadBalancerPoolsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1GetLoadBalancerPoolsOK), nil

}

/*
V1Ping Healthcheck endpoint.
*/
func (a *Client) V1Ping(params *V1PingParams) (*V1PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1PingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1Ping",
		Method:             "GET",
		PathPattern:        "/v1/ping",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1PingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1PingOK), nil

}

/*
V1UpdateConfig Wipes and overwrites the entire configuration.
*/
func (a *Client) V1UpdateConfig(params *V1UpdateConfigParams) (*V1UpdateConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1UpdateConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1UpdateConfig",
		Method:             "PUT",
		PathPattern:        "/v1/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1UpdateConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1UpdateConfigOK), nil

}

/*
V1UpdateLoadBalancerArtifact Creates or updates an configuration artifact for a load balancer pool. Not all artifacts can be changed.
*/
func (a *Client) V1UpdateLoadBalancerArtifact(params *V1UpdateLoadBalancerArtifactParams) (*V1UpdateLoadBalancerArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1UpdateLoadBalancerArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1UpdateLoadBalancerArtifact",
		Method:             "PUT",
		PathPattern:        "/v1/loadbalancers/{name}/artifacts/{artifactName}",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1UpdateLoadBalancerArtifactReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1UpdateLoadBalancerArtifactOK), nil

}

/*
V1UpdateLoadBalancerPool Updates a new load balancer pool.
*/
func (a *Client) V1UpdateLoadBalancerPool(params *V1UpdateLoadBalancerPoolParams) (*V1UpdateLoadBalancerPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1UpdateLoadBalancerPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1UpdateLoadBalancerPool",
		Method:             "PUT",
		PathPattern:        "/v1/loadbalancers/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1UpdateLoadBalancerPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1UpdateLoadBalancerPoolOK), nil

}

/*
V1Version Returns the installed Edge-LB package version.
*/
func (a *Client) V1Version(params *V1VersionParams) (*V1VersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1VersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V1Version",
		Method:             "GET",
		PathPattern:        "/v1/version",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V1VersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V1VersionOK), nil

}

/*
V2CreatePool Creates a new load balancer pool.
*/
func (a *Client) V2CreatePool(params *V2CreatePoolParams) (*V2CreatePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2CreatePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2CreatePool",
		Method:             "POST",
		PathPattern:        "/v2/pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2CreatePoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2CreatePoolOK), nil

}

/*
V2DeleteLBTemplate Resets the lb config template to default for a pool.
*/
func (a *Client) V2DeleteLBTemplate(params *V2DeleteLBTemplateParams) (*V2DeleteLBTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2DeleteLBTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2DeleteLBTemplate",
		Method:             "DELETE",
		PathPattern:        "/v2/pools/{name}/lbtemplate",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2DeleteLBTemplateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2DeleteLBTemplateOK), nil

}

/*
V2DeletePool Deletes a single load balancer pool based on the name supplied.
*/
func (a *Client) V2DeletePool(params *V2DeletePoolParams) (*V2DeletePoolNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2DeletePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2DeletePool",
		Method:             "DELETE",
		PathPattern:        "/v2/pools/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2DeletePoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2DeletePoolNoContent), nil

}

/*
V2GetConfig Get the entire configuration object including only v2 pools.
*/
func (a *Client) V2GetConfig(params *V2GetConfigParams) (*V2GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2GetConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2GetConfig",
		Method:             "GET",
		PathPattern:        "/v2/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2GetConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2GetConfigOK), nil

}

/*
V2GetDefaultLBTemplate Returns the default lb config template.
*/
func (a *Client) V2GetDefaultLBTemplate(params *V2GetDefaultLBTemplateParams) (*V2GetDefaultLBTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2GetDefaultLBTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2GetDefaultLBTemplate",
		Method:             "GET",
		PathPattern:        "/v2/lbtemplate",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2GetDefaultLBTemplateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2GetDefaultLBTemplateOK), nil

}

/*
V2GetLBConfig Returns the rendered lb config for a pool.
*/
func (a *Client) V2GetLBConfig(params *V2GetLBConfigParams) (*V2GetLBConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2GetLBConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2GetLBConfig",
		Method:             "GET",
		PathPattern:        "/v2/pools/{name}/lbconfig",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2GetLBConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2GetLBConfigOK), nil

}

/*
V2GetLBTemplate Returns the lb config template for a pool.
*/
func (a *Client) V2GetLBTemplate(params *V2GetLBTemplateParams) (*V2GetLBTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2GetLBTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2GetLBTemplate",
		Method:             "GET",
		PathPattern:        "/v2/pools/{name}/lbtemplate",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2GetLBTemplateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2GetLBTemplateOK), nil

}

/*
V2GetPool Returns a v2 load balancer pool based on a single name.
*/
func (a *Client) V2GetPool(params *V2GetPoolParams) (*V2GetPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2GetPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2GetPool",
		Method:             "GET",
		PathPattern:        "/v2/pools/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2GetPoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2GetPoolOK), nil

}

/*
V2GetPools Get all load balancer pools.
*/
func (a *Client) V2GetPools(params *V2GetPoolsParams) (*V2GetPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2GetPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2GetPools",
		Method:             "GET",
		PathPattern:        "/v2/pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2GetPoolsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2GetPoolsOK), nil

}

/*
V2UpdateLBTemplate Creates or updates the lb config template for a pool.
*/
func (a *Client) V2UpdateLBTemplate(params *V2UpdateLBTemplateParams) (*V2UpdateLBTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2UpdateLBTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2UpdateLBTemplate",
		Method:             "PUT",
		PathPattern:        "/v2/pools/{name}/lbtemplate",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2UpdateLBTemplateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2UpdateLBTemplateOK), nil

}

/*
V2UpdatePool Updates a new load balancer pool.
*/
func (a *Client) V2UpdatePool(params *V2UpdatePoolParams) (*V2UpdatePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV2UpdatePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2UpdatePool",
		Method:             "PUT",
		PathPattern:        "/v2/pools/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2UpdatePoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2UpdatePoolOK), nil

}

/*
Version Returns the installed Edge-LB package version.
*/
func (a *Client) Version(params *VersionParams) (*VersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Version",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
