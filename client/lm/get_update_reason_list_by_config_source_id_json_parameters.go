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

// NewGetUpdateReasonListByConfigSourceIDJSONParams creates a new GetUpdateReasonListByConfigSourceIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpdateReasonListByConfigSourceIDJSONParams() *GetUpdateReasonListByConfigSourceIDJSONParams {
	return &GetUpdateReasonListByConfigSourceIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpdateReasonListByConfigSourceIDJSONParamsWithTimeout creates a new GetUpdateReasonListByConfigSourceIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetUpdateReasonListByConfigSourceIDJSONParamsWithTimeout(timeout time.Duration) *GetUpdateReasonListByConfigSourceIDJSONParams {
	return &GetUpdateReasonListByConfigSourceIDJSONParams{
		timeout: timeout,
	}
}

// NewGetUpdateReasonListByConfigSourceIDJSONParamsWithContext creates a new GetUpdateReasonListByConfigSourceIDJSONParams object
// with the ability to set a context for a request.
func NewGetUpdateReasonListByConfigSourceIDJSONParamsWithContext(ctx context.Context) *GetUpdateReasonListByConfigSourceIDJSONParams {
	return &GetUpdateReasonListByConfigSourceIDJSONParams{
		Context: ctx,
	}
}

// NewGetUpdateReasonListByConfigSourceIDJSONParamsWithHTTPClient creates a new GetUpdateReasonListByConfigSourceIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpdateReasonListByConfigSourceIDJSONParamsWithHTTPClient(client *http.Client) *GetUpdateReasonListByConfigSourceIDJSONParams {
	return &GetUpdateReasonListByConfigSourceIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetUpdateReasonListByConfigSourceIDJSONParams contains all the parameters to send to the API endpoint

	for the get update reason list by config source Id Json operation.

	Typically these are written to a http.Request.
*/
type GetUpdateReasonListByConfigSourceIDJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get update reason list by config source Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) WithDefaults() *GetUpdateReasonListByConfigSourceIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get update reason list by config source Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetUpdateReasonListByConfigSourceIDJSONParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) WithTimeout(timeout time.Duration) *GetUpdateReasonListByConfigSourceIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) WithContext(ctx context.Context) *GetUpdateReasonListByConfigSourceIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) WithHTTPClient(client *http.Client) *GetUpdateReasonListByConfigSourceIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) WithUserAgent(userAgent *string) *GetUpdateReasonListByConfigSourceIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) WithID(id int32) *GetUpdateReasonListByConfigSourceIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get update reason list by config source Id Json params
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpdateReasonListByConfigSourceIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
