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

// UpdateDatasourceByIDReader is a Reader for the UpdateDatasourceByID structure.
type UpdateDatasourceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDatasourceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDatasourceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewUpdateDatasourceByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDatasourceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDatasourceByIDOK creates a UpdateDatasourceByIDOK with default headers values
func NewUpdateDatasourceByIDOK() *UpdateDatasourceByIDOK {
	return &UpdateDatasourceByIDOK{}
}

/*
UpdateDatasourceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateDatasourceByIDOK struct {
	Payload *models.DataSource
}

// IsSuccess returns true when this update datasource by Id o k response has a 2xx status code
func (o *UpdateDatasourceByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update datasource by Id o k response has a 3xx status code
func (o *UpdateDatasourceByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update datasource by Id o k response has a 4xx status code
func (o *UpdateDatasourceByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update datasource by Id o k response has a 5xx status code
func (o *UpdateDatasourceByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update datasource by Id o k response a status code equal to that given
func (o *UpdateDatasourceByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update datasource by Id o k response
func (o *UpdateDatasourceByIDOK) Code() int {
	return 200
}

func (o *UpdateDatasourceByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /setting/datasources/{id}][%d] updateDatasourceByIdOK %s", 200, payload)
}

func (o *UpdateDatasourceByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /setting/datasources/{id}][%d] updateDatasourceByIdOK %s", 200, payload)
}

func (o *UpdateDatasourceByIDOK) GetPayload() *models.DataSource {
	return o.Payload
}

func (o *UpdateDatasourceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatasourceByIDTooManyRequests creates a UpdateDatasourceByIDTooManyRequests with default headers values
func NewUpdateDatasourceByIDTooManyRequests() *UpdateDatasourceByIDTooManyRequests {
	return &UpdateDatasourceByIDTooManyRequests{}
}

/*
UpdateDatasourceByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDatasourceByIDTooManyRequests struct {

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

// IsSuccess returns true when this update datasource by Id too many requests response has a 2xx status code
func (o *UpdateDatasourceByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update datasource by Id too many requests response has a 3xx status code
func (o *UpdateDatasourceByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update datasource by Id too many requests response has a 4xx status code
func (o *UpdateDatasourceByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update datasource by Id too many requests response has a 5xx status code
func (o *UpdateDatasourceByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update datasource by Id too many requests response a status code equal to that given
func (o *UpdateDatasourceByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update datasource by Id too many requests response
func (o *UpdateDatasourceByIDTooManyRequests) Code() int {
	return 429
}

func (o *UpdateDatasourceByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /setting/datasources/{id}][%d] updateDatasourceByIdTooManyRequests", 429)
}

func (o *UpdateDatasourceByIDTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /setting/datasources/{id}][%d] updateDatasourceByIdTooManyRequests", 429)
}

func (o *UpdateDatasourceByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDatasourceByIDDefault creates a UpdateDatasourceByIDDefault with default headers values
func NewUpdateDatasourceByIDDefault(code int) *UpdateDatasourceByIDDefault {
	return &UpdateDatasourceByIDDefault{
		_statusCode: code,
	}
}

/*
UpdateDatasourceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateDatasourceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update datasource by Id default response has a 2xx status code
func (o *UpdateDatasourceByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update datasource by Id default response has a 3xx status code
func (o *UpdateDatasourceByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update datasource by Id default response has a 4xx status code
func (o *UpdateDatasourceByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update datasource by Id default response has a 5xx status code
func (o *UpdateDatasourceByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update datasource by Id default response a status code equal to that given
func (o *UpdateDatasourceByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update datasource by Id default response
func (o *UpdateDatasourceByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDatasourceByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /setting/datasources/{id}][%d] updateDatasourceById default %s", o._statusCode, payload)
}

func (o *UpdateDatasourceByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /setting/datasources/{id}][%d] updateDatasourceById default %s", o._statusCode, payload)
}

func (o *UpdateDatasourceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDatasourceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
