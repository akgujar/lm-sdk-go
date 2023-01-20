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

// NewGetExternalAPIStatsParams creates a new GetExternalAPIStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalAPIStatsParams() *GetExternalAPIStatsParams {
	return &GetExternalAPIStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalAPIStatsParamsWithTimeout creates a new GetExternalAPIStatsParams object
// with the ability to set a timeout on a request.
func NewGetExternalAPIStatsParamsWithTimeout(timeout time.Duration) *GetExternalAPIStatsParams {
	return &GetExternalAPIStatsParams{
		timeout: timeout,
	}
}

// NewGetExternalAPIStatsParamsWithContext creates a new GetExternalAPIStatsParams object
// with the ability to set a context for a request.
func NewGetExternalAPIStatsParamsWithContext(ctx context.Context) *GetExternalAPIStatsParams {
	return &GetExternalAPIStatsParams{
		Context: ctx,
	}
}

// NewGetExternalAPIStatsParamsWithHTTPClient creates a new GetExternalAPIStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalAPIStatsParamsWithHTTPClient(client *http.Client) *GetExternalAPIStatsParams {
	return &GetExternalAPIStatsParams{
		HTTPClient: client,
	}
}

/* GetExternalAPIStatsParams contains all the parameters to send to the API endpoint
   for the get external Api stats operation.

   Typically these are written to a http.Request.
*/
type GetExternalAPIStatsParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get external Api stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalAPIStatsParams) WithDefaults() *GetExternalAPIStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get external Api stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalAPIStatsParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetExternalAPIStatsParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get external Api stats params
func (o *GetExternalAPIStatsParams) WithTimeout(timeout time.Duration) *GetExternalAPIStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external Api stats params
func (o *GetExternalAPIStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external Api stats params
func (o *GetExternalAPIStatsParams) WithContext(ctx context.Context) *GetExternalAPIStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external Api stats params
func (o *GetExternalAPIStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external Api stats params
func (o *GetExternalAPIStatsParams) WithHTTPClient(client *http.Client) *GetExternalAPIStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external Api stats params
func (o *GetExternalAPIStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get external Api stats params
func (o *GetExternalAPIStatsParams) WithUserAgent(userAgent *string) *GetExternalAPIStatsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get external Api stats params
func (o *GetExternalAPIStatsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalAPIStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
