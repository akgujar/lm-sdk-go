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

// NewAddConfigsourceAuditVersionParams creates a new AddConfigsourceAuditVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddConfigsourceAuditVersionParams() *AddConfigsourceAuditVersionParams {
	return &AddConfigsourceAuditVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddConfigsourceAuditVersionParamsWithTimeout creates a new AddConfigsourceAuditVersionParams object
// with the ability to set a timeout on a request.
func NewAddConfigsourceAuditVersionParamsWithTimeout(timeout time.Duration) *AddConfigsourceAuditVersionParams {
	return &AddConfigsourceAuditVersionParams{
		timeout: timeout,
	}
}

// NewAddConfigsourceAuditVersionParamsWithContext creates a new AddConfigsourceAuditVersionParams object
// with the ability to set a context for a request.
func NewAddConfigsourceAuditVersionParamsWithContext(ctx context.Context) *AddConfigsourceAuditVersionParams {
	return &AddConfigsourceAuditVersionParams{
		Context: ctx,
	}
}

// NewAddConfigsourceAuditVersionParamsWithHTTPClient creates a new AddConfigsourceAuditVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddConfigsourceAuditVersionParamsWithHTTPClient(client *http.Client) *AddConfigsourceAuditVersionParams {
	return &AddConfigsourceAuditVersionParams{
		HTTPClient: client,
	}
}

/* AddConfigsourceAuditVersionParams contains all the parameters to send to the API endpoint
   for the add configsource audit version operation.

   Typically these are written to a http.Request.
*/
type AddConfigsourceAuditVersionParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.Audit

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add configsource audit version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddConfigsourceAuditVersionParams) WithDefaults() *AddConfigsourceAuditVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add configsource audit version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddConfigsourceAuditVersionParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := AddConfigsourceAuditVersionParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) WithTimeout(timeout time.Duration) *AddConfigsourceAuditVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) WithContext(ctx context.Context) *AddConfigsourceAuditVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) WithHTTPClient(client *http.Client) *AddConfigsourceAuditVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) WithUserAgent(userAgent *string) *AddConfigsourceAuditVersionParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) WithBody(body *models.Audit) *AddConfigsourceAuditVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) SetBody(body *models.Audit) {
	o.Body = body
}

// WithID adds the id to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) WithID(id int32) *AddConfigsourceAuditVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add configsource audit version params
func (o *AddConfigsourceAuditVersionParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddConfigsourceAuditVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
