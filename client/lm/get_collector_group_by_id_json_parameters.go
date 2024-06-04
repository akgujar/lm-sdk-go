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

// NewGetCollectorGroupByIDJSONParams creates a new GetCollectorGroupByIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCollectorGroupByIDJSONParams() *GetCollectorGroupByIDJSONParams {
	return &GetCollectorGroupByIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCollectorGroupByIDJSONParamsWithTimeout creates a new GetCollectorGroupByIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetCollectorGroupByIDJSONParamsWithTimeout(timeout time.Duration) *GetCollectorGroupByIDJSONParams {
	return &GetCollectorGroupByIDJSONParams{
		timeout: timeout,
	}
}

// NewGetCollectorGroupByIDJSONParamsWithContext creates a new GetCollectorGroupByIDJSONParams object
// with the ability to set a context for a request.
func NewGetCollectorGroupByIDJSONParamsWithContext(ctx context.Context) *GetCollectorGroupByIDJSONParams {
	return &GetCollectorGroupByIDJSONParams{
		Context: ctx,
	}
}

// NewGetCollectorGroupByIDJSONParamsWithHTTPClient creates a new GetCollectorGroupByIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCollectorGroupByIDJSONParamsWithHTTPClient(client *http.Client) *GetCollectorGroupByIDJSONParams {
	return &GetCollectorGroupByIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetCollectorGroupByIDJSONParams contains all the parameters to send to the API endpoint

	for the get collector group by Id Json operation.

	Typically these are written to a http.Request.
*/
type GetCollectorGroupByIDJSONParams struct {

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

// WithDefaults hydrates default values in the get collector group by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorGroupByIDJSONParams) WithDefaults() *GetCollectorGroupByIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get collector group by Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorGroupByIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetCollectorGroupByIDJSONParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) WithTimeout(timeout time.Duration) *GetCollectorGroupByIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) WithContext(ctx context.Context) *GetCollectorGroupByIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) WithHTTPClient(client *http.Client) *GetCollectorGroupByIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) WithUserAgent(userAgent *string) *GetCollectorGroupByIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) WithFields(fields *string) *GetCollectorGroupByIDJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) WithID(id int32) *GetCollectorGroupByIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get collector group by Id Json params
func (o *GetCollectorGroupByIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCollectorGroupByIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
