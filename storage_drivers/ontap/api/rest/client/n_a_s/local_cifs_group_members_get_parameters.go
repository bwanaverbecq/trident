// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewLocalCifsGroupMembersGetParams creates a new LocalCifsGroupMembersGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalCifsGroupMembersGetParams() *LocalCifsGroupMembersGetParams {
	return &LocalCifsGroupMembersGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalCifsGroupMembersGetParamsWithTimeout creates a new LocalCifsGroupMembersGetParams object
// with the ability to set a timeout on a request.
func NewLocalCifsGroupMembersGetParamsWithTimeout(timeout time.Duration) *LocalCifsGroupMembersGetParams {
	return &LocalCifsGroupMembersGetParams{
		timeout: timeout,
	}
}

// NewLocalCifsGroupMembersGetParamsWithContext creates a new LocalCifsGroupMembersGetParams object
// with the ability to set a context for a request.
func NewLocalCifsGroupMembersGetParamsWithContext(ctx context.Context) *LocalCifsGroupMembersGetParams {
	return &LocalCifsGroupMembersGetParams{
		Context: ctx,
	}
}

// NewLocalCifsGroupMembersGetParamsWithHTTPClient creates a new LocalCifsGroupMembersGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalCifsGroupMembersGetParamsWithHTTPClient(client *http.Client) *LocalCifsGroupMembersGetParams {
	return &LocalCifsGroupMembersGetParams{
		HTTPClient: client,
	}
}

/*
LocalCifsGroupMembersGetParams contains all the parameters to send to the API endpoint

	for the local cifs group members get operation.

	Typically these are written to a http.Request.
*/
type LocalCifsGroupMembersGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LocalCifsGroupSid.

	   Local group SID
	*/
	LocalCifsGroupSid string

	/* Name.

	   Member name
	*/
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local cifs group members get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupMembersGetParams) WithDefaults() *LocalCifsGroupMembersGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local cifs group members get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupMembersGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) WithTimeout(timeout time.Duration) *LocalCifsGroupMembersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) WithContext(ctx context.Context) *LocalCifsGroupMembersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) WithHTTPClient(client *http.Client) *LocalCifsGroupMembersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) WithFields(fields []string) *LocalCifsGroupMembersGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLocalCifsGroupSid adds the localCifsGroupSid to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) WithLocalCifsGroupSid(localCifsGroupSid string) *LocalCifsGroupMembersGetParams {
	o.SetLocalCifsGroupSid(localCifsGroupSid)
	return o
}

// SetLocalCifsGroupSid adds the localCifsGroupSid to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) SetLocalCifsGroupSid(localCifsGroupSid string) {
	o.LocalCifsGroupSid = localCifsGroupSid
}

// WithName adds the name to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) WithName(name string) *LocalCifsGroupMembersGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) WithSvmUUID(svmUUID string) *LocalCifsGroupMembersGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the local cifs group members get params
func (o *LocalCifsGroupMembersGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LocalCifsGroupMembersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param local_cifs_group.sid
	if err := r.SetPathParam("local_cifs_group.sid", o.LocalCifsGroupSid); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLocalCifsGroupMembersGet binds the parameter fields
func (o *LocalCifsGroupMembersGetParams) bindParamFields(formats strfmt.Registry) []string {
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
