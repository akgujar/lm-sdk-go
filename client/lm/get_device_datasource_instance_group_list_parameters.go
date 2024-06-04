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

// NewGetDeviceDatasourceInstanceGroupListParams creates a new GetDeviceDatasourceInstanceGroupListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceDatasourceInstanceGroupListParams() *GetDeviceDatasourceInstanceGroupListParams {
	return &GetDeviceDatasourceInstanceGroupListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupListParamsWithTimeout creates a new GetDeviceDatasourceInstanceGroupListParams object
// with the ability to set a timeout on a request.
func NewGetDeviceDatasourceInstanceGroupListParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupListParams {
	return &GetDeviceDatasourceInstanceGroupListParams{
		timeout: timeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupListParamsWithContext creates a new GetDeviceDatasourceInstanceGroupListParams object
// with the ability to set a context for a request.
func NewGetDeviceDatasourceInstanceGroupListParamsWithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupListParams {
	return &GetDeviceDatasourceInstanceGroupListParams{
		Context: ctx,
	}
}

// NewGetDeviceDatasourceInstanceGroupListParamsWithHTTPClient creates a new GetDeviceDatasourceInstanceGroupListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceDatasourceInstanceGroupListParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupListParams {
	return &GetDeviceDatasourceInstanceGroupListParams{
		HTTPClient: client,
	}
}

/*
GetDeviceDatasourceInstanceGroupListParams contains all the parameters to send to the API endpoint

	for the get device datasource instance group list operation.

	Typically these are written to a http.Request.
*/
type GetDeviceDatasourceInstanceGroupListParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	/* DeviceDsID.

	   The device-datasource ID you'd like to add an instance group for

	   Format: int32
	*/
	DeviceDsID int32

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

// WithDefaults hydrates default values in the get device datasource instance group list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDatasourceInstanceGroupListParams) WithDefaults() *GetDeviceDatasourceInstanceGroupListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device datasource instance group list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDatasourceInstanceGroupListParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetDeviceDatasourceInstanceGroupListParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithUserAgent(userAgent *string) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceDsID adds the deviceDsID to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithDeviceDsID(deviceDsID int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetDeviceDsID(deviceDsID)
	return o
}

// SetDeviceDsID adds the deviceDsId to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetDeviceDsID(deviceDsID int32) {
	o.DeviceDsID = deviceDsID
}

// WithDeviceID adds the deviceID to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithFields(fields *string) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithFilter(filter *string) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithOffset(offset *int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithSize(size *int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceInstanceGroupListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceDsId
	if err := r.SetPathParam("deviceDsId", swag.FormatInt32(o.DeviceDsID)); err != nil {
		return err
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
