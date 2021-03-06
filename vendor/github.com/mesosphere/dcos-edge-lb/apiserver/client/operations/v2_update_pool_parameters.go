// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/mesosphere/dcos-edge-lb/apiserver/models"
)

// NewV2UpdatePoolParams creates a new V2UpdatePoolParams object
// with the default values initialized.
func NewV2UpdatePoolParams() *V2UpdatePoolParams {
	var ()
	return &V2UpdatePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2UpdatePoolParamsWithTimeout creates a new V2UpdatePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2UpdatePoolParamsWithTimeout(timeout time.Duration) *V2UpdatePoolParams {
	var ()
	return &V2UpdatePoolParams{

		timeout: timeout,
	}
}

// NewV2UpdatePoolParamsWithContext creates a new V2UpdatePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2UpdatePoolParamsWithContext(ctx context.Context) *V2UpdatePoolParams {
	var ()
	return &V2UpdatePoolParams{

		Context: ctx,
	}
}

// NewV2UpdatePoolParamsWithHTTPClient creates a new V2UpdatePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2UpdatePoolParamsWithHTTPClient(client *http.Client) *V2UpdatePoolParams {
	var ()
	return &V2UpdatePoolParams{
		HTTPClient: client,
	}
}

/*V2UpdatePoolParams contains all the parameters to send to the API endpoint
for the v2 update pool operation typically these are written to a http.Request
*/
type V2UpdatePoolParams struct {

	/*Name*/
	Name string
	/*Pool*/
	Pool *models.V2Pool
	/*Token
	  DCOS Auth Token

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 update pool params
func (o *V2UpdatePoolParams) WithTimeout(timeout time.Duration) *V2UpdatePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 update pool params
func (o *V2UpdatePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 update pool params
func (o *V2UpdatePoolParams) WithContext(ctx context.Context) *V2UpdatePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 update pool params
func (o *V2UpdatePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 update pool params
func (o *V2UpdatePoolParams) WithHTTPClient(client *http.Client) *V2UpdatePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 update pool params
func (o *V2UpdatePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v2 update pool params
func (o *V2UpdatePoolParams) WithName(name string) *V2UpdatePoolParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v2 update pool params
func (o *V2UpdatePoolParams) SetName(name string) {
	o.Name = name
}

// WithPool adds the pool to the v2 update pool params
func (o *V2UpdatePoolParams) WithPool(pool *models.V2Pool) *V2UpdatePoolParams {
	o.SetPool(pool)
	return o
}

// SetPool adds the pool to the v2 update pool params
func (o *V2UpdatePoolParams) SetPool(pool *models.V2Pool) {
	o.Pool = pool
}

// WithToken adds the token to the v2 update pool params
func (o *V2UpdatePoolParams) WithToken(token *string) *V2UpdatePoolParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the v2 update pool params
func (o *V2UpdatePoolParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *V2UpdatePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Pool == nil {
		o.Pool = new(models.V2Pool)
	}

	if err := r.SetBodyParam(o.Pool); err != nil {
		return err
	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
