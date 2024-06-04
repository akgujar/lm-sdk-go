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
)

// NewGetAwsExternalIDParams creates a new GetAwsExternalIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAwsExternalIDParams() *GetAwsExternalIDParams {
	return &GetAwsExternalIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAwsExternalIDParamsWithTimeout creates a new GetAwsExternalIDParams object
// with the ability to set a timeout on a request.
func NewGetAwsExternalIDParamsWithTimeout(timeout time.Duration) *GetAwsExternalIDParams {
	return &GetAwsExternalIDParams{
		timeout: timeout,
	}
}

// NewGetAwsExternalIDParamsWithContext creates a new GetAwsExternalIDParams object
// with the ability to set a context for a request.
func NewGetAwsExternalIDParamsWithContext(ctx context.Context) *GetAwsExternalIDParams {
	return &GetAwsExternalIDParams{
		Context: ctx,
	}
}

// NewGetAwsExternalIDParamsWithHTTPClient creates a new GetAwsExternalIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAwsExternalIDParamsWithHTTPClient(client *http.Client) *GetAwsExternalIDParams {
	return &GetAwsExternalIDParams{
		HTTPClient: client,
	}
}

/*
GetAwsExternalIDParams contains all the parameters to send to the API endpoint

	for the get aws external Id operation.

	Typically these are written to a http.Request.
*/
type GetAwsExternalIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get aws external Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsExternalIDParams) WithDefaults() *GetAwsExternalIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get aws external Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsExternalIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetAwsExternalIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get aws external Id params
func (o *GetAwsExternalIDParams) WithTimeout(timeout time.Duration) *GetAwsExternalIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aws external Id params
func (o *GetAwsExternalIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aws external Id params
func (o *GetAwsExternalIDParams) WithContext(ctx context.Context) *GetAwsExternalIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aws external Id params
func (o *GetAwsExternalIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aws external Id params
func (o *GetAwsExternalIDParams) WithHTTPClient(client *http.Client) *GetAwsExternalIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aws external Id params
func (o *GetAwsExternalIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get aws external Id params
func (o *GetAwsExternalIDParams) WithUserAgent(userAgent *string) *GetAwsExternalIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get aws external Id params
func (o *GetAwsExternalIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetAwsExternalIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
