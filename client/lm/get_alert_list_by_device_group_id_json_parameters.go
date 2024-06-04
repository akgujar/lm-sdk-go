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

// NewGetAlertListByDeviceGroupIDJSONParams creates a new GetAlertListByDeviceGroupIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertListByDeviceGroupIDJSONParams() *GetAlertListByDeviceGroupIDJSONParams {
	return &GetAlertListByDeviceGroupIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertListByDeviceGroupIDJSONParamsWithTimeout creates a new GetAlertListByDeviceGroupIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetAlertListByDeviceGroupIDJSONParamsWithTimeout(timeout time.Duration) *GetAlertListByDeviceGroupIDJSONParams {
	return &GetAlertListByDeviceGroupIDJSONParams{
		timeout: timeout,
	}
}

// NewGetAlertListByDeviceGroupIDJSONParamsWithContext creates a new GetAlertListByDeviceGroupIDJSONParams object
// with the ability to set a context for a request.
func NewGetAlertListByDeviceGroupIDJSONParamsWithContext(ctx context.Context) *GetAlertListByDeviceGroupIDJSONParams {
	return &GetAlertListByDeviceGroupIDJSONParams{
		Context: ctx,
	}
}

// NewGetAlertListByDeviceGroupIDJSONParamsWithHTTPClient creates a new GetAlertListByDeviceGroupIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertListByDeviceGroupIDJSONParamsWithHTTPClient(client *http.Client) *GetAlertListByDeviceGroupIDJSONParams {
	return &GetAlertListByDeviceGroupIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetAlertListByDeviceGroupIDJSONParams contains all the parameters to send to the API endpoint

	for the get alert list by device group Id Json operation.

	Typically these are written to a http.Request.
*/
type GetAlertListByDeviceGroupIDJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// CustomColumns.
	CustomColumns *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

	// NeedMessage.
	NeedMessage *bool

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

// WithDefaults hydrates default values in the get alert list by device group Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertListByDeviceGroupIDJSONParams) WithDefaults() *GetAlertListByDeviceGroupIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert list by device group Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertListByDeviceGroupIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		needMessageDefault = bool(false)

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetAlertListByDeviceGroupIDJSONParams{
		UserAgent:   &userAgentDefault,
		NeedMessage: &needMessageDefault,
		Offset:      &offsetDefault,
		Size:        &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithTimeout(timeout time.Duration) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithContext(ctx context.Context) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithHTTPClient(client *http.Client) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithUserAgent(userAgent *string) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithCustomColumns adds the customColumns to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithCustomColumns(customColumns *string) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetCustomColumns(customColumns)
	return o
}

// SetCustomColumns adds the customColumns to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetCustomColumns(customColumns *string) {
	o.CustomColumns = customColumns
}

// WithFields adds the fields to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithFields(fields *string) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithFilter(filter *string) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithID(id int32) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WithNeedMessage adds the needMessage to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithNeedMessage(needMessage *bool) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetNeedMessage(needMessage)
	return o
}

// SetNeedMessage adds the needMessage to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetNeedMessage(needMessage *bool) {
	o.NeedMessage = needMessage
}

// WithOffset adds the offset to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithOffset(offset *int32) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) WithSize(size *int32) *GetAlertListByDeviceGroupIDJSONParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get alert list by device group Id Json params
func (o *GetAlertListByDeviceGroupIDJSONParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertListByDeviceGroupIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CustomColumns != nil {

		// query param customColumns
		var qrCustomColumns string

		if o.CustomColumns != nil {
			qrCustomColumns = *o.CustomColumns
		}
		qCustomColumns := qrCustomColumns
		if qCustomColumns != "" {

			if err := r.SetQueryParam("customColumns", qCustomColumns); err != nil {
				return err
			}
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

	if o.NeedMessage != nil {

		// query param needMessage
		var qrNeedMessage bool

		if o.NeedMessage != nil {
			qrNeedMessage = *o.NeedMessage
		}
		qNeedMessage := swag.FormatBool(qrNeedMessage)
		if qNeedMessage != "" {

			if err := r.SetQueryParam("needMessage", qNeedMessage); err != nil {
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