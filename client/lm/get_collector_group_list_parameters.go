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

// NewGetCollectorGroupListParams creates a new GetCollectorGroupListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCollectorGroupListParams() *GetCollectorGroupListParams {
	return &GetCollectorGroupListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCollectorGroupListParamsWithTimeout creates a new GetCollectorGroupListParams object
// with the ability to set a timeout on a request.
func NewGetCollectorGroupListParamsWithTimeout(timeout time.Duration) *GetCollectorGroupListParams {
	return &GetCollectorGroupListParams{
		timeout: timeout,
	}
}

// NewGetCollectorGroupListParamsWithContext creates a new GetCollectorGroupListParams object
// with the ability to set a context for a request.
func NewGetCollectorGroupListParamsWithContext(ctx context.Context) *GetCollectorGroupListParams {
	return &GetCollectorGroupListParams{
		Context: ctx,
	}
}

// NewGetCollectorGroupListParamsWithHTTPClient creates a new GetCollectorGroupListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCollectorGroupListParamsWithHTTPClient(client *http.Client) *GetCollectorGroupListParams {
	return &GetCollectorGroupListParams{
		HTTPClient: client,
	}
}

/* GetCollectorGroupListParams contains all the parameters to send to the API endpoint
   for the get collector group list operation.

   Typically these are written to a http.Request.
*/
type GetCollectorGroupListParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get collector group list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorGroupListParams) WithDefaults() *GetCollectorGroupListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get collector group list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectorGroupListParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetCollectorGroupListParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get collector group list params
func (o *GetCollectorGroupListParams) WithTimeout(timeout time.Duration) *GetCollectorGroupListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get collector group list params
func (o *GetCollectorGroupListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get collector group list params
func (o *GetCollectorGroupListParams) WithContext(ctx context.Context) *GetCollectorGroupListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get collector group list params
func (o *GetCollectorGroupListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get collector group list params
func (o *GetCollectorGroupListParams) WithHTTPClient(client *http.Client) *GetCollectorGroupListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get collector group list params
func (o *GetCollectorGroupListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get collector group list params
func (o *GetCollectorGroupListParams) WithUserAgent(userAgent *string) *GetCollectorGroupListParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get collector group list params
func (o *GetCollectorGroupListParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get collector group list params
func (o *GetCollectorGroupListParams) WithFields(fields *string) *GetCollectorGroupListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get collector group list params
func (o *GetCollectorGroupListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get collector group list params
func (o *GetCollectorGroupListParams) WithFilter(filter *string) *GetCollectorGroupListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get collector group list params
func (o *GetCollectorGroupListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get collector group list params
func (o *GetCollectorGroupListParams) WithOffset(offset *int32) *GetCollectorGroupListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get collector group list params
func (o *GetCollectorGroupListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get collector group list params
func (o *GetCollectorGroupListParams) WithSize(size *int32) *GetCollectorGroupListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get collector group list params
func (o *GetCollectorGroupListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetCollectorGroupListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
