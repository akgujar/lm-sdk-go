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

// NewGetImmediateDeviceListByDeviceGroupIDJSONParams creates a new GetImmediateDeviceListByDeviceGroupIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetImmediateDeviceListByDeviceGroupIDJSONParams() *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	return &GetImmediateDeviceListByDeviceGroupIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetImmediateDeviceListByDeviceGroupIDJSONParamsWithTimeout creates a new GetImmediateDeviceListByDeviceGroupIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetImmediateDeviceListByDeviceGroupIDJSONParamsWithTimeout(timeout time.Duration) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	return &GetImmediateDeviceListByDeviceGroupIDJSONParams{
		timeout: timeout,
	}
}

// NewGetImmediateDeviceListByDeviceGroupIDJSONParamsWithContext creates a new GetImmediateDeviceListByDeviceGroupIDJSONParams object
// with the ability to set a context for a request.
func NewGetImmediateDeviceListByDeviceGroupIDJSONParamsWithContext(ctx context.Context) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	return &GetImmediateDeviceListByDeviceGroupIDJSONParams{
		Context: ctx,
	}
}

// NewGetImmediateDeviceListByDeviceGroupIDJSONParamsWithHTTPClient creates a new GetImmediateDeviceListByDeviceGroupIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetImmediateDeviceListByDeviceGroupIDJSONParamsWithHTTPClient(client *http.Client) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	return &GetImmediateDeviceListByDeviceGroupIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetImmediateDeviceListByDeviceGroupIDJSONParams contains all the parameters to send to the API endpoint

	for the get immediate device list by device group Id Json operation.

	Typically these are written to a http.Request.
*/
type GetImmediateDeviceListByDeviceGroupIDJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

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

// WithDefaults hydrates default values in the get immediate device list by device group Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithDefaults() *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get immediate device list by device group Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetImmediateDeviceListByDeviceGroupIDJSONParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithTimeout(timeout time.Duration) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithContext(ctx context.Context) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithHTTPClient(client *http.Client) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithUserAgent(userAgent *string) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithFields(fields *string) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithFilter(filter *string) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithID(id int32) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithOffset(offset *int32) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WithSize(size *int32) *GetImmediateDeviceListByDeviceGroupIDJSONParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get immediate device list by device group Id Json params
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetImmediateDeviceListByDeviceGroupIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
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
