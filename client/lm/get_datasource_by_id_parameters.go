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

// NewGetDatasourceByIDParams creates a new GetDatasourceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatasourceByIDParams() *GetDatasourceByIDParams {
	return &GetDatasourceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasourceByIDParamsWithTimeout creates a new GetDatasourceByIDParams object
// with the ability to set a timeout on a request.
func NewGetDatasourceByIDParamsWithTimeout(timeout time.Duration) *GetDatasourceByIDParams {
	return &GetDatasourceByIDParams{
		timeout: timeout,
	}
}

// NewGetDatasourceByIDParamsWithContext creates a new GetDatasourceByIDParams object
// with the ability to set a context for a request.
func NewGetDatasourceByIDParamsWithContext(ctx context.Context) *GetDatasourceByIDParams {
	return &GetDatasourceByIDParams{
		Context: ctx,
	}
}

// NewGetDatasourceByIDParamsWithHTTPClient creates a new GetDatasourceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatasourceByIDParamsWithHTTPClient(client *http.Client) *GetDatasourceByIDParams {
	return &GetDatasourceByIDParams{
		HTTPClient: client,
	}
}

/*
GetDatasourceByIDParams contains all the parameters to send to the API endpoint

	for the get datasource by Id operation.

	Typically these are written to a http.Request.
*/
type GetDatasourceByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// Format.
	//
	// Default: "json"
	Format *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get datasource by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasourceByIDParams) WithDefaults() *GetDatasourceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get datasource by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasourceByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		formatDefault = string("json")
	)

	val := GetDatasourceByIDParams{
		UserAgent: &userAgentDefault,
		Format:    &formatDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithTimeout(timeout time.Duration) *GetDatasourceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithContext(ctx context.Context) *GetDatasourceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithHTTPClient(client *http.Client) *GetDatasourceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithUserAgent(userAgent *string) *GetDatasourceByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithFields(fields *string) *GetDatasourceByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFormat adds the format to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithFormat(format *string) *GetDatasourceByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get datasource by Id params
func (o *GetDatasourceByIDParams) WithID(id int32) *GetDatasourceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get datasource by Id params
func (o *GetDatasourceByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasourceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
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
