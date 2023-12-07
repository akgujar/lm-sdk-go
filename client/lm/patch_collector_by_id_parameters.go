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

// NewPatchCollectorByIDParams creates a new PatchCollectorByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchCollectorByIDParams() *PatchCollectorByIDParams {
	return &PatchCollectorByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCollectorByIDParamsWithTimeout creates a new PatchCollectorByIDParams object
// with the ability to set a timeout on a request.
func NewPatchCollectorByIDParamsWithTimeout(timeout time.Duration) *PatchCollectorByIDParams {
	return &PatchCollectorByIDParams{
		timeout: timeout,
	}
}

// NewPatchCollectorByIDParamsWithContext creates a new PatchCollectorByIDParams object
// with the ability to set a context for a request.
func NewPatchCollectorByIDParamsWithContext(ctx context.Context) *PatchCollectorByIDParams {
	return &PatchCollectorByIDParams{
		Context: ctx,
	}
}

// NewPatchCollectorByIDParamsWithHTTPClient creates a new PatchCollectorByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchCollectorByIDParamsWithHTTPClient(client *http.Client) *PatchCollectorByIDParams {
	return &PatchCollectorByIDParams{
		HTTPClient: client,
	}
}

/* PatchCollectorByIDParams contains all the parameters to send to the API endpoint
   for the patch collector by Id operation.

   Typically these are written to a http.Request.
*/
type PatchCollectorByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// AutoBalanceMonitoredDevices.
	AutoBalanceMonitoredDevices *bool

	// Body.
	Body *models.Collector

	// ForceUpdateFailedOverDevices.
	ForceUpdateFailedOverDevices *bool

	// ID.
	//
	// Format: int32
	ID int32

	// OpType.
	//
	// Default: "refresh"
	OpType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCollectorByIDParams) WithDefaults() *PatchCollectorByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchCollectorByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		autoBalanceMonitoredDevicesDefault = bool(false)

		forceUpdateFailedOverDevicesDefault = bool(false)

		opTypeDefault = string("refresh")
	)

	val := PatchCollectorByIDParams{
		UserAgent:                    &userAgentDefault,
		AutoBalanceMonitoredDevices:  &autoBalanceMonitoredDevicesDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,
		OpType:                       &opTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithTimeout(timeout time.Duration) *PatchCollectorByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithContext(ctx context.Context) *PatchCollectorByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithHTTPClient(client *http.Client) *PatchCollectorByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithUserAgent(userAgent *string) *PatchCollectorByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAutoBalanceMonitoredDevices adds the autoBalanceMonitoredDevices to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithAutoBalanceMonitoredDevices(autoBalanceMonitoredDevices *bool) *PatchCollectorByIDParams {
	o.SetAutoBalanceMonitoredDevices(autoBalanceMonitoredDevices)
	return o
}

// SetAutoBalanceMonitoredDevices adds the autoBalanceMonitoredDevices to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetAutoBalanceMonitoredDevices(autoBalanceMonitoredDevices *bool) {
	o.AutoBalanceMonitoredDevices = autoBalanceMonitoredDevices
}

// WithBody adds the body to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithBody(body *models.Collector) *PatchCollectorByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetBody(body *models.Collector) {
	o.Body = body
}

// WithForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) *PatchCollectorByIDParams {
	o.SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices)
	return o
}

// SetForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) {
	o.ForceUpdateFailedOverDevices = forceUpdateFailedOverDevices
}

// WithID adds the id to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithID(id int32) *PatchCollectorByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOpType adds the opType to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithOpType(opType *string) *PatchCollectorByIDParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCollectorByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AutoBalanceMonitoredDevices != nil {

		// query param autoBalanceMonitoredDevices
		var qrAutoBalanceMonitoredDevices bool

		if o.AutoBalanceMonitoredDevices != nil {
			qrAutoBalanceMonitoredDevices = *o.AutoBalanceMonitoredDevices
		}
		qAutoBalanceMonitoredDevices := swag.FormatBool(qrAutoBalanceMonitoredDevices)
		if qAutoBalanceMonitoredDevices != "" {

			if err := r.SetQueryParam("autoBalanceMonitoredDevices", qAutoBalanceMonitoredDevices); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ForceUpdateFailedOverDevices != nil {

		// query param forceUpdateFailedOverDevices
		var qrForceUpdateFailedOverDevices bool

		if o.ForceUpdateFailedOverDevices != nil {
			qrForceUpdateFailedOverDevices = *o.ForceUpdateFailedOverDevices
		}
		qForceUpdateFailedOverDevices := swag.FormatBool(qrForceUpdateFailedOverDevices)
		if qForceUpdateFailedOverDevices != "" {

			if err := r.SetQueryParam("forceUpdateFailedOverDevices", qForceUpdateFailedOverDevices); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.OpType != nil {

		// query param opType
		var qrOpType string

		if o.OpType != nil {
			qrOpType = *o.OpType
		}
		qOpType := qrOpType
		if qOpType != "" {

			if err := r.SetQueryParam("opType", qOpType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
