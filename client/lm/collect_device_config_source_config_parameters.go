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

// NewCollectDeviceConfigSourceConfigParams creates a new CollectDeviceConfigSourceConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCollectDeviceConfigSourceConfigParams() *CollectDeviceConfigSourceConfigParams {
	return &CollectDeviceConfigSourceConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCollectDeviceConfigSourceConfigParamsWithTimeout creates a new CollectDeviceConfigSourceConfigParams object
// with the ability to set a timeout on a request.
func NewCollectDeviceConfigSourceConfigParamsWithTimeout(timeout time.Duration) *CollectDeviceConfigSourceConfigParams {
	return &CollectDeviceConfigSourceConfigParams{
		timeout: timeout,
	}
}

// NewCollectDeviceConfigSourceConfigParamsWithContext creates a new CollectDeviceConfigSourceConfigParams object
// with the ability to set a context for a request.
func NewCollectDeviceConfigSourceConfigParamsWithContext(ctx context.Context) *CollectDeviceConfigSourceConfigParams {
	return &CollectDeviceConfigSourceConfigParams{
		Context: ctx,
	}
}

// NewCollectDeviceConfigSourceConfigParamsWithHTTPClient creates a new CollectDeviceConfigSourceConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCollectDeviceConfigSourceConfigParamsWithHTTPClient(client *http.Client) *CollectDeviceConfigSourceConfigParams {
	return &CollectDeviceConfigSourceConfigParams{
		HTTPClient: client,
	}
}

/*
CollectDeviceConfigSourceConfigParams contains all the parameters to send to the API endpoint

	for the collect device config source config operation.

	Typically these are written to a http.Request.
*/
type CollectDeviceConfigSourceConfigParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	// HdsID.
	//
	// Format: int32
	HdsID int32

	// InstanceID.
	//
	// Format: int32
	InstanceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the collect device config source config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectDeviceConfigSourceConfigParams) WithDefaults() *CollectDeviceConfigSourceConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the collect device config source config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectDeviceConfigSourceConfigParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := CollectDeviceConfigSourceConfigParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) WithTimeout(timeout time.Duration) *CollectDeviceConfigSourceConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) WithContext(ctx context.Context) *CollectDeviceConfigSourceConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) WithHTTPClient(client *http.Client) *CollectDeviceConfigSourceConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) WithUserAgent(userAgent *string) *CollectDeviceConfigSourceConfigParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceID adds the deviceID to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) WithDeviceID(deviceID int32) *CollectDeviceConfigSourceConfigParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithHdsID adds the hdsID to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) WithHdsID(hdsID int32) *CollectDeviceConfigSourceConfigParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithInstanceID adds the instanceID to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) WithInstanceID(instanceID int32) *CollectDeviceConfigSourceConfigParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the collect device config source config params
func (o *CollectDeviceConfigSourceConfigParams) SetInstanceID(instanceID int32) {
	o.InstanceID = instanceID
}

// WriteToRequest writes these params to a swagger request
func (o *CollectDeviceConfigSourceConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
	}

	// path param instanceId
	if err := r.SetPathParam("instanceId", swag.FormatInt32(o.InstanceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
