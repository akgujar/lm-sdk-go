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

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// NewGenerateReportByIDParams creates a new GenerateReportByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateReportByIDParams() *GenerateReportByIDParams {
	return &GenerateReportByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateReportByIDParamsWithTimeout creates a new GenerateReportByIDParams object
// with the ability to set a timeout on a request.
func NewGenerateReportByIDParamsWithTimeout(timeout time.Duration) *GenerateReportByIDParams {
	return &GenerateReportByIDParams{
		timeout: timeout,
	}
}

// NewGenerateReportByIDParamsWithContext creates a new GenerateReportByIDParams object
// with the ability to set a context for a request.
func NewGenerateReportByIDParamsWithContext(ctx context.Context) *GenerateReportByIDParams {
	return &GenerateReportByIDParams{
		Context: ctx,
	}
}

// NewGenerateReportByIDParamsWithHTTPClient creates a new GenerateReportByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateReportByIDParamsWithHTTPClient(client *http.Client) *GenerateReportByIDParams {
	return &GenerateReportByIDParams{
		HTTPClient: client,
	}
}

/*
GenerateReportByIDParams contains all the parameters to send to the API endpoint

	for the generate report by Id operation.

	Typically these are written to a http.Request.
*/
type GenerateReportByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.GenerateReportRequest

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate report by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateReportByIDParams) WithDefaults() *GenerateReportByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate report by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateReportByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := GenerateReportByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the generate report by Id params
func (o *GenerateReportByIDParams) WithTimeout(timeout time.Duration) *GenerateReportByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate report by Id params
func (o *GenerateReportByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate report by Id params
func (o *GenerateReportByIDParams) WithContext(ctx context.Context) *GenerateReportByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate report by Id params
func (o *GenerateReportByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate report by Id params
func (o *GenerateReportByIDParams) WithHTTPClient(client *http.Client) *GenerateReportByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate report by Id params
func (o *GenerateReportByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the generate report by Id params
func (o *GenerateReportByIDParams) WithUserAgent(userAgent *string) *GenerateReportByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the generate report by Id params
func (o *GenerateReportByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the generate report by Id params
func (o *GenerateReportByIDParams) WithBody(body *models.GenerateReportRequest) *GenerateReportByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate report by Id params
func (o *GenerateReportByIDParams) SetBody(body *models.GenerateReportRequest) {
	o.Body = body
}

// WithID adds the id to the generate report by Id params
func (o *GenerateReportByIDParams) WithID(id int32) *GenerateReportByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the generate report by Id params
func (o *GenerateReportByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateReportByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
