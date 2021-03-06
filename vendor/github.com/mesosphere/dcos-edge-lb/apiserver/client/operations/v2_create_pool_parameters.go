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

// NewV2CreatePoolParams creates a new V2CreatePoolParams object
// with the default values initialized.
func NewV2CreatePoolParams() *V2CreatePoolParams {
	var ()
	return &V2CreatePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2CreatePoolParamsWithTimeout creates a new V2CreatePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2CreatePoolParamsWithTimeout(timeout time.Duration) *V2CreatePoolParams {
	var ()
	return &V2CreatePoolParams{

		timeout: timeout,
	}
}

// NewV2CreatePoolParamsWithContext creates a new V2CreatePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2CreatePoolParamsWithContext(ctx context.Context) *V2CreatePoolParams {
	var ()
	return &V2CreatePoolParams{

		Context: ctx,
	}
}

// NewV2CreatePoolParamsWithHTTPClient creates a new V2CreatePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2CreatePoolParamsWithHTTPClient(client *http.Client) *V2CreatePoolParams {
	var ()
	return &V2CreatePoolParams{
		HTTPClient: client,
	}
}

/*V2CreatePoolParams contains all the parameters to send to the API endpoint
for the v2 create pool operation typically these are written to a http.Request
*/
type V2CreatePoolParams struct {

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

// WithTimeout adds the timeout to the v2 create pool params
func (o *V2CreatePoolParams) WithTimeout(timeout time.Duration) *V2CreatePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 create pool params
func (o *V2CreatePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 create pool params
func (o *V2CreatePoolParams) WithContext(ctx context.Context) *V2CreatePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 create pool params
func (o *V2CreatePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 create pool params
func (o *V2CreatePoolParams) WithHTTPClient(client *http.Client) *V2CreatePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 create pool params
func (o *V2CreatePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPool adds the pool to the v2 create pool params
func (o *V2CreatePoolParams) WithPool(pool *models.V2Pool) *V2CreatePoolParams {
	o.SetPool(pool)
	return o
}

// SetPool adds the pool to the v2 create pool params
func (o *V2CreatePoolParams) SetPool(pool *models.V2Pool) {
	o.Pool = pool
}

// WithToken adds the token to the v2 create pool params
func (o *V2CreatePoolParams) WithToken(token *string) *V2CreatePoolParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the v2 create pool params
func (o *V2CreatePoolParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *V2CreatePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
