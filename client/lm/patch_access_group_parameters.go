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

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchAccessGroupParams creates a new PatchAccessGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAccessGroupParams() *PatchAccessGroupParams {
	return &PatchAccessGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAccessGroupParamsWithTimeout creates a new PatchAccessGroupParams object
// with the ability to set a timeout on a request.
func NewPatchAccessGroupParamsWithTimeout(timeout time.Duration) *PatchAccessGroupParams {
	return &PatchAccessGroupParams{
		timeout: timeout,
	}
}

// NewPatchAccessGroupParamsWithContext creates a new PatchAccessGroupParams object
// with the ability to set a context for a request.
func NewPatchAccessGroupParamsWithContext(ctx context.Context) *PatchAccessGroupParams {
	return &PatchAccessGroupParams{
		Context: ctx,
	}
}

// NewPatchAccessGroupParamsWithHTTPClient creates a new PatchAccessGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAccessGroupParamsWithHTTPClient(client *http.Client) *PatchAccessGroupParams {
	return &PatchAccessGroupParams{
		HTTPClient: client,
	}
}

/* PatchAccessGroupParams contains all the parameters to send to the API endpoint
   for the patch access group operation.

   Typically these are written to a http.Request.
*/
type PatchAccessGroupParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.AccessGroup

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch access group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAccessGroupParams) WithDefaults() *PatchAccessGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch access group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAccessGroupParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := PatchAccessGroupParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch access group params
func (o *PatchAccessGroupParams) WithTimeout(timeout time.Duration) *PatchAccessGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch access group params
func (o *PatchAccessGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch access group params
func (o *PatchAccessGroupParams) WithContext(ctx context.Context) *PatchAccessGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch access group params
func (o *PatchAccessGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch access group params
func (o *PatchAccessGroupParams) WithHTTPClient(client *http.Client) *PatchAccessGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch access group params
func (o *PatchAccessGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch access group params
func (o *PatchAccessGroupParams) WithUserAgent(userAgent *string) *PatchAccessGroupParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch access group params
func (o *PatchAccessGroupParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch access group params
func (o *PatchAccessGroupParams) WithBody(body *models.AccessGroup) *PatchAccessGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch access group params
func (o *PatchAccessGroupParams) SetBody(body *models.AccessGroup) {
	o.Body = body
}

// WithID adds the id to the patch access group params
func (o *PatchAccessGroupParams) WithID(id int32) *PatchAccessGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch access group params
func (o *PatchAccessGroupParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAccessGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
