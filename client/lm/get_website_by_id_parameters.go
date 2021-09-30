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

// NewGetWebsiteByIDParams creates a new GetWebsiteByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsiteByIDParams() *GetWebsiteByIDParams {
	return &GetWebsiteByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteByIDParamsWithTimeout creates a new GetWebsiteByIDParams object
// with the ability to set a timeout on a request.
func NewGetWebsiteByIDParamsWithTimeout(timeout time.Duration) *GetWebsiteByIDParams {
	return &GetWebsiteByIDParams{
		timeout: timeout,
	}
}

// NewGetWebsiteByIDParamsWithContext creates a new GetWebsiteByIDParams object
// with the ability to set a context for a request.
func NewGetWebsiteByIDParamsWithContext(ctx context.Context) *GetWebsiteByIDParams {
	return &GetWebsiteByIDParams{
		Context: ctx,
	}
}

// NewGetWebsiteByIDParamsWithHTTPClient creates a new GetWebsiteByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsiteByIDParamsWithHTTPClient(client *http.Client) *GetWebsiteByIDParams {
	return &GetWebsiteByIDParams{
		HTTPClient: client,
	}
}

/* GetWebsiteByIDParams contains all the parameters to send to the API endpoint
   for the get website by Id operation.

   Typically these are written to a http.Request.
*/
type GetWebsiteByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// Format.
	//
	// Default: "json"
	Format *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get website by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteByIDParams) WithDefaults() *GetWebsiteByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")

		formatDefault = string("json")
	)

	val := GetWebsiteByIDParams{
		UserAgent: &userAgentDefault,
		Format:    &formatDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get website by Id params
func (o *GetWebsiteByIDParams) WithTimeout(timeout time.Duration) *GetWebsiteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website by Id params
func (o *GetWebsiteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website by Id params
func (o *GetWebsiteByIDParams) WithContext(ctx context.Context) *GetWebsiteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website by Id params
func (o *GetWebsiteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website by Id params
func (o *GetWebsiteByIDParams) WithHTTPClient(client *http.Client) *GetWebsiteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website by Id params
func (o *GetWebsiteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get website by Id params
func (o *GetWebsiteByIDParams) WithUserAgent(userAgent *string) *GetWebsiteByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get website by Id params
func (o *GetWebsiteByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFormat adds the format to the get website by Id params
func (o *GetWebsiteByIDParams) WithFormat(format *string) *GetWebsiteByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get website by Id params
func (o *GetWebsiteByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get website by Id params
func (o *GetWebsiteByIDParams) WithID(id int32) *GetWebsiteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get website by Id params
func (o *GetWebsiteByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
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
