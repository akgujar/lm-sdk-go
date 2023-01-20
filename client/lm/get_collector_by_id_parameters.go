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

// NewGetCollectorByIDParams creates a new GetCollectorByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCollectorByIDParams() *GetCollectorByIDParams {
	return &GetCollectorByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCollectorByIDParamsWithTimeout creates a new GetCollectorByIDParams object
// with the ability to set a timeout on a request.
func NewGetCollectorByIDParamsWithTimeout(timeout time.Duration) *GetCollectorByIDParams {
	return &GetCollectorByIDParams{
		timeout: timeout,
	}
}

// NewGetCollectorByIDParamsWithContext creates a new GetCollectorByIDParams object
// with the ability to set a context for a request.
func NewGetCollectorByIDParamsWithContext(ctx context.Context) *GetCollectorByIDParams {
	return &GetCollectorByIDParams{
		Context: ctx,
	}
}

// NewGetCollectorByIDParamsWithHTTPClient creates a new GetCollectorByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCollectorByIDParamsWithHTTPClient(client *http.Client) *GetCollectorByIDParams {
	return &GetCollectorByIDParams{
		HTTPClient: client,
	}
}

/* GetCollectorByIDParams contains all the parameters to send to the API endpoint
   for the get collector by Id operation.

   Typically these are written to a http.Request.
*/
type GetCollectorByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorByIDParams) WithDefaults() *GetCollectorByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get collector by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetCollectorByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get collector by Id params
func (o *GetCollectorByIDParams) WithTimeout(timeout time.Duration) *GetCollectorByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get collector by Id params
func (o *GetCollectorByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get collector by Id params
func (o *GetCollectorByIDParams) WithContext(ctx context.Context) *GetCollectorByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get collector by Id params
func (o *GetCollectorByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get collector by Id params
func (o *GetCollectorByIDParams) WithHTTPClient(client *http.Client) *GetCollectorByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get collector by Id params
func (o *GetCollectorByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get collector by Id params
func (o *GetCollectorByIDParams) WithUserAgent(userAgent *string) *GetCollectorByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get collector by Id params
func (o *GetCollectorByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get collector by Id params
func (o *GetCollectorByIDParams) WithFields(fields *string) *GetCollectorByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get collector by Id params
func (o *GetCollectorByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get collector by Id params
func (o *GetCollectorByIDParams) WithID(id int32) *GetCollectorByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get collector by Id params
func (o *GetCollectorByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCollectorByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
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
