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

// NewDeleteDeviceByIDParams creates a new DeleteDeviceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeviceByIDParams() *DeleteDeviceByIDParams {
	return &DeleteDeviceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceByIDParamsWithTimeout creates a new DeleteDeviceByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDeviceByIDParamsWithTimeout(timeout time.Duration) *DeleteDeviceByIDParams {
	return &DeleteDeviceByIDParams{
		timeout: timeout,
	}
}

// NewDeleteDeviceByIDParamsWithContext creates a new DeleteDeviceByIDParams object
// with the ability to set a context for a request.
func NewDeleteDeviceByIDParamsWithContext(ctx context.Context) *DeleteDeviceByIDParams {
	return &DeleteDeviceByIDParams{
		Context: ctx,
	}
}

// NewDeleteDeviceByIDParamsWithHTTPClient creates a new DeleteDeviceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeviceByIDParamsWithHTTPClient(client *http.Client) *DeleteDeviceByIDParams {
	return &DeleteDeviceByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteDeviceByIDParams contains all the parameters to send to the API endpoint

	for the delete device by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteDeviceByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// DeleteHard.
	//
	// Default: true
	DeleteHard *bool

	// End.
	//
	// Format: int64
	End *int64

	// ID.
	//
	// Format: int32
	ID int32

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

// WithDefaults hydrates default values in the delete device by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceByIDParams) WithDefaults() *DeleteDeviceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete device by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		deleteHardDefault = bool(true)
	)

	val := DeleteDeviceByIDParams{
		UserAgent:  &userAgentDefault,
		DeleteHard: &deleteHardDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithTimeout(timeout time.Duration) *DeleteDeviceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithContext(ctx context.Context) *DeleteDeviceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithHTTPClient(client *http.Client) *DeleteDeviceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithUserAgent(userAgent *string) *DeleteDeviceByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeleteHard adds the deleteHard to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithDeleteHard(deleteHard *bool) *DeleteDeviceByIDParams {
	o.SetDeleteHard(deleteHard)
	return o
}

// SetDeleteHard adds the deleteHard to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetDeleteHard(deleteHard *bool) {
	o.DeleteHard = deleteHard
}

// WithEnd adds the end to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithEnd(end *int64) *DeleteDeviceByIDParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetEnd(end *int64) {
	o.End = end
}

// WithID adds the id to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithID(id int32) *DeleteDeviceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetID(id int32) {
	o.ID = id
}

// WithNetflowFilter adds the netflowFilter to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithNetflowFilter(netflowFilter *string) *DeleteDeviceByIDParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithStart adds the start to the delete device by Id params
func (o *DeleteDeviceByIDParams) WithStart(start *int64) *DeleteDeviceByIDParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the delete device by Id params
func (o *DeleteDeviceByIDParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeleteHard != nil {

		// query param deleteHard
		var qrDeleteHard bool

		if o.DeleteHard != nil {
			qrDeleteHard = *o.DeleteHard
		}
		qDeleteHard := swag.FormatBool(qrDeleteHard)
		if qDeleteHard != "" {

			if err := r.SetQueryParam("deleteHard", qDeleteHard); err != nil {
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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
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
