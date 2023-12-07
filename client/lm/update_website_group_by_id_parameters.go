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

// NewUpdateWebsiteGroupByIDParams creates a new UpdateWebsiteGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateWebsiteGroupByIDParams() *UpdateWebsiteGroupByIDParams {
	return &UpdateWebsiteGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateWebsiteGroupByIDParamsWithTimeout creates a new UpdateWebsiteGroupByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateWebsiteGroupByIDParamsWithTimeout(timeout time.Duration) *UpdateWebsiteGroupByIDParams {
	return &UpdateWebsiteGroupByIDParams{
		timeout: timeout,
	}
}

// NewUpdateWebsiteGroupByIDParamsWithContext creates a new UpdateWebsiteGroupByIDParams object
// with the ability to set a context for a request.
func NewUpdateWebsiteGroupByIDParamsWithContext(ctx context.Context) *UpdateWebsiteGroupByIDParams {
	return &UpdateWebsiteGroupByIDParams{
		Context: ctx,
	}
}

// NewUpdateWebsiteGroupByIDParamsWithHTTPClient creates a new UpdateWebsiteGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateWebsiteGroupByIDParamsWithHTTPClient(client *http.Client) *UpdateWebsiteGroupByIDParams {
	return &UpdateWebsiteGroupByIDParams{
		HTTPClient: client,
	}
}

/* UpdateWebsiteGroupByIDParams contains all the parameters to send to the API endpoint
   for the update website group by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateWebsiteGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.WebsiteGroup

	// ID.
	//
	// Format: int32
	ID int32

	// OpType.
	//
	// Default: "refresh"
	OpType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWebsiteGroupByIDParams) WithDefaults() *UpdateWebsiteGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update website group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateWebsiteGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")

		opTypeDefault = string("refresh")
	)

	val := UpdateWebsiteGroupByIDParams{
		UserAgent: &userAgentDefault,
		OpType:    &opTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) WithTimeout(timeout time.Duration) *UpdateWebsiteGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) WithContext(ctx context.Context) *UpdateWebsiteGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) WithHTTPClient(client *http.Client) *UpdateWebsiteGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) WithUserAgent(userAgent *string) *UpdateWebsiteGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) WithBody(body *models.WebsiteGroup) *UpdateWebsiteGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) SetBody(body *models.WebsiteGroup) {
	o.Body = body
}

// WithID adds the id to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) WithID(id int32) *UpdateWebsiteGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOpType adds the opType to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) WithOpType(opType *string) *UpdateWebsiteGroupByIDParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the update website group by Id params
func (o *UpdateWebsiteGroupByIDParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateWebsiteGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OpType != nil {

		// query param opType
		var qrOpType string

		if o.OpType != nil {
			qrOpType = *o.OpType
		}
		qOpType := qrOpType
		if qOpType != "" {

			if err := r.SetQueryParam("opType", qOpType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
