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

// NewAddOpsNoteParams creates a new AddOpsNoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddOpsNoteParams() *AddOpsNoteParams {
	return &AddOpsNoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddOpsNoteParamsWithTimeout creates a new AddOpsNoteParams object
// with the ability to set a timeout on a request.
func NewAddOpsNoteParamsWithTimeout(timeout time.Duration) *AddOpsNoteParams {
	return &AddOpsNoteParams{
		timeout: timeout,
	}
}

// NewAddOpsNoteParamsWithContext creates a new AddOpsNoteParams object
// with the ability to set a context for a request.
func NewAddOpsNoteParamsWithContext(ctx context.Context) *AddOpsNoteParams {
	return &AddOpsNoteParams{
		Context: ctx,
	}
}

// NewAddOpsNoteParamsWithHTTPClient creates a new AddOpsNoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddOpsNoteParamsWithHTTPClient(client *http.Client) *AddOpsNoteParams {
	return &AddOpsNoteParams{
		HTTPClient: client,
	}
}

/*
AddOpsNoteParams contains all the parameters to send to the API endpoint

	for the add ops note operation.

	Typically these are written to a http.Request.
*/
type AddOpsNoteParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.OpsNote

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add ops note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOpsNoteParams) WithDefaults() *AddOpsNoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add ops note params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddOpsNoteParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := AddOpsNoteParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the add ops note params
func (o *AddOpsNoteParams) WithTimeout(timeout time.Duration) *AddOpsNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ops note params
func (o *AddOpsNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ops note params
func (o *AddOpsNoteParams) WithContext(ctx context.Context) *AddOpsNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ops note params
func (o *AddOpsNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ops note params
func (o *AddOpsNoteParams) WithHTTPClient(client *http.Client) *AddOpsNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ops note params
func (o *AddOpsNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the add ops note params
func (o *AddOpsNoteParams) WithUserAgent(userAgent *string) *AddOpsNoteParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the add ops note params
func (o *AddOpsNoteParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the add ops note params
func (o *AddOpsNoteParams) WithBody(body *models.OpsNote) *AddOpsNoteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add ops note params
func (o *AddOpsNoteParams) SetBody(body *models.OpsNote) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddOpsNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
