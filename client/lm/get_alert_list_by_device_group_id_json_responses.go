// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/v3/models"
)

// GetAlertListByDeviceGroupIDJSONReader is a Reader for the GetAlertListByDeviceGroupIDJSON structure.
type GetAlertListByDeviceGroupIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertListByDeviceGroupIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertListByDeviceGroupIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAlertListByDeviceGroupIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAlertListByDeviceGroupIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertListByDeviceGroupIDJSONOK creates a GetAlertListByDeviceGroupIDJSONOK with default headers values
func NewGetAlertListByDeviceGroupIDJSONOK() *GetAlertListByDeviceGroupIDJSONOK {
	return &GetAlertListByDeviceGroupIDJSONOK{}
}

/*
GetAlertListByDeviceGroupIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertListByDeviceGroupIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get alert list by device group Id Json o k response has a 2xx status code
func (o *GetAlertListByDeviceGroupIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert list by device group Id Json o k response has a 3xx status code
func (o *GetAlertListByDeviceGroupIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert list by device group Id Json o k response has a 4xx status code
func (o *GetAlertListByDeviceGroupIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert list by device group Id Json o k response has a 5xx status code
func (o *GetAlertListByDeviceGroupIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert list by device group Id Json o k response a status code equal to that given
func (o *GetAlertListByDeviceGroupIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert list by device group Id Json o k response
func (o *GetAlertListByDeviceGroupIDJSONOK) Code() int {
	return 200
}

func (o *GetAlertListByDeviceGroupIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/alerts?__json=][%d] getAlertListByDeviceGroupIdJsonOK %s", 200, payload)
}

func (o *GetAlertListByDeviceGroupIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/alerts?__json=][%d] getAlertListByDeviceGroupIdJsonOK %s", 200, payload)
}

func (o *GetAlertListByDeviceGroupIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAlertListByDeviceGroupIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertListByDeviceGroupIDJSONTooManyRequests creates a GetAlertListByDeviceGroupIDJSONTooManyRequests with default headers values
func NewGetAlertListByDeviceGroupIDJSONTooManyRequests() *GetAlertListByDeviceGroupIDJSONTooManyRequests {
	return &GetAlertListByDeviceGroupIDJSONTooManyRequests{}
}

/*
GetAlertListByDeviceGroupIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAlertListByDeviceGroupIDJSONTooManyRequests struct {

	/* Request limit per X-Rate-Limit-Window
	 */
	XRateLimitLimit int64

	/* The number of requests left for the time window
	 */
	XRateLimitRemaining int64

	/* The rolling time window length with the unit of second
	 */
	XRateLimitWindow int64
}

// IsSuccess returns true when this get alert list by device group Id Json too many requests response has a 2xx status code
func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert list by device group Id Json too many requests response has a 3xx status code
func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert list by device group Id Json too many requests response has a 4xx status code
func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert list by device group Id Json too many requests response has a 5xx status code
func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert list by device group Id Json too many requests response a status code equal to that given
func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get alert list by device group Id Json too many requests response
func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/groups/{id}/alerts?__json=][%d] getAlertListByDeviceGroupIdJsonTooManyRequests", 429)
}

func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /device/groups/{id}/alerts?__json=][%d] getAlertListByDeviceGroupIdJsonTooManyRequests", 429)
}

func (o *GetAlertListByDeviceGroupIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-rate-limit-limit
	hdrXRateLimitLimit := response.GetHeader("x-rate-limit-limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("x-rate-limit-limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header x-rate-limit-remaining
	hdrXRateLimitRemaining := response.GetHeader("x-rate-limit-remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("x-rate-limit-remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header x-rate-limit-window
	hdrXRateLimitWindow := response.GetHeader("x-rate-limit-window")

	if hdrXRateLimitWindow != "" {
		valxRateLimitWindow, err := swag.ConvertInt64(hdrXRateLimitWindow)
		if err != nil {
			return errors.InvalidType("x-rate-limit-window", "header", "int64", hdrXRateLimitWindow)
		}
		o.XRateLimitWindow = valxRateLimitWindow
	}

	return nil
}

// NewGetAlertListByDeviceGroupIDJSONDefault creates a GetAlertListByDeviceGroupIDJSONDefault with default headers values
func NewGetAlertListByDeviceGroupIDJSONDefault(code int) *GetAlertListByDeviceGroupIDJSONDefault {
	return &GetAlertListByDeviceGroupIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetAlertListByDeviceGroupIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertListByDeviceGroupIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get alert list by device group Id Json default response has a 2xx status code
func (o *GetAlertListByDeviceGroupIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get alert list by device group Id Json default response has a 3xx status code
func (o *GetAlertListByDeviceGroupIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get alert list by device group Id Json default response has a 4xx status code
func (o *GetAlertListByDeviceGroupIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get alert list by device group Id Json default response has a 5xx status code
func (o *GetAlertListByDeviceGroupIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get alert list by device group Id Json default response a status code equal to that given
func (o *GetAlertListByDeviceGroupIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get alert list by device group Id Json default response
func (o *GetAlertListByDeviceGroupIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertListByDeviceGroupIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/alerts?__json=][%d] getAlertListByDeviceGroupIdJson default %s", o._statusCode, payload)
}

func (o *GetAlertListByDeviceGroupIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/alerts?__json=][%d] getAlertListByDeviceGroupIdJson default %s", o._statusCode, payload)
}

func (o *GetAlertListByDeviceGroupIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAlertListByDeviceGroupIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}