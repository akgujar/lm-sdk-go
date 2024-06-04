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

// NewGetDevicePropertyListJSONParams creates a new GetDevicePropertyListJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDevicePropertyListJSONParams() *GetDevicePropertyListJSONParams {
	return &GetDevicePropertyListJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicePropertyListJSONParamsWithTimeout creates a new GetDevicePropertyListJSONParams object
// with the ability to set a timeout on a request.
func NewGetDevicePropertyListJSONParamsWithTimeout(timeout time.Duration) *GetDevicePropertyListJSONParams {
	return &GetDevicePropertyListJSONParams{
		timeout: timeout,
	}
}

// NewGetDevicePropertyListJSONParamsWithContext creates a new GetDevicePropertyListJSONParams object
// with the ability to set a context for a request.
func NewGetDevicePropertyListJSONParamsWithContext(ctx context.Context) *GetDevicePropertyListJSONParams {
	return &GetDevicePropertyListJSONParams{
		Context: ctx,
	}
}

// NewGetDevicePropertyListJSONParamsWithHTTPClient creates a new GetDevicePropertyListJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDevicePropertyListJSONParamsWithHTTPClient(client *http.Client) *GetDevicePropertyListJSONParams {
	return &GetDevicePropertyListJSONParams{
		HTTPClient: client,
	}
}

/*
GetDevicePropertyListJSONParams contains all the parameters to send to the API endpoint

	for the get device property list Json operation.

	Typically these are written to a http.Request.
*/
type GetDevicePropertyListJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

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

// WithDefaults hydrates default values in the get device property list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicePropertyListJSONParams) WithDefaults() *GetDevicePropertyListJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device property list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDevicePropertyListJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetDevicePropertyListJSONParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithTimeout(timeout time.Duration) *GetDevicePropertyListJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithContext(ctx context.Context) *GetDevicePropertyListJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithHTTPClient(client *http.Client) *GetDevicePropertyListJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithUserAgent(userAgent *string) *GetDevicePropertyListJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceID adds the deviceID to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithDeviceID(deviceID int32) *GetDevicePropertyListJSONParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithFields(fields *string) *GetDevicePropertyListJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithFilter(filter *string) *GetDevicePropertyListJSONParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithOffset(offset *int32) *GetDevicePropertyListJSONParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) WithSize(size *int32) *GetDevicePropertyListJSONParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device property list Json params
func (o *GetDevicePropertyListJSONParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicePropertyListJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
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
