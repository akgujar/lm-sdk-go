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

// NewGetDeviceGroupListJSONParams creates a new GetDeviceGroupListJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceGroupListJSONParams() *GetDeviceGroupListJSONParams {
	return &GetDeviceGroupListJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupListJSONParamsWithTimeout creates a new GetDeviceGroupListJSONParams object
// with the ability to set a timeout on a request.
func NewGetDeviceGroupListJSONParamsWithTimeout(timeout time.Duration) *GetDeviceGroupListJSONParams {
	return &GetDeviceGroupListJSONParams{
		timeout: timeout,
	}
}

// NewGetDeviceGroupListJSONParamsWithContext creates a new GetDeviceGroupListJSONParams object
// with the ability to set a context for a request.
func NewGetDeviceGroupListJSONParamsWithContext(ctx context.Context) *GetDeviceGroupListJSONParams {
	return &GetDeviceGroupListJSONParams{
		Context: ctx,
	}
}

// NewGetDeviceGroupListJSONParamsWithHTTPClient creates a new GetDeviceGroupListJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceGroupListJSONParamsWithHTTPClient(client *http.Client) *GetDeviceGroupListJSONParams {
	return &GetDeviceGroupListJSONParams{
		HTTPClient: client,
	}
}

/*
GetDeviceGroupListJSONParams contains all the parameters to send to the API endpoint

	for the get device group list Json operation.

	Typically these are written to a http.Request.
*/
type GetDeviceGroupListJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

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

// WithDefaults hydrates default values in the get device group list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupListJSONParams) WithDefaults() *GetDeviceGroupListJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device group list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupListJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetDeviceGroupListJSONParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithTimeout(timeout time.Duration) *GetDeviceGroupListJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithContext(ctx context.Context) *GetDeviceGroupListJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithHTTPClient(client *http.Client) *GetDeviceGroupListJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithUserAgent(userAgent *string) *GetDeviceGroupListJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithFields(fields *string) *GetDeviceGroupListJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithFilter(filter *string) *GetDeviceGroupListJSONParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithOffset(offset *int32) *GetDeviceGroupListJSONParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) WithSize(size *int32) *GetDeviceGroupListJSONParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device group list Json params
func (o *GetDeviceGroupListJSONParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupListJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
