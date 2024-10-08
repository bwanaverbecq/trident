// Code generated by go-swagger; DO NOT EDIT.

package storage

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
	"github.com/go-openapi/swag"
)

// NewStoragePoolGetParams creates a new StoragePoolGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStoragePoolGetParams() *StoragePoolGetParams {
	return &StoragePoolGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStoragePoolGetParamsWithTimeout creates a new StoragePoolGetParams object
// with the ability to set a timeout on a request.
func NewStoragePoolGetParamsWithTimeout(timeout time.Duration) *StoragePoolGetParams {
	return &StoragePoolGetParams{
		timeout: timeout,
	}
}

// NewStoragePoolGetParamsWithContext creates a new StoragePoolGetParams object
// with the ability to set a context for a request.
func NewStoragePoolGetParamsWithContext(ctx context.Context) *StoragePoolGetParams {
	return &StoragePoolGetParams{
		Context: ctx,
	}
}

// NewStoragePoolGetParamsWithHTTPClient creates a new StoragePoolGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStoragePoolGetParamsWithHTTPClient(client *http.Client) *StoragePoolGetParams {
	return &StoragePoolGetParams{
		HTTPClient: client,
	}
}

/*
StoragePoolGetParams contains all the parameters to send to the API endpoint

	for the storage pool get operation.

	Typically these are written to a http.Request.
*/
type StoragePoolGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Storage pool UUID.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage pool get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoragePoolGetParams) WithDefaults() *StoragePoolGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage pool get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoragePoolGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage pool get params
func (o *StoragePoolGetParams) WithTimeout(timeout time.Duration) *StoragePoolGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage pool get params
func (o *StoragePoolGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage pool get params
func (o *StoragePoolGetParams) WithContext(ctx context.Context) *StoragePoolGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage pool get params
func (o *StoragePoolGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage pool get params
func (o *StoragePoolGetParams) WithHTTPClient(client *http.Client) *StoragePoolGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage pool get params
func (o *StoragePoolGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the storage pool get params
func (o *StoragePoolGetParams) WithFields(fields []string) *StoragePoolGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the storage pool get params
func (o *StoragePoolGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the storage pool get params
func (o *StoragePoolGetParams) WithUUID(uuid string) *StoragePoolGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the storage pool get params
func (o *StoragePoolGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StoragePoolGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamStoragePoolGet binds the parameter fields
func (o *StoragePoolGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
