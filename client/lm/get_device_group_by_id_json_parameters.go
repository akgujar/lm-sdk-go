// Code generated by go-swagger; DO NOT EDIT.

package lm

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

// NewGetDeviceGroupByIDJSONParams creates a new GetDeviceGroupByIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceGroupByIDJSONParams() *GetDeviceGroupByIDJSONParams {
	return &GetDeviceGroupByIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupByIDJSONParamsWithTimeout creates a new GetDeviceGroupByIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetDeviceGroupByIDJSONParamsWithTimeout(timeout time.Duration) *GetDeviceGroupByIDJSONParams {
	return &GetDeviceGroupByIDJSONParams{
		timeout: timeout,
	}
}

// NewGetDeviceGroupByIDJSONParamsWithContext creates a new GetDeviceGroupByIDJSONParams object
// with the ability to set a context for a request.
func NewGetDeviceGroupByIDJSONParamsWithContext(ctx context.Context) *GetDeviceGroupByIDJSONParams {
	return &GetDeviceGroupByIDJSONParams{
		Context: ctx,
	}
}

// NewGetDeviceGroupByIDJSONParamsWithHTTPClient creates a new GetDeviceGroupByIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceGroupByIDJSONParamsWithHTTPClient(client *http.Client) *GetDeviceGroupByIDJSONParams {
	return &GetDeviceGroupByIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetDeviceGroupByIDJSONParams contains all the parameters to send to the API endpoint

	for the get device group by Id Json operation.

	Typically these are written to a http.Request.
*/
type GetDeviceGroupByIDJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device group by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupByIDJSONParams) WithDefaults() *GetDeviceGroupByIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device group by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupByIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetDeviceGroupByIDJSONParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) WithTimeout(timeout time.Duration) *GetDeviceGroupByIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) WithContext(ctx context.Context) *GetDeviceGroupByIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) WithHTTPClient(client *http.Client) *GetDeviceGroupByIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) WithUserAgent(userAgent *string) *GetDeviceGroupByIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) WithFields(fields *string) *GetDeviceGroupByIDJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) WithID(id int32) *GetDeviceGroupByIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device group by Id Json params
func (o *GetDeviceGroupByIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupByIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
