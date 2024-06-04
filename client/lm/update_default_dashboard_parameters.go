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

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// NewUpdateDefaultDashboardParams creates a new UpdateDefaultDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDefaultDashboardParams() *UpdateDefaultDashboardParams {
	return &UpdateDefaultDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDefaultDashboardParamsWithTimeout creates a new UpdateDefaultDashboardParams object
// with the ability to set a timeout on a request.
func NewUpdateDefaultDashboardParamsWithTimeout(timeout time.Duration) *UpdateDefaultDashboardParams {
	return &UpdateDefaultDashboardParams{
		timeout: timeout,
	}
}

// NewUpdateDefaultDashboardParamsWithContext creates a new UpdateDefaultDashboardParams object
// with the ability to set a context for a request.
func NewUpdateDefaultDashboardParamsWithContext(ctx context.Context) *UpdateDefaultDashboardParams {
	return &UpdateDefaultDashboardParams{
		Context: ctx,
	}
}

// NewUpdateDefaultDashboardParamsWithHTTPClient creates a new UpdateDefaultDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDefaultDashboardParamsWithHTTPClient(client *http.Client) *UpdateDefaultDashboardParams {
	return &UpdateDefaultDashboardParams{
		HTTPClient: client,
	}
}

/*
UpdateDefaultDashboardParams contains all the parameters to send to the API endpoint

	for the update default dashboard operation.

	Typically these are written to a http.Request.
*/
type UpdateDefaultDashboardParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.RestUserCustomizedDataV3

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update default dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDefaultDashboardParams) WithDefaults() *UpdateDefaultDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update default dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDefaultDashboardParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := UpdateDefaultDashboardParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update default dashboard params
func (o *UpdateDefaultDashboardParams) WithTimeout(timeout time.Duration) *UpdateDefaultDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update default dashboard params
func (o *UpdateDefaultDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update default dashboard params
func (o *UpdateDefaultDashboardParams) WithContext(ctx context.Context) *UpdateDefaultDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update default dashboard params
func (o *UpdateDefaultDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update default dashboard params
func (o *UpdateDefaultDashboardParams) WithHTTPClient(client *http.Client) *UpdateDefaultDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update default dashboard params
func (o *UpdateDefaultDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update default dashboard params
func (o *UpdateDefaultDashboardParams) WithUserAgent(userAgent *string) *UpdateDefaultDashboardParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update default dashboard params
func (o *UpdateDefaultDashboardParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update default dashboard params
func (o *UpdateDefaultDashboardParams) WithBody(body *models.RestUserCustomizedDataV3) *UpdateDefaultDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update default dashboard params
func (o *UpdateDefaultDashboardParams) SetBody(body *models.RestUserCustomizedDataV3) {
	o.Body = body
}

// WithID adds the id to the update default dashboard params
func (o *UpdateDefaultDashboardParams) WithID(id string) *UpdateDefaultDashboardParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update default dashboard params
func (o *UpdateDefaultDashboardParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDefaultDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
