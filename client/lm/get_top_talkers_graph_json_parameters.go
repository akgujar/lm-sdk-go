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

// NewGetTopTalkersGraphJSONParams creates a new GetTopTalkersGraphJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTopTalkersGraphJSONParams() *GetTopTalkersGraphJSONParams {
	return &GetTopTalkersGraphJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTopTalkersGraphJSONParamsWithTimeout creates a new GetTopTalkersGraphJSONParams object
// with the ability to set a timeout on a request.
func NewGetTopTalkersGraphJSONParamsWithTimeout(timeout time.Duration) *GetTopTalkersGraphJSONParams {
	return &GetTopTalkersGraphJSONParams{
		timeout: timeout,
	}
}

// NewGetTopTalkersGraphJSONParamsWithContext creates a new GetTopTalkersGraphJSONParams object
// with the ability to set a context for a request.
func NewGetTopTalkersGraphJSONParamsWithContext(ctx context.Context) *GetTopTalkersGraphJSONParams {
	return &GetTopTalkersGraphJSONParams{
		Context: ctx,
	}
}

// NewGetTopTalkersGraphJSONParamsWithHTTPClient creates a new GetTopTalkersGraphJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTopTalkersGraphJSONParamsWithHTTPClient(client *http.Client) *GetTopTalkersGraphJSONParams {
	return &GetTopTalkersGraphJSONParams{
		HTTPClient: client,
	}
}

/*
GetTopTalkersGraphJSONParams contains all the parameters to send to the API endpoint

	for the get top talkers graph Json operation.

	Typically these are written to a http.Request.
*/
type GetTopTalkersGraphJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// End.
	//
	// Format: int64
	End *int64

	// Format.
	Format *string

	// ID.
	//
	// Format: int32
	ID int32

	// Keyword.
	Keyword *string

	// NetflowFilter.
	NetflowFilter *string

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get top talkers graph Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTopTalkersGraphJSONParams) WithDefaults() *GetTopTalkersGraphJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get top talkers graph Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTopTalkersGraphJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GetTopTalkersGraphJSONParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithTimeout(timeout time.Duration) *GetTopTalkersGraphJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithContext(ctx context.Context) *GetTopTalkersGraphJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithHTTPClient(client *http.Client) *GetTopTalkersGraphJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithUserAgent(userAgent *string) *GetTopTalkersGraphJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithEnd adds the end to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithEnd(end *int64) *GetTopTalkersGraphJSONParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithFormat(format *string) *GetTopTalkersGraphJSONParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithID(id int32) *GetTopTalkersGraphJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetID(id int32) {
	o.ID = id
}

// WithKeyword adds the keyword to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithKeyword(keyword *string) *GetTopTalkersGraphJSONParams {
	o.SetKeyword(keyword)
	return o
}

// SetKeyword adds the keyword to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetKeyword(keyword *string) {
	o.Keyword = keyword
}

// WithNetflowFilter adds the netflowFilter to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithNetflowFilter(netflowFilter *string) *GetTopTalkersGraphJSONParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithStart adds the start to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) WithStart(start *int64) *GetTopTalkersGraphJSONParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get top talkers graph Json params
func (o *GetTopTalkersGraphJSONParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetTopTalkersGraphJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
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

	if o.Keyword != nil {

		// query param keyword
		var qrKeyword string

		if o.Keyword != nil {
			qrKeyword = *o.Keyword
		}
		qKeyword := qrKeyword
		if qKeyword != "" {

			if err := r.SetQueryParam("keyword", qKeyword); err != nil {
				return err
			}
		}
	}

	if o.NetflowFilter != nil {

		// query param netflowFilter
		var qrNetflowFilter string

		if o.NetflowFilter != nil {
			qrNetflowFilter = *o.NetflowFilter
		}
		qNetflowFilter := qrNetflowFilter
		if qNetflowFilter != "" {

			if err := r.SetQueryParam("netflowFilter", qNetflowFilter); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
