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

// NewDeleteSDTByIDParams creates a new DeleteSDTByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSDTByIDParams() *DeleteSDTByIDParams {
	return &DeleteSDTByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSDTByIDParamsWithTimeout creates a new DeleteSDTByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteSDTByIDParamsWithTimeout(timeout time.Duration) *DeleteSDTByIDParams {
	return &DeleteSDTByIDParams{
		timeout: timeout,
	}
}

// NewDeleteSDTByIDParamsWithContext creates a new DeleteSDTByIDParams object
// with the ability to set a context for a request.
func NewDeleteSDTByIDParamsWithContext(ctx context.Context) *DeleteSDTByIDParams {
	return &DeleteSDTByIDParams{
		Context: ctx,
	}
}

// NewDeleteSDTByIDParamsWithHTTPClient creates a new DeleteSDTByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSDTByIDParamsWithHTTPClient(client *http.Client) *DeleteSDTByIDParams {
	return &DeleteSDTByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteSDTByIDParams contains all the parameters to send to the API endpoint

	for the delete SDT by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteSDTByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete SDT by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSDTByIDParams) WithDefaults() *DeleteSDTByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete SDT by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSDTByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := DeleteSDTByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete SDT by Id params
func (o *DeleteSDTByIDParams) WithTimeout(timeout time.Duration) *DeleteSDTByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete SDT by Id params
func (o *DeleteSDTByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete SDT by Id params
func (o *DeleteSDTByIDParams) WithContext(ctx context.Context) *DeleteSDTByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete SDT by Id params
func (o *DeleteSDTByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete SDT by Id params
func (o *DeleteSDTByIDParams) WithHTTPClient(client *http.Client) *DeleteSDTByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete SDT by Id params
func (o *DeleteSDTByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete SDT by Id params
func (o *DeleteSDTByIDParams) WithUserAgent(userAgent *string) *DeleteSDTByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete SDT by Id params
func (o *DeleteSDTByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the delete SDT by Id params
func (o *DeleteSDTByIDParams) WithID(id string) *DeleteSDTByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete SDT by Id params
func (o *DeleteSDTByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSDTByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
