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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchDevicePropertyByNameParams creates a new PatchDevicePropertyByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDevicePropertyByNameParams() *PatchDevicePropertyByNameParams {
	return &PatchDevicePropertyByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDevicePropertyByNameParamsWithTimeout creates a new PatchDevicePropertyByNameParams object
// with the ability to set a timeout on a request.
func NewPatchDevicePropertyByNameParamsWithTimeout(timeout time.Duration) *PatchDevicePropertyByNameParams {
	return &PatchDevicePropertyByNameParams{
		timeout: timeout,
	}
}

// NewPatchDevicePropertyByNameParamsWithContext creates a new PatchDevicePropertyByNameParams object
// with the ability to set a context for a request.
func NewPatchDevicePropertyByNameParamsWithContext(ctx context.Context) *PatchDevicePropertyByNameParams {
	return &PatchDevicePropertyByNameParams{
		Context: ctx,
	}
}

// NewPatchDevicePropertyByNameParamsWithHTTPClient creates a new PatchDevicePropertyByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDevicePropertyByNameParamsWithHTTPClient(client *http.Client) *PatchDevicePropertyByNameParams {
	return &PatchDevicePropertyByNameParams{
		HTTPClient: client,
	}
}

/* PatchDevicePropertyByNameParams contains all the parameters to send to the API endpoint
   for the patch device property by name operation.

   Typically these are written to a http.Request.
*/
type PatchDevicePropertyByNameParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.EntityProperty

	// DeviceID.
	//
	// Format: int32
	DeviceID int32

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch device property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDevicePropertyByNameParams) WithDefaults() *PatchDevicePropertyByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch device property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDevicePropertyByNameParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := PatchDevicePropertyByNameParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) WithTimeout(timeout time.Duration) *PatchDevicePropertyByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) WithContext(ctx context.Context) *PatchDevicePropertyByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) WithHTTPClient(client *http.Client) *PatchDevicePropertyByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) WithUserAgent(userAgent *string) *PatchDevicePropertyByNameParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) WithBody(body *models.EntityProperty) *PatchDevicePropertyByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) SetBody(body *models.EntityProperty) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) WithDeviceID(deviceID int32) *PatchDevicePropertyByNameParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithName adds the name to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) WithName(name string) *PatchDevicePropertyByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch device property by name params
func (o *PatchDevicePropertyByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDevicePropertyByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
