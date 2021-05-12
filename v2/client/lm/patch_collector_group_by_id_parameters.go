// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// NewPatchCollectorGroupByIDParams creates a new PatchCollectorGroupByIDParams object
// with the default values initialized.
func NewPatchCollectorGroupByIDParams() *PatchCollectorGroupByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorGroupByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCollectorGroupByIDParamsWithTimeout creates a new PatchCollectorGroupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCollectorGroupByIDParamsWithTimeout(timeout time.Duration) *PatchCollectorGroupByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorGroupByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,

		timeout: timeout,
	}
}

// NewPatchCollectorGroupByIDParamsWithContext creates a new PatchCollectorGroupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCollectorGroupByIDParamsWithContext(ctx context.Context) *PatchCollectorGroupByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorGroupByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,

		Context: ctx,
	}
}

// NewPatchCollectorGroupByIDParamsWithHTTPClient creates a new PatchCollectorGroupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCollectorGroupByIDParamsWithHTTPClient(client *http.Client) *PatchCollectorGroupByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorGroupByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,
		HTTPClient:                   client,
	}
}

/*PatchCollectorGroupByIDParams contains all the parameters to send to the API endpoint
for the patch collector group by Id operation typically these are written to a http.Request
*/
type PatchCollectorGroupByIDParams struct {

	/*Body*/
	Body *models.CollectorGroup
	/*CollectorLoadBalanced*/
	CollectorLoadBalanced *bool
	/*ForceUpdateFailedOverDevices*/
	ForceUpdateFailedOverDevices *bool
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) WithTimeout(timeout time.Duration) *PatchCollectorGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) WithContext(ctx context.Context) *PatchCollectorGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) WithHTTPClient(client *http.Client) *PatchCollectorGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) WithBody(body *models.CollectorGroup) *PatchCollectorGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) SetBody(body *models.CollectorGroup) {
	o.Body = body
}

// WithCollectorLoadBalanced adds the collectorLoadBalanced to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) WithCollectorLoadBalanced(collectorLoadBalanced *bool) *PatchCollectorGroupByIDParams {
	o.SetCollectorLoadBalanced(collectorLoadBalanced)
	return o
}

// SetCollectorLoadBalanced adds the collectorLoadBalanced to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) SetCollectorLoadBalanced(collectorLoadBalanced *bool) {
	o.CollectorLoadBalanced = collectorLoadBalanced
}

// WithForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) WithForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) *PatchCollectorGroupByIDParams {
	o.SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices)
	return o
}

// SetForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) {
	o.ForceUpdateFailedOverDevices = forceUpdateFailedOverDevices
}

// WithID adds the id to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) WithID(id int32) *PatchCollectorGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch collector group by Id params
func (o *PatchCollectorGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCollectorGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CollectorLoadBalanced != nil {

		// query param collectorLoadBalanced
		var qrCollectorLoadBalanced bool
		if o.CollectorLoadBalanced != nil {
			qrCollectorLoadBalanced = *o.CollectorLoadBalanced
		}
		qCollectorLoadBalanced := swag.FormatBool(qrCollectorLoadBalanced)
		if qCollectorLoadBalanced != "" {
			if err := r.SetQueryParam("collectorLoadBalanced", qCollectorLoadBalanced); err != nil {
				return err
			}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}