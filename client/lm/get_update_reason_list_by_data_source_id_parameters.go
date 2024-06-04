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

// NewGetUpdateReasonListByDataSourceIDParams creates a new GetUpdateReasonListByDataSourceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpdateReasonListByDataSourceIDParams() *GetUpdateReasonListByDataSourceIDParams {
	return &GetUpdateReasonListByDataSourceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpdateReasonListByDataSourceIDParamsWithTimeout creates a new GetUpdateReasonListByDataSourceIDParams object
// with the ability to set a timeout on a request.
func NewGetUpdateReasonListByDataSourceIDParamsWithTimeout(timeout time.Duration) *GetUpdateReasonListByDataSourceIDParams {
	return &GetUpdateReasonListByDataSourceIDParams{
		timeout: timeout,
	}
}

// NewGetUpdateReasonListByDataSourceIDParamsWithContext creates a new GetUpdateReasonListByDataSourceIDParams object
// with the ability to set a context for a request.
func NewGetUpdateReasonListByDataSourceIDParamsWithContext(ctx context.Context) *GetUpdateReasonListByDataSourceIDParams {
	return &GetUpdateReasonListByDataSourceIDParams{
		Context: ctx,
	}
}

// NewGetUpdateReasonListByDataSourceIDParamsWithHTTPClient creates a new GetUpdateReasonListByDataSourceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpdateReasonListByDataSourceIDParamsWithHTTPClient(client *http.Client) *GetUpdateReasonListByDataSourceIDParams {
	return &GetUpdateReasonListByDataSourceIDParams{
		HTTPClient: client,
	}
}

/*
GetUpdateReasonListByDataSourceIDParams contains all the parameters to send to the API endpoint

	for the get update reason list by data source Id operation.

	Typically these are written to a http.Request.
*/
type GetUpdateReasonListByDataSourceIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

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

// WithDefaults hydrates default values in the get update reason list by data source Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpdateReasonListByDataSourceIDParams) WithDefaults() *GetUpdateReasonListByDataSourceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get update reason list by data source Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpdateReasonListByDataSourceIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetUpdateReasonListByDataSourceIDParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithTimeout(timeout time.Duration) *GetUpdateReasonListByDataSourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithContext(ctx context.Context) *GetUpdateReasonListByDataSourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithHTTPClient(client *http.Client) *GetUpdateReasonListByDataSourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithUserAgent(userAgent *string) *GetUpdateReasonListByDataSourceIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithFields(fields *string) *GetUpdateReasonListByDataSourceIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithFilter(filter *string) *GetUpdateReasonListByDataSourceIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithID(id int32) *GetUpdateReasonListByDataSourceIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithOffset(offset *int32) *GetUpdateReasonListByDataSourceIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) WithSize(size *int32) *GetUpdateReasonListByDataSourceIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get update reason list by data source Id params
func (o *GetUpdateReasonListByDataSourceIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpdateReasonListByDataSourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
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
