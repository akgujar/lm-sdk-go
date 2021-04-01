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

// NewGetEscalationChainByIDParams creates a new GetEscalationChainByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEscalationChainByIDParams() *GetEscalationChainByIDParams {
	return &GetEscalationChainByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEscalationChainByIDParamsWithTimeout creates a new GetEscalationChainByIDParams object
// with the ability to set a timeout on a request.
func NewGetEscalationChainByIDParamsWithTimeout(timeout time.Duration) *GetEscalationChainByIDParams {
	return &GetEscalationChainByIDParams{
		timeout: timeout,
	}
}

// NewGetEscalationChainByIDParamsWithContext creates a new GetEscalationChainByIDParams object
// with the ability to set a context for a request.
func NewGetEscalationChainByIDParamsWithContext(ctx context.Context) *GetEscalationChainByIDParams {
	return &GetEscalationChainByIDParams{
		Context: ctx,
	}
}

// NewGetEscalationChainByIDParamsWithHTTPClient creates a new GetEscalationChainByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEscalationChainByIDParamsWithHTTPClient(client *http.Client) *GetEscalationChainByIDParams {
	return &GetEscalationChainByIDParams{
		HTTPClient: client,
	}
}

/* GetEscalationChainByIDParams contains all the parameters to send to the API endpoint
   for the get escalation chain by Id operation.

   Typically these are written to a http.Request.
*/
type GetEscalationChainByIDParams struct {

	// Fields.
	Fields *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get escalation chain by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEscalationChainByIDParams) WithDefaults() *GetEscalationChainByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get escalation chain by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEscalationChainByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) WithTimeout(timeout time.Duration) *GetEscalationChainByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) WithContext(ctx context.Context) *GetEscalationChainByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) WithHTTPClient(client *http.Client) *GetEscalationChainByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) WithFields(fields *string) *GetEscalationChainByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) WithID(id int32) *GetEscalationChainByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get escalation chain by Id params
func (o *GetEscalationChainByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEscalationChainByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}