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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewGetCollectorByHostNameParams creates a new GetCollectorByHostNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCollectorByHostNameParams() *GetCollectorByHostNameParams {
	return &GetCollectorByHostNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCollectorByHostNameParamsWithTimeout creates a new GetCollectorByHostNameParams object
// with the ability to set a timeout on a request.
func NewGetCollectorByHostNameParamsWithTimeout(timeout time.Duration) *GetCollectorByHostNameParams {
	return &GetCollectorByHostNameParams{
		timeout: timeout,
	}
}

// NewGetCollectorByHostNameParamsWithContext creates a new GetCollectorByHostNameParams object
// with the ability to set a context for a request.
func NewGetCollectorByHostNameParamsWithContext(ctx context.Context) *GetCollectorByHostNameParams {
	return &GetCollectorByHostNameParams{
		Context: ctx,
	}
}

// NewGetCollectorByHostNameParamsWithHTTPClient creates a new GetCollectorByHostNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCollectorByHostNameParamsWithHTTPClient(client *http.Client) *GetCollectorByHostNameParams {
	return &GetCollectorByHostNameParams{
		HTTPClient: client,
	}
}

/* GetCollectorByHostNameParams contains all the parameters to send to the API endpoint
   for the get collector by host name operation.

   Typically these are written to a http.Request.
*/
type GetCollectorByHostNameParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.Collector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get collector by host name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorByHostNameParams) WithDefaults() *GetCollectorByHostNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get collector by host name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorByHostNameParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetCollectorByHostNameParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get collector by host name params
func (o *GetCollectorByHostNameParams) WithTimeout(timeout time.Duration) *GetCollectorByHostNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get collector by host name params
func (o *GetCollectorByHostNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get collector by host name params
func (o *GetCollectorByHostNameParams) WithContext(ctx context.Context) *GetCollectorByHostNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get collector by host name params
func (o *GetCollectorByHostNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get collector by host name params
func (o *GetCollectorByHostNameParams) WithHTTPClient(client *http.Client) *GetCollectorByHostNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get collector by host name params
func (o *GetCollectorByHostNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get collector by host name params
func (o *GetCollectorByHostNameParams) WithUserAgent(userAgent *string) *GetCollectorByHostNameParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get collector by host name params
func (o *GetCollectorByHostNameParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the get collector by host name params
func (o *GetCollectorByHostNameParams) WithBody(body *models.Collector) *GetCollectorByHostNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get collector by host name params
func (o *GetCollectorByHostNameParams) SetBody(body *models.Collector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetCollectorByHostNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
