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

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// NewUpdateCollectorGroupByIDParams creates a new UpdateCollectorGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCollectorGroupByIDParams() *UpdateCollectorGroupByIDParams {
	return &UpdateCollectorGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCollectorGroupByIDParamsWithTimeout creates a new UpdateCollectorGroupByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateCollectorGroupByIDParamsWithTimeout(timeout time.Duration) *UpdateCollectorGroupByIDParams {
	return &UpdateCollectorGroupByIDParams{
		timeout: timeout,
	}
}

// NewUpdateCollectorGroupByIDParamsWithContext creates a new UpdateCollectorGroupByIDParams object
// with the ability to set a context for a request.
func NewUpdateCollectorGroupByIDParamsWithContext(ctx context.Context) *UpdateCollectorGroupByIDParams {
	return &UpdateCollectorGroupByIDParams{
		Context: ctx,
	}
}

// NewUpdateCollectorGroupByIDParamsWithHTTPClient creates a new UpdateCollectorGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCollectorGroupByIDParamsWithHTTPClient(client *http.Client) *UpdateCollectorGroupByIDParams {
	return &UpdateCollectorGroupByIDParams{
		HTTPClient: client,
	}
}

/* UpdateCollectorGroupByIDParams contains all the parameters to send to the API endpoint
   for the update collector group by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateCollectorGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// AutoBalanceMonitoredDevices.
	AutoBalanceMonitoredDevices *bool

	// Body.
	Body *models.CollectorGroup

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

// WithDefaults hydrates default values in the update collector group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCollectorGroupByIDParams) WithDefaults() *UpdateCollectorGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update collector group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCollectorGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		autoBalanceMonitoredDevicesDefault = bool(false)

		forceUpdateFailedOverDevicesDefault = bool(false)

		opTypeDefault = string("refresh")
	)

	val := UpdateCollectorGroupByIDParams{
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

// WithTimeout adds the timeout to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithTimeout(timeout time.Duration) *UpdateCollectorGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithContext(ctx context.Context) *UpdateCollectorGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithHTTPClient(client *http.Client) *UpdateCollectorGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithUserAgent(userAgent *string) *UpdateCollectorGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAutoBalanceMonitoredDevices adds the autoBalanceMonitoredDevices to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithAutoBalanceMonitoredDevices(autoBalanceMonitoredDevices *bool) *UpdateCollectorGroupByIDParams {
	o.SetAutoBalanceMonitoredDevices(autoBalanceMonitoredDevices)
	return o
}

// SetAutoBalanceMonitoredDevices adds the autoBalanceMonitoredDevices to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetAutoBalanceMonitoredDevices(autoBalanceMonitoredDevices *bool) {
	o.AutoBalanceMonitoredDevices = autoBalanceMonitoredDevices
}

// WithBody adds the body to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithBody(body *models.CollectorGroup) *UpdateCollectorGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetBody(body *models.CollectorGroup) {
	o.Body = body
}

// WithForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) *UpdateCollectorGroupByIDParams {
	o.SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices)
	return o
}

// SetForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) {
	o.ForceUpdateFailedOverDevices = forceUpdateFailedOverDevices
}

// WithID adds the id to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithID(id int32) *UpdateCollectorGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOpType adds the opType to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) WithOpType(opType *string) *UpdateCollectorGroupByIDParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the update collector group by Id params
func (o *UpdateCollectorGroupByIDParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCollectorGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
