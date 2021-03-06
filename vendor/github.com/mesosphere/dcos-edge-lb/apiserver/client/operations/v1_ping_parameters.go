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
)

// NewV1PingParams creates a new V1PingParams object
// with the default values initialized.
func NewV1PingParams() *V1PingParams {

	return &V1PingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1PingParamsWithTimeout creates a new V1PingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1PingParamsWithTimeout(timeout time.Duration) *V1PingParams {

	return &V1PingParams{

		timeout: timeout,
	}
}

// NewV1PingParamsWithContext creates a new V1PingParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1PingParamsWithContext(ctx context.Context) *V1PingParams {

	return &V1PingParams{

		Context: ctx,
	}
}

// NewV1PingParamsWithHTTPClient creates a new V1PingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1PingParamsWithHTTPClient(client *http.Client) *V1PingParams {

	return &V1PingParams{
		HTTPClient: client,
	}
}

/*V1PingParams contains all the parameters to send to the API endpoint
for the v1 ping operation typically these are written to a http.Request
*/
type V1PingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 ping params
func (o *V1PingParams) WithTimeout(timeout time.Duration) *V1PingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 ping params
func (o *V1PingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 ping params
func (o *V1PingParams) WithContext(ctx context.Context) *V1PingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 ping params
func (o *V1PingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 ping params
func (o *V1PingParams) WithHTTPClient(client *http.Client) *V1PingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 ping params
func (o *V1PingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1PingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
