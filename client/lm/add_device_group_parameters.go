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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddDeviceGroupParams creates a new AddDeviceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDeviceGroupParams() *AddDeviceGroupParams {
	return &AddDeviceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceGroupParamsWithTimeout creates a new AddDeviceGroupParams object
// with the ability to set a timeout on a request.
func NewAddDeviceGroupParamsWithTimeout(timeout time.Duration) *AddDeviceGroupParams {
	return &AddDeviceGroupParams{
		timeout: timeout,
	}
}

// NewAddDeviceGroupParamsWithContext creates a new AddDeviceGroupParams object
// with the ability to set a context for a request.
func NewAddDeviceGroupParamsWithContext(ctx context.Context) *AddDeviceGroupParams {
	return &AddDeviceGroupParams{
		Context: ctx,
	}
}

// NewAddDeviceGroupParamsWithHTTPClient creates a new AddDeviceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDeviceGroupParamsWithHTTPClient(client *http.Client) *AddDeviceGroupParams {
	return &AddDeviceGroupParams{
		HTTPClient: client,
	}
}

/* AddDeviceGroupParams contains all the parameters to send to the API endpoint
   for the add device group operation.

   Typically these are written to a http.Request.
*/
type AddDeviceGroupParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// Body.
	Body *models.DeviceGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add device group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceGroupParams) WithDefaults() *AddDeviceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add device group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDeviceGroupParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")
	)

	val := AddDeviceGroupParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add device group params
func (o *AddDeviceGroupParams) WithTimeout(timeout time.Duration) *AddDeviceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device group params
func (o *AddDeviceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device group params
func (o *AddDeviceGroupParams) WithContext(ctx context.Context) *AddDeviceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device group params
func (o *AddDeviceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device group params
func (o *AddDeviceGroupParams) WithHTTPClient(client *http.Client) *AddDeviceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device group params
func (o *AddDeviceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add device group params
func (o *AddDeviceGroupParams) WithUserAgent(userAgent *string) *AddDeviceGroupParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add device group params
func (o *AddDeviceGroupParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add device group params
func (o *AddDeviceGroupParams) WithBody(body *models.DeviceGroup) *AddDeviceGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device group params
func (o *AddDeviceGroupParams) SetBody(body *models.DeviceGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
