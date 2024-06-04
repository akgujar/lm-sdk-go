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

// AddDatasourceByIDReader is a Reader for the AddDatasourceByID structure.
type AddDatasourceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDatasourceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDatasourceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewAddDatasourceByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddDatasourceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDatasourceByIDOK creates a AddDatasourceByIDOK with default headers values
func NewAddDatasourceByIDOK() *AddDatasourceByIDOK {
	return &AddDatasourceByIDOK{}
}

/*
AddDatasourceByIDOK describes a response with status code 200, with default header values.

successful operation
*/
type AddDatasourceByIDOK struct {
	Payload *models.DataSource
}

// IsSuccess returns true when this add datasource by Id o k response has a 2xx status code
func (o *AddDatasourceByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add datasource by Id o k response has a 3xx status code
func (o *AddDatasourceByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add datasource by Id o k response has a 4xx status code
func (o *AddDatasourceByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add datasource by Id o k response has a 5xx status code
func (o *AddDatasourceByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add datasource by Id o k response a status code equal to that given
func (o *AddDatasourceByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add datasource by Id o k response
func (o *AddDatasourceByIDOK) Code() int {
	return 200
}

func (o *AddDatasourceByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/datasources][%d] addDatasourceByIdOK %s", 200, payload)
}

func (o *AddDatasourceByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/datasources][%d] addDatasourceByIdOK %s", 200, payload)
}

func (o *AddDatasourceByIDOK) GetPayload() *models.DataSource {
	return o.Payload
}

func (o *AddDatasourceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDatasourceByIDTooManyRequests creates a AddDatasourceByIDTooManyRequests with default headers values
func NewAddDatasourceByIDTooManyRequests() *AddDatasourceByIDTooManyRequests {
	return &AddDatasourceByIDTooManyRequests{}
}

/*
AddDatasourceByIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AddDatasourceByIDTooManyRequests struct {

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

// IsSuccess returns true when this add datasource by Id too many requests response has a 2xx status code
func (o *AddDatasourceByIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add datasource by Id too many requests response has a 3xx status code
func (o *AddDatasourceByIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add datasource by Id too many requests response has a 4xx status code
func (o *AddDatasourceByIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this add datasource by Id too many requests response has a 5xx status code
func (o *AddDatasourceByIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this add datasource by Id too many requests response a status code equal to that given
func (o *AddDatasourceByIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the add datasource by Id too many requests response
func (o *AddDatasourceByIDTooManyRequests) Code() int {
	return 429
}

func (o *AddDatasourceByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /setting/datasources][%d] addDatasourceByIdTooManyRequests", 429)
}

func (o *AddDatasourceByIDTooManyRequests) String() string {
	return fmt.Sprintf("[POST /setting/datasources][%d] addDatasourceByIdTooManyRequests", 429)
}

func (o *AddDatasourceByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddDatasourceByIDDefault creates a AddDatasourceByIDDefault with default headers values
func NewAddDatasourceByIDDefault(code int) *AddDatasourceByIDDefault {
	return &AddDatasourceByIDDefault{
		_statusCode: code,
	}
}

/*
AddDatasourceByIDDefault describes a response with status code -1, with default header values.

Error
*/
type AddDatasourceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add datasource by Id default response has a 2xx status code
func (o *AddDatasourceByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add datasource by Id default response has a 3xx status code
func (o *AddDatasourceByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add datasource by Id default response has a 4xx status code
func (o *AddDatasourceByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add datasource by Id default response has a 5xx status code
func (o *AddDatasourceByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add datasource by Id default response a status code equal to that given
func (o *AddDatasourceByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add datasource by Id default response
func (o *AddDatasourceByIDDefault) Code() int {
	return o._statusCode
}

func (o *AddDatasourceByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/datasources][%d] addDatasourceById default %s", o._statusCode, payload)
}

func (o *AddDatasourceByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /setting/datasources][%d] addDatasourceById default %s", o._statusCode, payload)
}

func (o *AddDatasourceByIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddDatasourceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
