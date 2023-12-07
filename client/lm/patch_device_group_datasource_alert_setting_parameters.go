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

// NewPatchDeviceGroupDatasourceAlertSettingParams creates a new PatchDeviceGroupDatasourceAlertSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDeviceGroupDatasourceAlertSettingParams() *PatchDeviceGroupDatasourceAlertSettingParams {
	return &PatchDeviceGroupDatasourceAlertSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeviceGroupDatasourceAlertSettingParamsWithTimeout creates a new PatchDeviceGroupDatasourceAlertSettingParams object
// with the ability to set a timeout on a request.
func NewPatchDeviceGroupDatasourceAlertSettingParamsWithTimeout(timeout time.Duration) *PatchDeviceGroupDatasourceAlertSettingParams {
	return &PatchDeviceGroupDatasourceAlertSettingParams{
		timeout: timeout,
	}
}

// NewPatchDeviceGroupDatasourceAlertSettingParamsWithContext creates a new PatchDeviceGroupDatasourceAlertSettingParams object
// with the ability to set a context for a request.
func NewPatchDeviceGroupDatasourceAlertSettingParamsWithContext(ctx context.Context) *PatchDeviceGroupDatasourceAlertSettingParams {
	return &PatchDeviceGroupDatasourceAlertSettingParams{
		Context: ctx,
	}
}

// NewPatchDeviceGroupDatasourceAlertSettingParamsWithHTTPClient creates a new PatchDeviceGroupDatasourceAlertSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDeviceGroupDatasourceAlertSettingParamsWithHTTPClient(client *http.Client) *PatchDeviceGroupDatasourceAlertSettingParams {
	return &PatchDeviceGroupDatasourceAlertSettingParams{
		HTTPClient: client,
	}
}

/* PatchDeviceGroupDatasourceAlertSettingParams contains all the parameters to send to the API endpoint
   for the patch device group datasource alert setting operation.

   Typically these are written to a http.Request.
*/
type PatchDeviceGroupDatasourceAlertSettingParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/GO-SDK"
	UserAgent *string

	// Body.
	Body *models.DeviceGroupDataSourceAlertConfig

	// DeviceGroupID.
	//
	// Format: int32
	DeviceGroupID int32

	// DsID.
	//
	// Format: int32
	DsID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch device group datasource alert setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithDefaults() *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch device group datasource alert setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/GO-SDK")
	)

	val := PatchDeviceGroupDatasourceAlertSettingParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithTimeout(timeout time.Duration) *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithContext(ctx context.Context) *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithHTTPClient(client *http.Client) *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithUserAgent(userAgent *string) *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithBody(body *models.DeviceGroupDataSourceAlertConfig) *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetBody(body *models.DeviceGroupDataSourceAlertConfig) {
	o.Body = body
}

// WithDeviceGroupID adds the deviceGroupID to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithDeviceGroupID(deviceGroupID int32) *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithDsID adds the dsID to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WithDsID(dsID int32) *PatchDeviceGroupDatasourceAlertSettingParams {
	o.SetDsID(dsID)
	return o
}

// SetDsID adds the dsId to the patch device group datasource alert setting params
func (o *PatchDeviceGroupDatasourceAlertSettingParams) SetDsID(dsID int32) {
	o.DsID = dsID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeviceGroupDatasourceAlertSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

	// path param dsId
	if err := r.SetPathParam("dsId", swag.FormatInt32(o.DsID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
