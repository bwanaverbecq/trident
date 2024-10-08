// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewNetworkIPRouteDeleteParams creates a new NetworkIPRouteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPRouteDeleteParams() *NetworkIPRouteDeleteParams {
	return &NetworkIPRouteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPRouteDeleteParamsWithTimeout creates a new NetworkIPRouteDeleteParams object
// with the ability to set a timeout on a request.
func NewNetworkIPRouteDeleteParamsWithTimeout(timeout time.Duration) *NetworkIPRouteDeleteParams {
	return &NetworkIPRouteDeleteParams{
		timeout: timeout,
	}
}

// NewNetworkIPRouteDeleteParamsWithContext creates a new NetworkIPRouteDeleteParams object
// with the ability to set a context for a request.
func NewNetworkIPRouteDeleteParamsWithContext(ctx context.Context) *NetworkIPRouteDeleteParams {
	return &NetworkIPRouteDeleteParams{
		Context: ctx,
	}
}

// NewNetworkIPRouteDeleteParamsWithHTTPClient creates a new NetworkIPRouteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPRouteDeleteParamsWithHTTPClient(client *http.Client) *NetworkIPRouteDeleteParams {
	return &NetworkIPRouteDeleteParams{
		HTTPClient: client,
	}
}

/*
NetworkIPRouteDeleteParams contains all the parameters to send to the API endpoint

	for the network ip route delete operation.

	Typically these are written to a http.Request.
*/
type NetworkIPRouteDeleteParams struct {

	/* UUID.

	   Route UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip route delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRouteDeleteParams) WithDefaults() *NetworkIPRouteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip route delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRouteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) WithTimeout(timeout time.Duration) *NetworkIPRouteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) WithContext(ctx context.Context) *NetworkIPRouteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) WithHTTPClient(client *http.Client) *NetworkIPRouteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) WithUUID(uuid string) *NetworkIPRouteDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the network ip route delete params
func (o *NetworkIPRouteDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPRouteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
