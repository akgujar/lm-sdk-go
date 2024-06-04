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

// NewGetDeviceDatasourceInstanceByIDJSONParams creates a new GetDeviceDatasourceInstanceByIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceDatasourceInstanceByIDJSONParams() *GetDeviceDatasourceInstanceByIDJSONParams {
	return &GetDeviceDatasourceInstanceByIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceInstanceByIDJSONParamsWithTimeout creates a new GetDeviceDatasourceInstanceByIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetDeviceDatasourceInstanceByIDJSONParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceByIDJSONParams {
	return &GetDeviceDatasourceInstanceByIDJSONParams{
		timeout: timeout,
	}
}

// NewGetDeviceDatasourceInstanceByIDJSONParamsWithContext creates a new GetDeviceDatasourceInstanceByIDJSONParams object
// with the ability to set a context for a request.
func NewGetDeviceDatasourceInstanceByIDJSONParamsWithContext(ctx context.Context) *GetDeviceDatasourceInstanceByIDJSONParams {
	return &GetDeviceDatasourceInstanceByIDJSONParams{
		Context: ctx,
	}
}

// NewGetDeviceDatasourceInstanceByIDJSONParamsWithHTTPClient creates a new GetDeviceDatasourceInstanceByIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceDatasourceInstanceByIDJSONParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceByIDJSONParams {
	return &GetDeviceDatasourceInstanceByIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetDeviceDatasourceInstanceByIDJSONParams contains all the parameters to send to the API endpoint

	for the get device datasource instance by Id Json operation.

	Typically these are written to a http.Request.
*/
type GetDeviceDatasourceInstanceByIDJSONParams struct {

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

	/* HdsID.

	   The device-datasource ID

	   Format: int32
	*/
	HdsID int32

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device datasource instance by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithDefaults() *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device datasource instance by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetDeviceDatasourceInstanceByIDJSONParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithContext(ctx context.Context) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithUserAgent(userAgent *string) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceID adds the deviceID to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithFields(fields *string) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHdsID adds the hdsID to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithHdsID(hdsID int32) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WithID(id int32) *GetDeviceDatasourceInstanceByIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device datasource instance by Id Json params
func (o *GetDeviceDatasourceInstanceByIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceInstanceByIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
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