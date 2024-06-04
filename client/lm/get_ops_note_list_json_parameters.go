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

// NewGetOpsNoteListJSONParams creates a new GetOpsNoteListJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOpsNoteListJSONParams() *GetOpsNoteListJSONParams {
	return &GetOpsNoteListJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOpsNoteListJSONParamsWithTimeout creates a new GetOpsNoteListJSONParams object
// with the ability to set a timeout on a request.
func NewGetOpsNoteListJSONParamsWithTimeout(timeout time.Duration) *GetOpsNoteListJSONParams {
	return &GetOpsNoteListJSONParams{
		timeout: timeout,
	}
}

// NewGetOpsNoteListJSONParamsWithContext creates a new GetOpsNoteListJSONParams object
// with the ability to set a context for a request.
func NewGetOpsNoteListJSONParamsWithContext(ctx context.Context) *GetOpsNoteListJSONParams {
	return &GetOpsNoteListJSONParams{
		Context: ctx,
	}
}

// NewGetOpsNoteListJSONParamsWithHTTPClient creates a new GetOpsNoteListJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOpsNoteListJSONParamsWithHTTPClient(client *http.Client) *GetOpsNoteListJSONParams {
	return &GetOpsNoteListJSONParams{
		HTTPClient: client,
	}
}

/*
GetOpsNoteListJSONParams contains all the parameters to send to the API endpoint

	for the get ops note list Json operation.

	Typically these are written to a http.Request.
*/
type GetOpsNoteListJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Fields.
	Fields *string

	/* Filter.

	   Filter the response based on tags, createdBy, happenedOn, monitorObjectGroups, monitorObjectNames, or _all field values
	*/
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

// WithDefaults hydrates default values in the get ops note list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOpsNoteListJSONParams) WithDefaults() *GetOpsNoteListJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ops note list Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOpsNoteListJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetOpsNoteListJSONParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithTimeout(timeout time.Duration) *GetOpsNoteListJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithContext(ctx context.Context) *GetOpsNoteListJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithHTTPClient(client *http.Client) *GetOpsNoteListJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithUserAgent(userAgent *string) *GetOpsNoteListJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithFields(fields *string) *GetOpsNoteListJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithFilter(filter *string) *GetOpsNoteListJSONParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithOffset(offset *int32) *GetOpsNoteListJSONParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) WithSize(size *int32) *GetOpsNoteListJSONParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get ops note list Json params
func (o *GetOpsNoteListJSONParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetOpsNoteListJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
