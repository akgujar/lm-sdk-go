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

// NewGetDatasourceListJSONParams creates a new GetDatasourceListJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatasourceListJSONParams() *GetDatasourceListJSONParams {
	return &GetDatasourceListJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasourceListJSONParamsWithTimeout creates a new GetDatasourceListJSONParams object
// with the ability to set a timeout on a request.
func NewGetDatasourceListJSONParamsWithTimeout(timeout time.Duration) *GetDatasourceListJSONParams {
	return &GetDatasourceListJSONParams{
		timeout: timeout,
	}
}

// NewGetDatasourceListJSONParamsWithContext creates a new GetDatasourceListJSONParams object
// with the ability to set a context for a request.
func NewGetDatasourceListJSONParamsWithContext(ctx context.Context) *GetDatasourceListJSONParams {
	return &GetDatasourceListJSONParams{
		Context: ctx,
	}
}

// NewGetDatasourceListJSONParamsWithHTTPClient creates a new GetDatasourceListJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatasourceListJSONParamsWithHTTPClient(client *http.Client) *GetDatasourceListJSONParams {
	return &GetDatasourceListJSONParams{
		HTTPClient: client,
	}
}

/*
GetDatasourceListJSONParams contains all the parameters to send to the API endpoint

	for the get datasource list Json operation.

	Typically these are written to a http.Request.
*/
type GetDatasourceListJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// Format.
	//
	// Default: "json"
	Format *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get datasource list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasourceListJSONParams) WithDefaults() *GetDatasourceListJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get datasource list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasourceListJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		formatDefault = string("json")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetDatasourceListJSONParams{
		UserAgent: &userAgentDefault,
		Format:    &formatDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithTimeout(timeout time.Duration) *GetDatasourceListJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithContext(ctx context.Context) *GetDatasourceListJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithHTTPClient(client *http.Client) *GetDatasourceListJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithUserAgent(userAgent *string) *GetDatasourceListJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithFields(fields *string) *GetDatasourceListJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithFilter(filter *string) *GetDatasourceListJSONParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithFormat adds the format to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithFormat(format *string) *GetDatasourceListJSONParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetFormat(format *string) {
	o.Format = format
}

// WithOffset adds the offset to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithOffset(offset *int32) *GetDatasourceListJSONParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get datasource list Json params
func (o *GetDatasourceListJSONParams) WithSize(size *int32) *GetDatasourceListJSONParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get datasource list Json params
func (o *GetDatasourceListJSONParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasourceListJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
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

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
