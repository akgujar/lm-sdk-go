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

// GetImmediateDeviceListByDeviceGroupIDJSONReader is a Reader for the GetImmediateDeviceListByDeviceGroupIDJSON structure.
type GetImmediateDeviceListByDeviceGroupIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImmediateDeviceListByDeviceGroupIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImmediateDeviceListByDeviceGroupIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetImmediateDeviceListByDeviceGroupIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetImmediateDeviceListByDeviceGroupIDJSONOK creates a GetImmediateDeviceListByDeviceGroupIDJSONOK with default headers values
func NewGetImmediateDeviceListByDeviceGroupIDJSONOK() *GetImmediateDeviceListByDeviceGroupIDJSONOK {
	return &GetImmediateDeviceListByDeviceGroupIDJSONOK{}
}

/*
GetImmediateDeviceListByDeviceGroupIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetImmediateDeviceListByDeviceGroupIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get immediate device list by device group Id Json o k response has a 2xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get immediate device list by device group Id Json o k response has a 3xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get immediate device list by device group Id Json o k response has a 4xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get immediate device list by device group Id Json o k response has a 5xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get immediate device list by device group Id Json o k response a status code equal to that given
func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get immediate device list by device group Id Json o k response
func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) Code() int {
	return 200
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/devices?__json=][%d] getImmediateDeviceListByDeviceGroupIdJsonOK %s", 200, payload)
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/devices?__json=][%d] getImmediateDeviceListByDeviceGroupIdJsonOK %s", 200, payload)
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests creates a GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests with default headers values
func NewGetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests() *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests {
	return &GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests{}
}

/*
GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get immediate device list by device group Id Json too many requests response has a 2xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get immediate device list by device group Id Json too many requests response has a 3xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get immediate device list by device group Id Json too many requests response has a 4xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get immediate device list by device group Id Json too many requests response has a 5xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get immediate device list by device group Id Json too many requests response a status code equal to that given
func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get immediate device list by device group Id Json too many requests response
func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device/groups/{id}/devices?__json=][%d] getImmediateDeviceListByDeviceGroupIdJsonTooManyRequests", 429)
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /device/groups/{id}/devices?__json=][%d] getImmediateDeviceListByDeviceGroupIdJsonTooManyRequests", 429)
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetImmediateDeviceListByDeviceGroupIDJSONDefault creates a GetImmediateDeviceListByDeviceGroupIDJSONDefault with default headers values
func NewGetImmediateDeviceListByDeviceGroupIDJSONDefault(code int) *GetImmediateDeviceListByDeviceGroupIDJSONDefault {
	return &GetImmediateDeviceListByDeviceGroupIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetImmediateDeviceListByDeviceGroupIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetImmediateDeviceListByDeviceGroupIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get immediate device list by device group Id Json default response has a 2xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get immediate device list by device group Id Json default response has a 3xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get immediate device list by device group Id Json default response has a 4xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get immediate device list by device group Id Json default response has a 5xx status code
func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get immediate device list by device group Id Json default response a status code equal to that given
func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get immediate device list by device group Id Json default response
func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/devices?__json=][%d] getImmediateDeviceListByDeviceGroupIdJson default %s", o._statusCode, payload)
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /device/groups/{id}/devices?__json=][%d] getImmediateDeviceListByDeviceGroupIdJson default %s", o._statusCode, payload)
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetImmediateDeviceListByDeviceGroupIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
