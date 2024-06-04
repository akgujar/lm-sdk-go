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

// NewGetAlertListByDeviceIDJSONParams creates a new GetAlertListByDeviceIDJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertListByDeviceIDJSONParams() *GetAlertListByDeviceIDJSONParams {
	return &GetAlertListByDeviceIDJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertListByDeviceIDJSONParamsWithTimeout creates a new GetAlertListByDeviceIDJSONParams object
// with the ability to set a timeout on a request.
func NewGetAlertListByDeviceIDJSONParamsWithTimeout(timeout time.Duration) *GetAlertListByDeviceIDJSONParams {
	return &GetAlertListByDeviceIDJSONParams{
		timeout: timeout,
	}
}

// NewGetAlertListByDeviceIDJSONParamsWithContext creates a new GetAlertListByDeviceIDJSONParams object
// with the ability to set a context for a request.
func NewGetAlertListByDeviceIDJSONParamsWithContext(ctx context.Context) *GetAlertListByDeviceIDJSONParams {
	return &GetAlertListByDeviceIDJSONParams{
		Context: ctx,
	}
}

// NewGetAlertListByDeviceIDJSONParamsWithHTTPClient creates a new GetAlertListByDeviceIDJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertListByDeviceIDJSONParamsWithHTTPClient(client *http.Client) *GetAlertListByDeviceIDJSONParams {
	return &GetAlertListByDeviceIDJSONParams{
		HTTPClient: client,
	}
}

/*
GetAlertListByDeviceIDJSONParams contains all the parameters to send to the API endpoint

	for the get alert list by device Id Json operation.

	Typically these are written to a http.Request.
*/
type GetAlertListByDeviceIDJSONParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Bound.
	//
	// Default: "instances"
	Bound *string

	// CustomColumns.
	CustomColumns *string

	// End.
	//
	// Format: int64
	End *int64

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

	// NetflowFilter.
	NetflowFilter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert list by device Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertListByDeviceIDJSONParams) WithDefaults() *GetAlertListByDeviceIDJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert list by device Id Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertListByDeviceIDJSONParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		boundDefault = string("instances")

		needMessageDefault = bool(false)

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetAlertListByDeviceIDJSONParams{
		UserAgent:   &userAgentDefault,
		Bound:       &boundDefault,
		NeedMessage: &needMessageDefault,
		Offset:      &offsetDefault,
		Size:        &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithTimeout(timeout time.Duration) *GetAlertListByDeviceIDJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithContext(ctx context.Context) *GetAlertListByDeviceIDJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithHTTPClient(client *http.Client) *GetAlertListByDeviceIDJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithUserAgent(userAgent *string) *GetAlertListByDeviceIDJSONParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBound adds the bound to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithBound(bound *string) *GetAlertListByDeviceIDJSONParams {
	o.SetBound(bound)
	return o
}

// SetBound adds the bound to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetBound(bound *string) {
	o.Bound = bound
}

// WithCustomColumns adds the customColumns to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithCustomColumns(customColumns *string) *GetAlertListByDeviceIDJSONParams {
	o.SetCustomColumns(customColumns)
	return o
}

// SetCustomColumns adds the customColumns to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetCustomColumns(customColumns *string) {
	o.CustomColumns = customColumns
}

// WithEnd adds the end to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithEnd(end *int64) *GetAlertListByDeviceIDJSONParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetEnd(end *int64) {
	o.End = end
}

// WithFields adds the fields to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithFields(fields *string) *GetAlertListByDeviceIDJSONParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithFilter(filter *string) *GetAlertListByDeviceIDJSONParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithID(id int32) *GetAlertListByDeviceIDJSONParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetID(id int32) {
	o.ID = id
}

// WithNeedMessage adds the needMessage to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithNeedMessage(needMessage *bool) *GetAlertListByDeviceIDJSONParams {
	o.SetNeedMessage(needMessage)
	return o
}

// SetNeedMessage adds the needMessage to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetNeedMessage(needMessage *bool) {
	o.NeedMessage = needMessage
}

// WithNetflowFilter adds the netflowFilter to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithNetflowFilter(netflowFilter *string) *GetAlertListByDeviceIDJSONParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithOffset adds the offset to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithOffset(offset *int32) *GetAlertListByDeviceIDJSONParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithSize(size *int32) *GetAlertListByDeviceIDJSONParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetSize(size *int32) {
	o.Size = size
}

// WithStart adds the start to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) WithStart(start *int64) *GetAlertListByDeviceIDJSONParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get alert list by device Id Json params
func (o *GetAlertListByDeviceIDJSONParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertListByDeviceIDJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Bound != nil {

		// query param bound
		var qrBound string

		if o.Bound != nil {
			qrBound = *o.Bound
		}
		qBound := qrBound
		if qBound != "" {

			if err := r.SetQueryParam("bound", qBound); err != nil {
				return err
			}
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