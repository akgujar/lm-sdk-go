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

// NewGetDebugCommandResultJSONParams creates a new GetDebugCommandResultJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDebugCommandResultJSONParams() *GetDebugCommandResultJSONParams {
	return &GetDebugCommandResultJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDebugCommandResultJSONParamsWithTimeout creates a new GetDebugCommandResultJSONParams object
// with the ability to set a timeout on a request.
func NewGetDebugCommandResultJSONParamsWithTimeout(timeout time.Duration) *GetDebugCommandResultJSONParams {
	return &GetDebugCommandResultJSONParams{
		timeout: timeout,
	}
}

// NewGetDebugCommandResultJSONParamsWithContext creates a new GetDebugCommandResultJSONParams object
// with the ability to set a context for a request.
func NewGetDebugCommandResultJSONParamsWithContext(ctx context.Context) *GetDebugCommandResultJSONParams {
	return &GetDebugCommandResultJSONParams{
		Context: ctx,
	}
}

// NewGetDebugCommandResultJSONParamsWithHTTPClient creates a new GetDebugCommandResultJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDebugCommandResultJSONParamsWithHTTPClient(client *http.Client) *GetDebugCommandResultJSONParams {
	return &GetDebugCommandResultJSONParams{
		HTTPClient: client,
	}
}

/*
GetDebugCommandResultJSONParams contains all the parameters to send to the API endpoint

	for the get debug command result Json operation.

	Typically these are written to a http.Request.
*/
type GetDebugCommandResultJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// CollectorID.
	//
	// Format: int32
	// Default: -1
	CollectorID *int32

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get debug command result Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDebugCommandResultJSONParams) WithDefaults() *GetDebugCommandResultJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get debug command result Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDebugCommandResultJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		collectorIDDefault = int32(-1)
	)

	val := GetDebugCommandResultJSONParams{
		UserAgent:   &userAgentDefault,
		CollectorID: &collectorIDDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) WithTimeout(timeout time.Duration) *GetDebugCommandResultJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) WithContext(ctx context.Context) *GetDebugCommandResultJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) WithHTTPClient(client *http.Client) *GetDebugCommandResultJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) WithUserAgent(userAgent *string) *GetDebugCommandResultJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithCollectorID adds the collectorID to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) WithCollectorID(collectorID *int32) *GetDebugCommandResultJSONParams {
	o.SetCollectorID(collectorID)
	return o
}

// SetCollectorID adds the collectorId to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) SetCollectorID(collectorID *int32) {
	o.CollectorID = collectorID
}

// WithID adds the id to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) WithID(id string) *GetDebugCommandResultJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get debug command result Json params
func (o *GetDebugCommandResultJSONParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDebugCommandResultJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CollectorID != nil {

		// query param collectorId
		var qrCollectorID int32

		if o.CollectorID != nil {
			qrCollectorID = *o.CollectorID
		}
		qCollectorID := swag.FormatInt32(qrCollectorID)
		if qCollectorID != "" {

			if err := r.SetQueryParam("collectorId", qCollectorID); err != nil {
				return err
			}
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