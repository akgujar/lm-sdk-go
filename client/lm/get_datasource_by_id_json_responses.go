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

// GetDatasourceByIDJSONReader is a Reader for the GetDatasourceByIDJSON structure.
type GetDatasourceByIDJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatasourceByIDJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatasourceByIDJSONOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetDatasourceByIDJSONTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDatasourceByIDJSONDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDatasourceByIDJSONOK creates a GetDatasourceByIDJSONOK with default headers values
func NewGetDatasourceByIDJSONOK() *GetDatasourceByIDJSONOK {
	return &GetDatasourceByIDJSONOK{}
}

/*
GetDatasourceByIDJSONOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDatasourceByIDJSONOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get datasource by Id Json o k response has a 2xx status code
func (o *GetDatasourceByIDJSONOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get datasource by Id Json o k response has a 3xx status code
func (o *GetDatasourceByIDJSONOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datasource by Id Json o k response has a 4xx status code
func (o *GetDatasourceByIDJSONOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get datasource by Id Json o k response has a 5xx status code
func (o *GetDatasourceByIDJSONOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get datasource by Id Json o k response a status code equal to that given
func (o *GetDatasourceByIDJSONOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get datasource by Id Json o k response
func (o *GetDatasourceByIDJSONOK) Code() int {
	return 200
}

func (o *GetDatasourceByIDJSONOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/datasources/{id}?__json=][%d] getDatasourceByIdJsonOK %s", 200, payload)
}

func (o *GetDatasourceByIDJSONOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/datasources/{id}?__json=][%d] getDatasourceByIdJsonOK %s", 200, payload)
}

func (o *GetDatasourceByIDJSONOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDatasourceByIDJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasourceByIDJSONTooManyRequests creates a GetDatasourceByIDJSONTooManyRequests with default headers values
func NewGetDatasourceByIDJSONTooManyRequests() *GetDatasourceByIDJSONTooManyRequests {
	return &GetDatasourceByIDJSONTooManyRequests{}
}

/*
GetDatasourceByIDJSONTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDatasourceByIDJSONTooManyRequests struct {

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

// IsSuccess returns true when this get datasource by Id Json too many requests response has a 2xx status code
func (o *GetDatasourceByIDJSONTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datasource by Id Json too many requests response has a 3xx status code
func (o *GetDatasourceByIDJSONTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datasource by Id Json too many requests response has a 4xx status code
func (o *GetDatasourceByIDJSONTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get datasource by Id Json too many requests response has a 5xx status code
func (o *GetDatasourceByIDJSONTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get datasource by Id Json too many requests response a status code equal to that given
func (o *GetDatasourceByIDJSONTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get datasource by Id Json too many requests response
func (o *GetDatasourceByIDJSONTooManyRequests) Code() int {
	return 429
}

func (o *GetDatasourceByIDJSONTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{id}?__json=][%d] getDatasourceByIdJsonTooManyRequests", 429)
}

func (o *GetDatasourceByIDJSONTooManyRequests) String() string {
	return fmt.Sprintf("[GET /setting/datasources/{id}?__json=][%d] getDatasourceByIdJsonTooManyRequests", 429)
}

func (o *GetDatasourceByIDJSONTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDatasourceByIDJSONDefault creates a GetDatasourceByIDJSONDefault with default headers values
func NewGetDatasourceByIDJSONDefault(code int) *GetDatasourceByIDJSONDefault {
	return &GetDatasourceByIDJSONDefault{
		_statusCode: code,
	}
}

/*
GetDatasourceByIDJSONDefault describes a response with status code -1, with default header values.

Error
*/
type GetDatasourceByIDJSONDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get datasource by Id Json default response has a 2xx status code
func (o *GetDatasourceByIDJSONDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get datasource by Id Json default response has a 3xx status code
func (o *GetDatasourceByIDJSONDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get datasource by Id Json default response has a 4xx status code
func (o *GetDatasourceByIDJSONDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get datasource by Id Json default response has a 5xx status code
func (o *GetDatasourceByIDJSONDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get datasource by Id Json default response a status code equal to that given
func (o *GetDatasourceByIDJSONDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get datasource by Id Json default response
func (o *GetDatasourceByIDJSONDefault) Code() int {
	return o._statusCode
}

func (o *GetDatasourceByIDJSONDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/datasources/{id}?__json=][%d] getDatasourceByIdJson default %s", o._statusCode, payload)
}

func (o *GetDatasourceByIDJSONDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /setting/datasources/{id}?__json=][%d] getDatasourceByIdJson default %s", o._statusCode, payload)
}

func (o *GetDatasourceByIDJSONDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDatasourceByIDJSONDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}