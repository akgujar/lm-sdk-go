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

// NewGetWebsiteByIDJSONParams creates a new GetWebsiteByIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsiteByIDJSONParams() *GetWebsiteByIDJSONParams {
	return &GetWebsiteByIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteByIDJSONParamsWithTimeout creates a new GetWebsiteByIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetWebsiteByIDJSONParamsWithTimeout(timeout time.Duration) *GetWebsiteByIDJSONParams {
	return &GetWebsiteByIDJSONParams{
		timeout: timeout,
	}
}

// NewGetWebsiteByIDJSONParamsWithContext creates a new GetWebsiteByIDJSONParams object
// with the ability to set a context for a request.
func NewGetWebsiteByIDJSONParamsWithContext(ctx context.Context) *GetWebsiteByIDJSONParams {
	return &GetWebsiteByIDJSONParams{
		Context: ctx,
	}
}

// NewGetWebsiteByIDJSONParamsWithHTTPClient creates a new GetWebsiteByIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsiteByIDJSONParamsWithHTTPClient(client *http.Client) *GetWebsiteByIDJSONParams {
	return &GetWebsiteByIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetWebsiteByIDJSONParams contains all the parameters to send to the API endpoint

	for the get website by Id Json operation.

	Typically these are written to a http.Request.
*/
type GetWebsiteByIDJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
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

// WithDefaults hydrates default values in the get website by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteByIDJSONParams) WithDefaults() *GetWebsiteByIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteByIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		formatDefault = string("json")
	)

	val := GetWebsiteByIDJSONParams{
		UserAgent: &userAgentDefault,
		Format:    &formatDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) WithTimeout(timeout time.Duration) *GetWebsiteByIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) WithContext(ctx context.Context) *GetWebsiteByIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) WithHTTPClient(client *http.Client) *GetWebsiteByIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) WithUserAgent(userAgent *string) *GetWebsiteByIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFormat adds the format to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) WithFormat(format *string) *GetWebsiteByIDJSONParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) WithID(id int32) *GetWebsiteByIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get website by Id Json params
func (o *GetWebsiteByIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteByIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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